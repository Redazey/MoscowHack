package tests

import (
	"log"
	pbAuth "moscowhack/gen/go/auth"
	"moscowhack/internal/app/lib/jwt"
	"moscowhack/tests/suite"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	ctx, st := suite.New(t)

	req := &pbAuth.AuthRequest{
		Username: gofakeit.Name(),
		Password: gofakeit.Password(true, true, true, true, false, 10),
	}
	exceptedKey, _ := jwt.Keygen(req.Username, req.Password)

	t.Run("NewUserRegistration Test", func(t *testing.T) {
		response, err := st.AuthClient.Registration(ctx, req)
		if err != nil {
			log.Fatalf("Error when calling Registration: %v", err)
		}

		assert.Equal(t, exceptedKey, response.Key)
	})
	t.Run("UserLogin Test", func(t *testing.T) {
		response, err := st.AuthClient.Login(ctx, req)
		if err != nil {
			log.Fatalf("Error when calling Login: %v", err)
		}

		assert.Equal(t, exceptedKey, response.Key)
	})
}
