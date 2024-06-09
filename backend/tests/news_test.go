package tests

import (
	"log"
	pbNews "moscowhack/gen/go/news"
	"moscowhack/tests/suite"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestNews(t *testing.T) {
	ctx, st := suite.New(t)
	var (
		Id         = gofakeit.Int32()
		Title      = gofakeit.Book().Title
		Text       = "Test text for text news"
		Datetime   = "2024-02-27 19:38:05"
		Categories = "1,2"
	)

	exceptedNews := pbNews.NewsResponse{
		News: map[string]*pbNews.NewsItem{
			"testnews": {
				Id:         Id,
				Title:      Title,
				Text:       Text,
				Datetime:   Datetime,
				Categories: Categories,
			},
		},
	}

	NewsReq := &pbNews.NewsRequest{
		Id:         Id,
		Title:      Title,
		Text:       Text,
		Datetime:   Datetime,
		Categories: Categories,
	}

	req := &pbNews.NewsRequest{
		Id:         Id,
		Categories: Categories,
	}

	t.Run("AddNews Test", func(t *testing.T) {

		response, err := st.NewsClient.AddNews(ctx, NewsReq)
		if err != nil {
			log.Fatalf("Error when adding news: %v", err)
		}

		assert.Equal(t, Id, response.Id)
	})

	t.Run("GetNews Test", func(t *testing.T) {

		response, err := st.NewsClient.GetNews(ctx, req)
		if err != nil {
			log.Fatalf("Error when calling Registration: %v", err)
		}

		assert.NotNil(t, response.News)
	})
	t.Run("GetNewsById Test", func(t *testing.T) {
		response, err := st.NewsClient.GetNewsById(ctx, req)
		if err != nil {
			log.Fatalf("Error when calling Login: %v", err)
		}

		assert.Equal(t, exceptedNews.News, response.News, "Excepted value: %v, recieved: %v", exceptedNews.News, response.News)
	})
	t.Run("GetNewsByCategory Test", func(t *testing.T) {
		response, err := st.NewsClient.GetNewsByCategory(ctx, req)
		if err != nil {
			log.Fatalf("Error when calling Login: %v", err)
		}

		assert.Equal(t, exceptedNews.News, response.News, "Excepted value: %v, recieved: %v", exceptedNews.News, response.News)
	})
	t.Run("DelNews Test", func(t *testing.T) {
		response, err := st.NewsClient.DelNews(ctx, NewsReq)
		if err != nil {
			log.Fatalf("Error when deleting news: %v", err)
		}

		assert.Equal(t, "", response.Err)
	})
}
