package news

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "moscowhack/gen/go/news"
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

func (s *Service) GetNewsService() (map[string]*pb.GetNewsItem, error) {
	// Initialize newsSlice
	newsMap, err := getNewsFromCache("news")
	if err != nil {
		return nil, err
	}
	if newsMap != nil {
		newsContentMap, errMap := convertNewsMap(newsMap)
		if errMap != nil {
			return nil, errMap
		}

		return newsContentMap, nil
	}

	// Данных нет
	query := `SELECT n.id, n.title, n.text, n.datetime, string_agg(c.name, ',') AS categories
		FROM news n
		JOIN "categoriesNews" cn ON n.id = cn."newsID"
		JOIN categories c ON cn."categoryID" = c.id
		GROUP BY n.id;`
	newsMap, err = fetchNewsFromDB(query)
	if err != nil {
		return nil, err
	}

	err = cache.SaveCache("news", newsMap)
	if err != nil {
		return nil, err
	}

	newsContentMap, errMap := convertNewsMap(newsMap)
	if errMap != nil {
		return nil, errMap
	}

	return newsContentMap, nil
}

func (s *Service) GetNewsByIdService(id int32) (*pb.GetNewsByIdResponse, error) {
	// Initialize newsSlice
	newsMap, err := getNewsFromCache("news_" + fmt.Sprint(id))
	if err != nil {
		return nil, err
	}
	if newsMap != nil {
		newsContent := &pb.GetNewsByIdResponse{
			Title:      newsMap[fmt.Sprint(id)]["title"].(string),
			Text:       newsMap[fmt.Sprint(id)]["text"].(string),
			Datetime:   newsMap[fmt.Sprint(id)]["datetime"].(string),
			Categories: newsMap[fmt.Sprint(id)]["categories"].(string),
		}

		return newsContent, nil
	}

	// Данных нет
	query := `SELECT n.id, n.title, n.text, n.datetime, string_agg(c.name, ',') AS categories
		FROM news n
		JOIN "categoriesNews" cn ON n.id = cn."newsID"
		JOIN categories c ON cn."categoryID" = c.id
		WHERE n.id = $1 GROUP BY n.id;`
	newsMap, err = fetchNewsFromDB(query, id)
	if err != nil {
		return nil, err
	}

	err = cache.SaveCache("news_"+fmt.Sprint(id), newsMap)
	if err != nil {
		return nil, err
	}

	newsContent := &pb.GetNewsByIdResponse{
		Title:      newsMap[fmt.Sprint(id)]["title"].(string),
		Text:       newsMap[fmt.Sprint(id)]["text"].(string),
		Datetime:   newsMap[fmt.Sprint(id)]["datetime"].(string),
		Categories: newsMap[fmt.Sprint(id)]["categories"].(string),
	}

	return newsContent, nil
}

func (s *Service) GetNewsByCategoryService(categoryId string) (map[string]*pb.GetNewsItem, error) {
	// Initialize newsSlice
	newsMap, err := getNewsFromCache("news_category_" + categoryId)
	if err != nil {
		return nil, err
	}
	if newsMap != nil {
		newsContentMap, errMap := convertNewsMap(newsMap)
		if errMap != nil {
			return nil, errMap
		}

		return newsContentMap, nil
	}

	// Данных нет
	query := `SELECT n.*, c."name"
    FROM "news" n
    JOIN "categoriesNews" cn ON n."id" = cn."newsID"
    JOIN "categories" c ON cn."categoryID" = c."id"
    WHERE cn."categoryID" IN (` + categoryId + `)`
	newsMap, err = fetchNewsFromDB(query)
	if err != nil {
		return nil, err
	}

	err = cache.SaveCache("news_category_"+fmt.Sprint(categoryId), newsMap)
	if err != nil {
		return nil, err
	}

	newsContentMap, errMap := convertNewsMap(newsMap)
	if errMap != nil {
		return nil, errMap
	}

	return newsContentMap, nil
}

