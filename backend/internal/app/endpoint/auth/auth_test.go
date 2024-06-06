package auth_test

import (
	"fmt"
	"io"
	"log"
	"moscowhack/internal/app/config"
	"moscowhack/internal/app/service/cacher"
	"moscowhack/pkg/logger"
	"net/http"
	"testing"

	"go.uber.org/zap"
)

func TestAuth(t *testing.T) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Ошибка при попытке спарсить .env файл в структуру: %v", err)
	}

	logger.Init(cfg.LoggerLevel)
	cacher.Init(cfg.Cache.CacheEXTime)

	client := &http.Client{}

	t.Run("NewUserRegistration Test", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/NewUserRegistration", nil)
		if err != nil {
			logger.Error("Ошибка при отправке запроса: ", zap.Error(err))
			return
		}

		req.Header.Set("username", "testuser")
		req.Header.Set("password", "testpass")
		res, err := client.Do(req)
		if err != nil {
			logger.Error("Ошибка при отправке запроса: ", zap.Error(err))
			return
		}

		receivedData, _ := io.ReadAll(res.Body)

		fmt.Printf("Получены данные от сервера: %s", receivedData)
	})

	t.Run("UserLogin Test", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/NewUserRegistration", nil)
		if err != nil {
			logger.Error("Ошибка при отправке запроса: ", zap.Error(err))
			return
		}

		req.Header.Set("username", "testuser")
		req.Header.Set("password", "testpass")
		res, err := client.Do(req)
		if err != nil {
			logger.Error("Ошибка при отправке запроса: ", zap.Error(err))
			return
		}

		receivedData, _ := io.ReadAll(res.Body)
		fmt.Printf("Получены данные от сервера: %v", receivedData)
	})
}
