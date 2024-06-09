package news

import (
	"context"
	"fmt"
	"log"
	"moscowhack/gen/go/news"
	"moscowhack/pkg/cache"
	"moscowhack/pkg/db"
	"strconv"
	"strings"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) GetNewsService(ctx context.Context) (map[string]*news.NewsItem, error) {
	// Initialize newsSlice
	newsMap := make(map[string]map[string]interface{})

	newsCheck, err := cache.IsExistInCache("news")
	if newsCheck && err == nil {
		newsMap, err = cache.ReadCache("news")
		if err != nil {
			return nil, err
		}

		newsContentMap := createNewsContentMap(newsMap)

		return newsContentMap, nil
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
		var id int
		var title, text, datetime, categories string
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

		newsMap[fmt.Sprint(id)] = map[string]interface{}{
			"id":         fmt.Sprint(id),
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

	return newsContentMap, nil
}

func (s *Service) GetNewsByIdService(ctx context.Context, id int32) (map[string]*news.NewsItem, error) {
	// Initialize newsSlice
	newsMap := make(map[string]map[string]interface{})

	newsCheck, err := cache.IsExistInCache("news_" + fmt.Sprint(id))
	if newsCheck && err == nil {
		newsMap, err = cache.ReadCache("news_" + fmt.Sprint(id))
		if err != nil {
			return nil, err
		}

		newsContentMap := createNewsContentMap(newsMap)

		return newsContentMap, nil
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
		var id int
		var title, text, datetime, categories string
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

		newsMap[fmt.Sprint(id)] = map[string]interface{}{
			"id":         fmt.Sprint(id),
			"title":      title,
			"text":       text,
			"datetime":   datetime,
			"categories": categories,
		}
	}

	err = cache.SaveCache("news_"+fmt.Sprint(id), newsMap)
	if err != nil {
		return nil, err
	}

	newsContentMap := createNewsContentMap(newsMap)

	return newsContentMap, nil
}

func (s *Service) GetNewsByCategoryService(ctx context.Context, categoryId string) (map[string]*news.NewsItem, error) {
	// Initialize newsSlice
	newsMap := make(map[string]map[string]interface{})

	newsCheck, err := cache.IsExistInCache("news_category_" + categoryId)
	if newsCheck && err == nil {
		newsMap, err = cache.ReadCache("news_category_" + categoryId)
		if err != nil {
			return nil, err
		}

		newsContentMap := createNewsContentMap(newsMap)

		return newsContentMap, nil
	}

	// Данных нет
	rows, err := db.Conn.Query(`SELECT n.*, c."name"
    FROM "news" n
    JOIN "categoriesNews" cn ON n."id" = cn."newsID"
    JOIN "categories" c ON cn."categoryID" = c."id"
    WHERE cn."categoryID" IN (` + categoryId + `)`)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title, text, datetime, categories string
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

		newsMap[fmt.Sprint(id)] = map[string]interface{}{
			"id":         fmt.Sprint(id),
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

	return newsContentMap, nil
}

func (s *Service) AddNewsService(ctx context.Context, title string, text string, datetime string, categories string) (int32, error) {
	t, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err != nil {
		return 0, err
	}

	// Начало транзакции
	tx, err := db.Conn.Beginx()
	if err != nil {
		return 0, err
	}

	// Добавление новости в таблицу news
	var newsID int
	err = tx.QueryRowx("INSERT INTO news (title, text, datetime) VALUES ($1, $2, $3) RETURNING id", title, text, t).Scan(&newsID)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return 0, err
		}
		return 0, err
	}

	categoriesSlice := strings.Split(categories, ",")

	// Привязка категорий к новости
	for _, category := range categoriesSlice {
		var categoryID int
		categoryID, err = strconv.Atoi(category)
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				return 0, err
			}
			return 0, err
		}

		_, err = tx.Exec("INSERT INTO \"categoriesNews\" (\"newsID\", \"categoryID\") VALUES ($1, $2)", newsID, categoryID)
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				return 0, err
			}
			return 0, err
		}
	}

	// Фиксация транзакции
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int32(newsID), nil
}

func (s *Service) DelNewsService(ctx context.Context, newsID int32) error {
	// Начало транзакции
	tx, err := db.Conn.Beginx()
	if err != nil {
		return err
	}

	// Удаление связей новости с категориями
	_, err = tx.Exec("DELETE FROM \"categoriesNews\" WHERE \"newsID\" = $1", newsID)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errRollback
		}
		return err
	}

	// Удаление самой новости
	_, err = tx.Exec("DELETE FROM news WHERE id = $1", newsID)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errRollback
		}
		return err
	}

	// Фиксация транзакции
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func createNewsContentMap(newsMap map[string]map[string]interface{}) map[string]*news.NewsItem {
	fmt.Println(newsMap)

	newsContentMap := make(map[string]*news.NewsItem)
	for _, data := range newsMap {
		fmt.Println(data["id"].(string))
		id, err := strconv.ParseInt(strings.TrimSpace(data["id"].(string)), 10, 32)
		if err != nil {
			log.Fatalf("Error converting string to int32: %v", err)
		}

		newsContent := &news.NewsItem{
			Id:         int32(id),
			Title:      data["title"].(string),
			Text:       data["text"].(string),
			Datetime:   data["datetime"].(string),
			Categories: data["categories"].(string),
		}
		newsContentMap[data["id"].(string)] = newsContent
	}
	return newsContentMap
}
