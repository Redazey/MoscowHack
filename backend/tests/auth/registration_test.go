package auth_tests

import (
	"log"
	pbAuth "moscowhack/gen/go/auth"
	"moscowhack/internal/app/lib/jwt"
	"moscowhack/tests/suite"
	"testing"

	gofakeit "github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	ctx, st := suite.New(t)
	ClearTable("users")

	RegReq := &pbAuth.RegistrationRequest{
		Username: gofakeit.Name(),
		Password: gofakeit.Password(true, true, true, true, false, 10),
		RoleId:   1,
	}

	exceptedKey, _ := jwt.Keygen(RegReq.Username, RegReq.Password, st.Cfg.JwtSecret)

	t.Run("NewUserRegistration Test", func(t *testing.T) {
		response, err := st.AuthClient.Registration(ctx, RegReq)
		if err != nil {
			log.Fatalf("Error when calling Registration: %v", err)
		}

		assert.Equal(t, exceptedKey, response.Key)
	})
}
