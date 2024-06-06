package cache_test

import (
	"log"
	"moscowhack/internal/app/config"
	"moscowhack/internal/app/service/cacher"
	"moscowhack/pkg/cache"
	"moscowhack/pkg/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dummydata = map[string]string{
	"username": "testuser",
	"password": "exampass",
	"roleid":   "1",
}

func TestCreateDummyData(t *testing.T) {
	// инициализируем конфиг, логгер и кэш
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Ошибка при попытке спарсить .env файл в структуру: %v", err)
	}

	logger.Init(cfg.LoggerLevel)
	cacher.Init(cfg.Cache.CacheEXTime)

	cache.ClearCache()

	convertedData, hashKey := cache.ConvertMap(dummydata, "username", "password")

	t.Run("FillingWithData Test", func(t *testing.T) {
		err := cache.SaveCache("test", convertedData)
		assert.NoError(t, err, "ошибка при заполнении Redis: %v", err)
	})

	// Тестирование кейса, когда данные есть в кэше
	t.Run("DataInCache Test", func(t *testing.T) {
		result, err := cache.IsDataInCache("test", hashKey, "Password")

		assert.NoError(t, err, "ошибок при выполнении не найдено")
		assert.Equal(t, "exampass", result, "ожидаемое значение - \"exampass\", вышло: %v", result)
	})

	// Тестирование кейса, когда данных нет в кэше
	t.Run("NoDataInCache", func(t *testing.T) {
		result, err := cache.IsDataInCache("test", "ghostuser", "password")
		assert.Nil(t, err, "неожиданная ошибка: %v", err)
		assert.Nil(t, result, "ожидалось nil, вышло: %v", result)
	})

	// Тестирование чтения кэша с определенным ключом
	t.Run("ReadSpecificKey", func(t *testing.T) {
		expectedValue := map[string]interface{}{
			"Content-Type": "application/json",
			"Password":     "exampass",
			"Roleid":       "1",
			"Username":     "testuser",
		}
		result, err := cache.ReadCache("test")

		assert.NoError(t, err, "Expected no error")
		assert.Equal(t, expectedValue, result[hashKey], "Expected value %s", expectedValue)
	})
}
