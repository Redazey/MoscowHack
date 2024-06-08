package tests

import (
	"context"
	"fmt"
	"log"
	pbAuth "moscowhack/gen/go/auth"
	"testing"

	"google.golang.org/grpc"
)

func TestAuth(t *testing.T) {
	req := &pbAuth.AuthRequest{
		Username: "testUser",
		Password: "testPassword",
	}

	t.Run("NewUserRegistration Test", func(t *testing.T) {
		conn, err := grpc.NewClient("localhost:50051")
		if err != nil {
			log.Fatalf("Failed to dial server: %v", err)
		}
		defer conn.Close()

		// Инициализируем клиент
		client := pbAuth.AuthServiceClient()

		// Пример вызова функции на сервере
		request := &pb.YourRequest{
			// Укажите данные для запроса
		}
		response, err := client.YourFunction(context.Background(), request)
		if err != nil {
			log.Fatalf("Error when calling YourFunction: %v", err)
		}

		log.Printf("Response from server: %v", response)
		fmt.Printf("Получены данные от сервера: %s", authResp)
	})
	/*
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
		})*/
}
