package news

import (
	"context"
	"encoding/json"
	"fmt"
	"moscowhack/gen/go/news"
	"moscowhack/pkg/cache"
	"moscowhack/pkg/db"
	"strings"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) GetNewsService(ctx context.Context) (*news.NewsItem, error) {
	// Initialize newsSlice
	newsMap := make(map[string]map[string]interface{})

	newsCheck, err := cache.IsExistInCache("news")
	fmt.Println(newsCheck)
	if newsCheck && err == nil {
		newsMap, err = cache.ReadCache("news")
		if err != nil {
			return nil, err
		}

		newsContentMap := make(map[string]*news.NewsContent)
		for id, data := range newsMap {
			newsContent := &news.NewsContent{
				Id:         id,
				Title:      data["title"].(string),
				Text:       data["text"].(string),
				Datetime:   data["datetime"].(string),
				Categories: data["categories"].(string),
			}
			newsContentMap[id] = newsContent
		}

		newsItem := news.NewsItem{NewsItem: newsContentMap}

		return &newsItem, nil
	}

	// Данных нет
	rows, err := db.Conn.Query(`
		SELECT n.id, n.title, n.text, n.datetime, string_agg(c.name, ',') AS categories
		FROM news n
		JOIN "categoriesNews" cn ON n.id = cn."newsID"
		JOIN categories c ON cn."categoryID" = c.id
		GROUP BY n.id;
	`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, title, text, datetime, categories string
		err := rows.Scan(
			&id,
			&title,
			&text,
			&datetime,
			&categories,
		)
		if err != nil {
			return nil, err
		}

		newsMap[id] = map[string]interface{}{
			"title":      title,
			"text":       text,
			"datetime":   datetime,
			"categories": categories,
		}
	}

	err = cache.SaveCache("news", newsMap)
	if err != nil {
		return nil, err
	}

	newsContentMap := make(map[string]*news.NewsContent)
	for id, data := range newsMap {
		newsContent := &news.NewsContent{
			Id:         id,
			Title:      data["title"].(string),
			Text:       data["text"].(string),
			Datetime:   data["datetime"].(string),
			Categories: data["categories"].(string),
		}
		newsContentMap[id] = newsContent
	}

	newsItem := news.NewsItem{NewsItem: newsContentMap}

	return &newsItem, nil
}

func (s *Service) GetNewsByIdService(ctx context.Context, id int) (*news.NewsItem, error) {
	// Initialize newsSlice
	var newsSlice map[string]*news.NewsContent

	// Try to get news from Redis cache
	newsData, err := cache.Rdb.Get(cache.Ctx, "News_"+fmt.Sprint(id)).Result()
	if err == nil && newsData != "" {
		// If news is found in cache, unmarshal and return
		if err := json.Unmarshal([]byte(newsData), &newsSlice); err == nil {
			newsItem := news.NewsItem{NewsItem: newsSlice}

			return &newsItem, nil
		}
	}

	// Data not in Redis, get from database
	rows, err := db.Conn.Queryx(`
		SELECT n.id, n.title, n.text, n.datetime, array_agg(c.name) AS categories
		FROM news n
		JOIN "categoriesNews" cn ON n.id = cn."newsID"
		JOIN categories c ON cn."categoryID" = c.id
		GROUP BY n.id WHERE n.id = $1`, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var newsContent news.NewsContent
		var categories []string
		err := rows.Scan(
			&newsContent.Id,
			&newsContent.Title,
			&newsContent.Text,
			&newsContent.Datetime,
			&categories,
		)
		if err != nil {
			return nil, err
		}
		newsContent.Categories = strings.Join(categories, ",")
		// Store news content in newsSlice
		newsSlice[newsContent.Id] = &newsContent
	}

	// Save data to Redis
	newsJSON, err := json.Marshal(newsSlice)
	if err != nil {
		return nil, err
	}
	err = cache.Rdb.Set(cache.Ctx, "News_"+fmt.Sprint(id), newsJSON, 1*time.Minute).Err()
	if err != nil {
		return nil, err
	}

	newsItem := news.NewsItem{NewsItem: newsSlice}

	return &newsItem, nil
}
