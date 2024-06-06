package jwtAuth_test

import (
	"moscowhack/internal/app/config"
	"moscowhack/internal/app/errorz"
	"moscowhack/internal/app/service/jwtAuth"
	"moscowhack/pkg/logger"

	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestJwtAuth(t *testing.T) {
	s := jwtAuth.New()

	config.Init()
	config := config.GetConfig()
	logger.Init(config.LoggerMode)
	godotenv.Load(config.EnvPath)

	err := godotenv.Load(config.EnvPath)
	assert.Nil(t, err, "Ошибка при открытии env файла: %v", err)

	secret := os.Getenv("JWT_KEY")
	username := "username"
	password := "testpwd"

	t.Run("GenerateToken", func(t *testing.T) {
		// Вызов тестируемой функции
		tokenString, err := s.Keygen(username, password)
		assert.Nil(t, err, "Не ожидаем ошибку, получили: %v", err)

		// Проверка наличия сообщения возвращаемого токена
		assert.NotNil(t, tokenString, "Ожидаем наличие сообщения в токене")

		// Проверка на генерацию корректного JWT-токена
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		assert.Nil(t, err, "Неверный секретный ключ: %v", err)

		assert.NotNil(t, token.Claims.(jwt.MapClaims)["username"],
			"Ошибка в данных токена: отсутствует информация о пользователе")

		exp := token.Claims.(jwt.MapClaims)["exp"].(float64)
		expTime := time.Unix(int64(exp), 0)
		assert.False(t, expTime.Before(time.Now()), "Срок действия токена истек")
	})
	t.Run("TokenAuth", func(t *testing.T) {
		tokenString, err := s.Keygen(username, password)
		// Вызываем тестируемую функцию
		isTokenValid, err := s.TokenAuth(tokenString)
		assert.Equal(t, isTokenValid, true)
		assert.Nil(t, err, "Не ожидаем ошибку, получили: %v", err)

		// Проверяем корректность обработки истекшего токена
		claims := jwt.MapClaims{}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		assert.Nil(t, err, "Неверный секретный ключ: %v", err)

		_ = token.Claims.Valid()

		now := time.Now()
		claims["exp"] = now.Unix()

		tokenString, _ = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))

		isTokenValid, err = s.TokenAuth(tokenString)
		assert.Equal(t, isTokenValid, false)
		assert.EqualError(t, err, errorz.ErrTokenExpired.Error(),
			"Ожидаем ошибку TokenExpired, получили другую ошибку: %v", err)
	})
}
