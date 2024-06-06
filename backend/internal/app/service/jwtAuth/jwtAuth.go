package jwtAuth

import (
	"os"
	"time"

	"moscowhack/internal/app/errorz"
	"moscowhack/pkg/logger"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Keygen(username string, pwd string) (string, error) {
	// Создаем новый JWT токен
	token := jwt.New(jwt.SigningMethodHS256)

	// Указываем параметры для токена
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["password"] = pwd
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	secretKey := []byte(os.Getenv("JWT_KEY"))

	// Подписываем токен с помощью указанного секретного ключа
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		logger.Error("ошибка при подписи токена: ", zap.Error(err))
		return "", err
	}

	return tokenString, nil
}

func (s *Service) TokenAuth(tokenString string) (bool, error) {
	secretKey := []byte(os.Getenv("JWT_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		logger.Error("ошибка при чтении токена: ", zap.Error(err))
		return false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {

		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)

		if time.Now().After(expirationTime) {
			return false, errorz.ErrTokenExpired
		}

	} else {
		logger.Error("невалидный токен: ", zap.Error(errorz.ErrValidation))
		return false, errorz.ErrValidation
	}

	return true, nil
}