func (s *Service) AddNewsService(title string, text string, datetime string, categories string) (int32, error) {
	t, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err != nil {
		return 0, err
	}

	// Начало транзакции
	tx, err := db.Conn.Beginx()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// Добавление новости в таблицу news
	var newsID int32
	err = tx.QueryRowx("INSERT INTO news (title, text, datetime) VALUES ($1, $2, $3) RETURNING id", title, text, t).Scan(&newsID)
	if err != nil {
		return 0, err
	}

	categoriesSlice := strings.Split(categories, ",")

	// Привязка категорий к новости
	for _, category := range categoriesSlice {
		var categoryID int
		categoryID, err = strconv.Atoi(category)
		if err != nil {
			return 0, err
		}

		_, err = tx.Exec("INSERT INTO \"categoriesNews\" (\"newsID\", \"categoryID\") VALUES ($1, $2)", newsID, categoryID)
		if err != nil {
			return 0, err
		}
	}

	// Фиксация транзакции
	if err = tx.Commit(); err != nil {
		return 0, err
	}

	// Очистка кеша
	if err = cache.DeleteCache("news"); err != nil {
		return 0, err
	}
	if err = cache.DeleteCacheByPattern("news_category_*"); err != nil {
		return 0, err
	}

	return newsID, nil
}

func (s *Service) DelNewsService(newsID int32) error {
	// Начало транзакции
	tx, err := db.Conn.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Удаление связей новости с категориями
	_, err = tx.Exec("DELETE FROM \"categoriesNews\" WHERE \"newsID\" = $1", newsID)
	if err != nil {
		return err
	}

	// Удаление самой новости
	_, err = tx.Exec("DELETE FROM news WHERE id = $1", newsID)
	if err != nil {
		return err
	}

	// Фиксация транзакции
	if err = tx.Commit(); err != nil {
		return err
	}

	// Очистка кеша
	if err = cache.DeleteCache("news"); err != nil {
		return err
	}
	if err = cache.DeleteCacheByPattern("news_category_*"); err != nil {
		return err
	}
	if err = cache.DeleteCache("news_" + fmt.Sprint(newsID)); err != nil {
		return err
	}

	return nil
}

func getNewsFromCache(cacheKey string) (map[string]map[string]interface{}, error) {
	if exists, err := cache.IsExistInCache(cacheKey); err != nil {
		return nil, err
	} else if !exists {
		return nil, nil
	}

	newsMap, err := cache.ReadCache(cacheKey)
	if err != nil {
		return nil, err
	}

	if _, notFound := newsMap["notFound"]; notFound {
		return nil, status.Error(codes.NotFound, "нет значений в БД")
	}

	return newsMap, nil
}

func convertNewsMap(newsMap map[string]map[string]interface{}) (map[string]*pb.GetNewsItem, error) {
	newsContentMap := make(map[string]*pb.GetNewsItem)
	for _, data := range newsMap {
		id, err := strconv.ParseInt(strings.TrimSpace(data["id"].(string)), 10, 32)
		if err != nil {
			return nil, status.Error(codes.Internal, error.Error(err))
		}

		newsContent := &pb.GetNewsItem{
			Id:         int32(id),
			Title:      data["title"].(string),
			Text:       data["text"].(string),
			Datetime:   data["datetime"].(string),
			Categories: data["categories"].(string),
		}
		newsContentMap[data["id"].(string)] = newsContent
	}

	return newsContentMap, nil
}

func fetchNewsFromDB(query string, args ...interface{}) (map[string]map[string]interface{}, error) {
	rows, err := db.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	newsMap := make(map[string]map[string]interface{})
	found := false
	for rows.Next() {
		var id int32
		var title, text, datetime, categories string
		err := rows.Scan(&id, &title, &text, &datetime, &categories)
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

		found = true
	}

	if !found {
		newsMap["notFound"] = map[string]interface{}{"message": "нет значений в БД"}
		return nil, status.Error(codes.NotFound, "нет значений в БД")
	}

	return newsMap, nil
}
