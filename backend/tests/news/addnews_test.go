package news_tests

import (
	"log"
	pbNews "moscowhack/gen/go/news"
	"moscowhack/tests/suite"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestAddNews(t *testing.T) {
	ctx, st := suite.New(t)
	var (
		Id         = gofakeit.Int32()
		Title      = gofakeit.Book().Title
		Text       = "Test text for text news"
		Datetime   = "2024-02-27 19:38:05"
		Categories = "1,2"
	)

	NewsReq := &pbNews.NewsRequest{
		Id:         Id,
		Title:      Title,
		Text:       Text,
		Datetime:   Datetime,
		Categories: Categories,
	}

	t.Run("AddNews Test", func(t *testing.T) {

		response, err := st.NewsClient.AddNews(ctx, NewsReq)
		if err != nil {
			log.Fatalf("Error when adding news: %v", err)
		}

		assert.Equal(t, Id, response.Id)
	})
}
