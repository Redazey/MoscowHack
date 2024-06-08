package tests

import (
	"log"
	pbNews "moscowhack/gen/go/news"
	"moscowhack/internal/app/lib/jwt"
	"moscowhack/tests/suite"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestNews(t *testing.T) {
	ctx, st := suite.New(t)

	/*
		id:         gofakeit.Int8(),
		title:      gofakeit.BookTitle(),
		text:       "Test text for test news",
		datetime:   gofakeit.Date(),
		categories: gofakeit.Categories(),
	*/

	req := &pbNews.NewsRequest{
		Id:       string(gofakeit.Int8()),
		Category: gofakeit.CarModel(),
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
