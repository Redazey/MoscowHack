package getNews

import (
	"encoding/json"
	"fmt"
	"moscowhack/pkg/cache"
	"moscowhack/pkg/db"
	"moscowhack/protos/news"
	"strings"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) GetNewsFromDB(id int) (*news.NewsItem, error) {
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
		GROUP BY n.id`)
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
