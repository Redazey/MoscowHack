package news

import (
	"context"
	"errors"
	"moscowhack/gen/go/news"
	"reflect"
	"testing"
)

// Mock cache
type mockCache struct {
	isExistInCacheFunc func(key string) (bool, error)
	readCacheFunc      func(key string) (map[string]map[string]interface{}, error)
	saveCacheFunc      func(key string, data map[string]map[string]interface{}) error
}

func (m *mockCache) IsExistInCache(key string) (bool, error) {
	if m.isExistInCacheFunc != nil {
		return m.isExistInCacheFunc(key)
	}
	return false, nil
}

func (m *mockCache) ReadCache(key string) (map[string]map[string]interface{}, error) {
	if m.readCacheFunc != nil {
		return m.readCacheFunc(key)
	}
	return nil, errors.New("read cache function not implemented")
}

func (m *mockCache) SaveCache(key string, data map[string]map[string]interface{}) error {
	if m.saveCacheFunc != nil {
		return m.saveCacheFunc(key, data)
	}
	return errors.New("save cache function not implemented")
}

// Mock DB Rows
type mockRows struct {
	rowsData   [][]interface{}
	currentRow int
}

func (m *mockRows) Next() bool {
	m.currentRow++
	return m.currentRow < len(m.rowsData)
}

func (m *mockRows) Scan(dest ...interface{}) error {
	for i, d := range dest {
		*(d.(*string)) = m.rowsData[m.currentRow][i].(string)
	}
	return nil
}

func (m *mockRows) Close() error {
	return nil
}

// Mock DB
type mockDB struct {
	queryFunc func(query string, args ...interface{}) (dbRows, error)
}

func (m *mockDB) Query(query string, args ...interface{}) (dbRows, error) {
	if m.queryFunc != nil {
		return m.queryFunc(query, args...)
	}
	return nil, errors.New("query function not implemented")
}

// Mock DB Rows Interface
type dbRows interface {
	Next() bool
	Scan(dest ...interface{}) error
	Close() error
}

func TestService_GetNewsService(t *testing.T) {
	service := &Service{}

	// Setup mocks
	service.cache = &mockCache{
		isExistInCacheFunc: func(key string) (bool, error) {
			return true, nil
		},
		readCacheFunc: func(key string) (map[string]map[string]interface{}, error) {
			newsMap := make(map[string]map[string]interface{})
			newsMap["1"] = map[string]interface{}{
				"id":         "1",
				"title":      "Test News",
				"text":       "This is a test news article",
				"datetime":   "2024-06-08",
				"categories": "test,unit",
			}
			return newsMap, nil
		},
	}

	ctx := context.Background()

	newsItem, err := service.GetNewsService(ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := &news.NewsItem{
		NewsItem: map[string]*news.NewsContent{
			"1": {
				Id:         "1",
				Title:      "Test News",
				Text:       "This is a test news article",
				Datetime:   "2024-06-08",
				Categories: "test,unit",
			},
		},
	}

	if !reflect.DeepEqual(newsItem, expected) {
		t.Errorf("Expected %+v, got %+v", expected, newsItem)
	}
}

func TestService_GetNewsByIdService(t *testing.T) {
	service := &Service{}

	// Setup mocks
	service.cache = &mockCache{
		isExistInCacheFunc: func(key string) (bool, error) {
			return true, nil
		},
		readCacheFunc: func(key string) (map[string]map[string]interface{}, error) {
			newsMap := make(map[string]map[string]interface{})
			newsMap["1"] = map[string]interface{}{
				"id":         "1",
				"title":      "Test News",
				"text":       "This is a test news article",
				"datetime":   "2024-06-08",
				"categories": "test,unit",
			}
			return newsMap, nil
		},
	}

	ctx := context.Background()

	newsItem, err := service.GetNewsByIdService(ctx, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := &news.NewsItem{
		NewsItem: map[string]*news.NewsContent{
			"1": {
				Id:         "1",
				Title:      "Test News",
				Text:       "This is a test news article",
				Datetime:   "2024-06-08",
				Categories: "test,unit",
			},
		},
	}

	if !reflect.DeepEqual(newsItem, expected) {
		t.Errorf("Expected %+v, got %+v", expected, newsItem)
	}
}

func TestService_GetNewsByCategoryService(t *testing.T) {
	service := &Service{}

	// Setup mocks
	service.cache = &mockCache{
		isExistInCacheFunc: func(key string) (bool, error) {
			return true, nil
		},
		readCacheFunc: func(key string) (map[string]map[string]interface{}, error) {
			newsMap := make(map[string]map[string]interface{})
			newsMap["1"] = map[string]interface{}{
				"id":         "1",
				"title":      "Test News",
				"text":       "This is a test news article",
				"datetime":   "2024-06-08",
				"categories": "test,unit",
			}
			return newsMap, nil
		},
	}

	ctx := context.Background()

	newsItem, err := service.GetNewsByCategoryService(ctx, []string{"test"})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := &news.NewsItem{
		NewsItem: map[string]*news.NewsContent{
			"1": {
				Id:         "1",
				Title:      "Test News",
				Text:       "This is a test news article",
				Datetime:   "2024-06-08",
				Categories: "test,unit",
			},
		},
	}

	if !reflect.DeepEqual(newsItem, expected) {
		t.Errorf("Expected %+v, got %+v", expected, newsItem)
	}
}

func TestService_GetNewsService_Error(t *testing.T) {
	service := &Service{}

	// Setup mocks
	service.cache = &mockCache{
		readCacheFunc: func(key string) (map[string]map[string]interface{}, error) {
			return nil, errors.New("cache error")
		},
	}

	ctx := context.Background()

	_, err := service.GetNewsService(ctx)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestService_GetNewsByIdService_Error(t *testing.T) {
	service := &Service{}

	// Setup mocks
	service.cache = &mockCache{
		readCacheFunc: func(key string) (map[string]map[string]interface{}, error) {
			return nil, errors.New("cache error")
		},
	}

	ctx := context.Background()

	_, err := service.GetNewsByIdService(ctx, 1)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestService_GetNewsByCategoryService_Error(t *testing.T) {
	service := &Service{}

	// Setup mocks
	service.cache = &mockCache{
		readCacheFunc: func(key string) (map[string]map[string]interface{}, error) {
			return nil, errors.New("cache error")
		},
	}

	ctx := context.Background()

	_, err := service.GetNewsByCategoryService(ctx, []string{"test"})
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}
