package news_tests

import (
	"log"
	pbNews "moscowhack/gen/go/news"
	"moscowhack/tests/suite"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestGetNews(t *testing.T) {
	ctx, st := suite.New(t)
	var (
		Id         = gofakeit.Int32()
		Title      = gofakeit.Book().Title
		Text       = "Test text for text news"
		Datetime   = "2024-02-27 19:38:05"
		Categories = "1,2"
	)

	req := &pbNews.NewsRequest{
		Id:         Id,
		Categories: Categories,
	}

	t.Run("GetNews Test", func(t *testing.T) {

		response, err := st.NewsClient.GetNews(ctx, req)
		if err != nil {
			log.Fatalf("Error when calling Registration: %v", err)
		}

		assert.NotNil(t, response.News)
	})
}
