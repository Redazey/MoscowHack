package auth

import (
	"context"
	"database/sql"
	"moscowhack/config"
	pb "moscowhack/gen/go/auth"
	"moscowhack/internal/app/errorz"
	"moscowhack/internal/app/lib/db"
	"moscowhack/internal/app/lib/jwt"
	"moscowhack/pkg/cache"
	"moscowhack/pkg/logger"

	"go.uber.org/zap"
)

type Service struct {
	Cfg *config.Configuration
}

func New(cfg *config.Configuration) *Service {
	return &Service{
		Cfg: cfg,
	}
}

func (s *Service) UserLogin(ctx context.Context, email string, password string) (string, error) {
	hashKey := cache.GetHash(email, password)

	cachePwd, err := cache.IsDataInCache("users", hashKey, "password")
	if err != nil {
		logger.Error("ошибка при поиске данных в кэше Redis: ", zap.Error(err))
		return "", err
	}

	if cachePwd != "" && cachePwd != password {
		return "", errorz.ErrUserNotFound
	} else if cachePwd == "" {
		dbMap, err := db.FetchUserData(email)
		if err != nil {
			return "", err
		}
		if dbMap != nil && dbMap["password"] != password {
			return "", err
		}
	}

	// генерируем jwt токен и данных юзера для использования в дальнейшем
	key, err := jwt.Keygen(email, password, s.Cfg.JwtSecret)
	if err != nil {
		logger.Error("ошибка при генерации токена: ", zap.Error(err))
		return "", err
	}

	// сохраняем залогиненого юзера в кэш
	err = cache.SaveCache("users", userData)
	if err != nil {
		return "", err
	}

	// авторизуем его
	return key, nil
}

func (s *Service) NewUserRegistration(ctx context.Context, req *pb.RegistrationRequest) (string, error) {
	msg := map[string]interface{}{
		"surname":    req.Surname,
		"name":       req.Name,
		"patronymic": req.Patronymic,
		"email":      req.Email,
		"password":   req.Password,
		"roleid":     req.RoleId,
		"birthdate":  req.Birthdate,
		"photourl":   req.Photourl,
		"push":       req.Push,
	}
	userData, hashKey := cache.ConvertMap(msg, "email", "password")

	cachePwd, err := cache.IsDataInCache("users", hashKey, "password")
	if err != nil {
		return "", err
	}

	// если пароль у юзера есть, значит и юзер существует
	if cachePwd == "" {
		dbMap, err := db.FetchUserData(req.Email)
		if err != sql.ErrNoRows && err != nil {
			return "", err
		}

		if len(dbMap) != 0 {
			return "", errorz.ErrUserExists
		}
	}

	key, err := jwt.Keygen(req.Email, req.Password, s.Cfg.JwtSecret)
	if err != nil {
		logger.Error("ошибка при генерации токена: ", zap.Error(err))
		return "", err
	}

	return key, nil
}

func (s *Service) IsAdmin(tokenString string) (bool, error) {
	UserData, err := jwt.UserDataFromJwt(tokenString, s.Cfg.JwtSecret)
	if err != nil {
		return false, err
	}

	_, hashKey := cache.ConvertMap(UserData, "email", "password")

	roleId, err := cache.IsDataInCache("users", hashKey, "roleId")
	if err != nil {
		return false, err
	}

	if roleId == 0 {
		DbUserData, err := db.FetchUserData(UserData["email"])
		if err != nil {
			return false, err
		}

		if DbUserData == nil {
			return false, errorz.ErrUserNotFound
		}

		roleId = UserData["roleid"]

		if roleId == 0 {
			return false, nil
		}
	}

	return true, nil
}
