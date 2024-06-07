package auth

import (
	"context"
	"database/sql"
	pb "moscowhack/gen/go/auth"
	"moscowhack/internal/app/errorz"
	"moscowhack/internal/app/service/db"
	"moscowhack/pkg/cache"
	"moscowhack/pkg/logger"

	"go.uber.org/zap"
)

type Service interface {
	Keygen(string, string) (string, error)
	TokenAuth(string) (bool, error)
}

type Endpoint struct {
	s      Service
	server AuthServiceServer
}

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) UserLogin(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
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
	key, err := e.s.Keygen(req.Username, req.Password)
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
			cache.SaveCache("users", userData)

			// авторизуем его
			return &pb.AuthResponse{Key: key}, nil
		}
	}

	return nil, errorz.ErrUserNotFound
}

// передаем в эту функцию username и password
func (e *Endpoint) NewUserRegistration(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
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

	key, err := e.s.Keygen(req.Username, req.Password)
	if err != nil {
		logger.Error("ошибка при генерации токена: %s", zap.Error(err))
		return nil, err
	}

	return &pb.AuthResponse{Key: key}, nil
}
