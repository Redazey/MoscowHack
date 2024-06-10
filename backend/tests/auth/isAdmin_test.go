package auth_tests

import (
	"log"
	pbAuth "moscowhack/gen/go/auth"
	"moscowhack/internal/app/lib/jwt"
	"moscowhack/tests/suite"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAdmin(t *testing.T) {
	ctx, st := suite.New(t)
	// 1 - админ, 2 - юзер
	AdminData, err := MockUser(1)
	if err != nil {
		log.Fatalf("Ошибка при добавлении тестового админа в бд: %v", err)
	}

	UserData, err := MockUser(2)
	if err != nil {
		log.Fatalf("Ошибка при добавлении тестового пользователя в бд: %v", err)
	}

	t.Run("IsAdmin Test", func(t *testing.T) {
		tokenString, _ := jwt.Keygen(AdminData["email"].(string), AdminData["password"].(string), st.Cfg.JwtSecret)

		IsAdminReq := &pbAuth.IsAdminRequest{
			JwtToken: tokenString,
		}

		response, err := st.AuthClient.IsAdmin(ctx, IsAdminReq)
		if err != nil {
			log.Fatalf("Ошибка при вызове функции IsAdmin: %v", err)
		}

		assert.Equal(t, true, response.IsAdmin)
	})

	t.Run("IsNotAdmin Test", func(t *testing.T) {
		tokenString, _ := jwt.Keygen(UserData["email"].(string), UserData["password"].(string), st.Cfg.JwtSecret)

		IsAdminReq := &pbAuth.IsAdminRequest{
			JwtToken: tokenString,
		}

		response, err := st.AuthClient.IsAdmin(ctx, IsAdminReq)
		if err != nil {
			log.Fatalf("Ошибка при вызове функции IsAdmin: %v", err)
		}

		assert.Equal(t, true, response.IsAdmin)
	})
}
