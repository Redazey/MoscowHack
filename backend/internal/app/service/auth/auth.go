package auth

import (
	"context"
	"database/sql"
	"moscowhack/internal/app/errorz"
	"moscowhack/internal/app/lib/db"
	"moscowhack/internal/app/lib/jwt"
	"moscowhack/pkg/cache"
	"moscowhack/pkg/logger"

	"go.uber.org/zap"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) UserLogin(ctx context.Context, username string, password string) (string, error) {
	msg := map[string]string{
		"username": username,
		"password": password,
	}
	userData, hashKey := cache.ConvertMap(msg, "username", "password")

	cachePwd, err := cache.IsDataInCache("users", hashKey, "password")
	if err != nil {
		logger.Error("ошибка при поиске данных в кэше Redis: ", zap.Error(err))
		return "", err
	}

	if cachePwd != "" && cachePwd != password {
		return "", errorz.ErrUserNotFound
	} else if cachePwd == "" {
		dbMap, err := db.FetchUserData(username)
		if err != nil {
			return "", err
		}
		if dbMap != nil && dbMap["password"] != password {
			return "", err
		}
	}

	// генерируем jwt токен и данных юзера для использования в дальнейшем
	key, err := jwt.Keygen(username, password)
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

func (s *Service) NewUserRegistration(ctx context.Context, username string, password string) (string, error) {
	msg := map[string]string{
		"username": username,
		"password": password,
	}
	userData, hashKey := cache.ConvertMap(msg, "username", "password")

	cachePwd, err := cache.IsDataInCache("users", hashKey, "password")
	if err != nil {
		return "", err
	}

	// если пароль у юзера есть, значит и юзер существует
	if cachePwd == "" {
		dbMap, err := db.FetchUserData(username)
		if err != sql.ErrNoRows && err != nil {
			return "", err
		}

		if len(dbMap) != 0 {
			return "", errorz.ErrUserExists
		}
	}

	err = cache.SaveCache("users", userData)
	if err != nil {
		logger.Error("ошибка при регистрации пользователя: ", zap.Error(err))
		return "", err
	}

	key, err := jwt.Keygen(username, password)
	if err != nil {
		logger.Error("ошибка при генерации токена: ", zap.Error(err))
		return "", err
	}

	return key, nil
}
