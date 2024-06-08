package news

import (
	"context"
	"fmt"
	"moscowhack/gen/go/news"
	"moscowhack/pkg/cache"
	"moscowhack/pkg/db"
	"strings"
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
	if newsCheck && err == nil {
		newsMap, err = cache.ReadCache("news")
		if err != nil {
			return nil, err
		}

		newsContentMap := createNewsContentMap(newsMap)
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
			"id":         id,
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

	newsContentMap := createNewsContentMap(newsMap)
	newsItem := news.NewsItem{NewsItem: newsContentMap}

	return &newsItem, nil
}

func (s *Service) GetNewsByIdService(ctx context.Context, id int) (*news.NewsItem, error) {
	// Initialize newsSlice
	newsMap := make(map[string]map[string]interface{})

	newsCheck, err := cache.IsExistInCache("news_" + fmt.Sprint(id))
	if newsCheck && err == nil {
		newsMap, err = cache.ReadCache("news_" + fmt.Sprint(id))
		if err != nil {
			return nil, err
		}

		newsContentMap := createNewsContentMap(newsMap)
		newsItem := news.NewsItem{NewsItem: newsContentMap}

		return &newsItem, nil
	}

	// Данных нет
	rows, err := db.Conn.Query(`
		SELECT n.id, n.title, n.text, n.datetime, string_agg(c.name, ',') AS categories
		FROM news n
		JOIN "categoriesNews" cn ON n.id = cn."newsID"
		JOIN categories c ON cn."categoryID" = c.id
		WHERE n.id = $1 GROUP BY n.id;
	`, id)
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
			"id":         id,
			"title":      title,
			"text":       text,
			"datetime":   datetime,
			"categories": categories,
		}
		fmt.Println(newsMap[id])
	}

	err = cache.SaveCache("news_"+fmt.Sprint(id), newsMap)
	if err != nil {
		return nil, err
	}

	newsContentMap := createNewsContentMap(newsMap)
	newsItem := news.NewsItem{NewsItem: newsContentMap}

	return &newsItem, nil
}

func (s *Service) GetNewsByCategoryService(ctx context.Context, categoryId []string) (*news.NewsItem, error) {
	// Initialize newsSlice
	newsMap := make(map[string]map[string]interface{})

	// Преобразование []string в строку с разделителем ","
	categoryIdString := strings.Join(categoryId, ",")

	newsCheck, err := cache.IsExistInCache("news_category_" + categoryIdString)
	if newsCheck && err == nil {
		newsMap, err = cache.ReadCache("news_category_" + categoryIdString)
		if err != nil {
			return nil, err
		}

		newsContentMap := createNewsContentMap(newsMap)
		newsItem := news.NewsItem{NewsItem: newsContentMap}

		return &newsItem, nil
	}

	// Данных нет
	rows, err := db.Conn.Query(`SELECT n.*, c."name"
    FROM "news" n
    JOIN "categoriesNews" cn ON n."id" = cn."newsID"
    JOIN "categories" c ON cn."categoryID" = c."id"
    WHERE cn."categoryID" IN (` + categoryIdString + `)`)

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
			"id":         id,
			"title":      title,
			"text":       text,
			"datetime":   datetime,
			"categories": categories,
		}
	}

	err = cache.SaveCache("news_category_"+fmt.Sprint(categoryId), newsMap)
	if err != nil {
		return nil, err
	}

	newsContentMap := createNewsContentMap(newsMap)
	newsItem := news.NewsItem{NewsItem: newsContentMap}

	return &newsItem, nil
}

func createNewsContentMap(newsMap map[string]map[string]interface{}) map[string]*news.NewsContent {
	newsContentMap := make(map[string]*news.NewsContent)
	for _, data := range newsMap {
		newsContent := &news.NewsContent{
			Id:         data["id"].(string),
			Title:      data["title"].(string),
			Text:       data["text"].(string),
			Datetime:   data["datetime"].(string),
			Categories: data["categories"].(string),
		}
		newsContentMap[data["id"].(string)] = newsContent
	}
	return newsContentMap
}
