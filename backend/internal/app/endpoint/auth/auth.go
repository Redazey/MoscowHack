package auth

import (
	"context"
	"database/sql"
	"errors"
	pb "moscowhack/gen/go/auth"
	"moscowhack/internal/app/errorz"
	"moscowhack/internal/app/lib/db"
	"moscowhack/pkg/cache"
	"moscowhack/pkg/logger"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Auth interface {
	Keygen(string, string) (string, error)
	TokenAuth(string) (bool, error)
}

type Endpoint struct {
	auth Auth
	pb.UnimplementedAuthServiceServer
}

func Register(gRPCServer *grpc.Server, auth Auth) {
	pb.RegisterAuthServiceServer(gRPCServer, &Endpoint{auth: auth})
}

func New(auth Auth) *Endpoint {
	return &Endpoint{
		auth: auth,
	}
}

func (e *Endpoint) UserLogin(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	if req.Username == "" && req.Password == "" {
		return nil, errors.New("username и password пусты в UserLogin")
	}

	msg := map[string]string{
		"username": req.Username,
		"password": req.Password,
	}
	userData, hashKey := cache.ConvertMap(msg, "username", "password")

	cachePwd, err := cache.IsDataInCache("users", hashKey, "password")
	if err != nil {
		logger.Error("ошибка при поиске данных в кэше Redis: ", zap.Error(err))
		return nil, err
	}

	// генерируем jwt токен и данных юзера для использования в дальнейшем
	key, err := e.auth.Keygen(req.Username, req.Password)
	if err != nil {
		logger.Error("ошибка при генерации токена: %s", zap.Error(err))
		return nil, err
	}

	if cachePwd != "" && cachePwd == req.Password {
		return &pb.AuthResponse{Key: key}, nil
	} else if cachePwd == nil {
		dbMap, err := db.FetchUserData(req.Username)
		if err != nil {
			return nil, err
		}
		if dbMap != nil && dbMap["password"] == req.Password {
			// сохраняем залогиненого юзера в кэш
			err := cache.SaveCache("users", userData)
			if err != nil {
				return nil, err
			}

			// авторизуем его
			return &pb.AuthResponse{Key: key}, nil
		}
	}

	return nil, errorz.ErrUserNotFound
}

// передаем в эту функцию username и password
func (e *Endpoint) NewUserRegistration(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	if req.Username == "" && req.Password == "" {
		return nil, errors.New("username и password пусты в UserLogin")
	}

	msg := map[string]string{
		"username": req.Username,
		"password": req.Password,
	}
	userData, hashKey := cache.ConvertMap(msg, "username", "password")

	cachePwd, err := cache.IsDataInCache("users", hashKey, "password")
	if err != nil {
		return nil, err
	}

	// если пароль у юзера есть, значит и юзер существует
	if cachePwd == "" {
		dbMap, err := db.FetchUserData(req.Username)
		if err != sql.ErrNoRows && err != nil {
			return nil, err
		}

		if len(dbMap) != 0 {
			logger.Info("такой пользователь уже существует")
			return nil, errorz.ErrUserExists
		}
	}

	err = cache.SaveCache("users", userData)
	if err != nil {
		logger.Error("ошибка при регистрации пользователя: %s", zap.Error(err))
		return nil, err
	}

	key, err := e.auth.Keygen(req.Username, req.Password)
	if err != nil {
		logger.Error("ошибка при генерации токена: %s", zap.Error(err))
		return nil, err
	}

	return &pb.AuthResponse{Key: key}, nil
}
