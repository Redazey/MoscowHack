package cache

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ctx         = context.Background()
	CacheConn   *redis.Client
	CacheEXTime = 15
)

func Init(Addr string, Password string, DB int) error {
	CacheConn = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
		DB:       DB,
	})

	err := CacheConn.Ping(ctx).Err()
	if err != nil {
		return err
	}

	return nil
}

/*
функция для подготовки мапы к записи в кэш
columns - поля, по которым будет сгенерирован md5hash

формат входящей мапы:

	map[string]string = {
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

формат выходящей мапы:

	map[string]map[string]interface{} = {
		"md5hash": {
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		},
	}
	+ md5hash key
*/
func ConvertMap(inputMap http.Header, columns ...string) (map[string]map[string]interface{}, string) {
	var mainKey string

	hash := md5.Sum([]byte(strings.Join(columns, "")))
	mainKey = hex.EncodeToString(hash[:])

	outputMap := make(map[string]map[string]interface{})
	outputMap[mainKey] = make(map[string]interface{})

	for key, value := range inputMap {
		outputMap[mainKey][key] = value[0]
	}

	return outputMap, mainKey
}

/*
поиск данных в кэше по md5 хэш-ключу

запрашиваем поиск по ключу input, и что должно вернуться по ключу output
*/
func IsDataInCache(table string, input string, output string) (interface{}, error) {
	cacheMap, err := ReadCache(table)
	if cacheMap[input] != nil && err == nil {
		return cacheMap[input][output], nil
	} else if err != nil {
		return nil, err
	}

	return nil, nil
}

/*
функция для записи данных в кэш, принимает мапы, после конвертации функцией ConvertMap

Вид входящей мапы:

	map[string]map[string]interface{}{
		"md5hash": {
			"username": "exampleUser",
			"password": "examplePass",
			"roleid":   "exampleRoleid",
		},
	}
*/
func SaveCache(table string, cacheMap map[string]map[string]interface{}) error {
	if cacheMap == nil {
		return nil
	}

	for key, args := range cacheMap {
		// Устанавливаем значение в хэш-таблицу
		jsonMap, err := json.Marshal(args)
		if err != nil {
			return err
		}

		err = CacheConn.HSet(ctx, table, key, jsonMap).Err()
		if err != nil {
			return err
		}

		// Устанавливаем время жизни ключа
		err = CacheConn.Expire(ctx, key, time.Minute*time.Duration(CacheEXTime)).Err()
		if err != nil {
			return err
		}
	}

	// удаляем устаревшие данные
	err := DeleteEX(table)
	if err != nil {
		return err
	}
	return nil
}

/*
Функция для чтения значений по хэш-ключу

возвращает map вида:

	map[string]map[string]interface{} = {
		"md5hash": {
			"username": "exampleUser",
			"password": "examplePass",
			"roleid":   "exampleRoleid",
		},
	}
*/
func ReadCache(table string) (map[string]map[string]interface{}, error) {
	cacheMap := make(map[string]map[string]interface{})

	// Получаем все поля и значения из хэша
	result, err := CacheConn.HGetAll(ctx, table).Result()
	if err != nil {
		return nil, err
	}

	// Преобразуем результат в map[string]interface{}
	for key, value := range result {
		var tempMap map[string]interface{}
		err := json.Unmarshal([]byte(value), &tempMap)
		if err != nil {
			return nil, err
		}
		cacheMap[key] = tempMap
	}

	return cacheMap, nil
}

/*
Функция, которая удаляет все протухшие ключ-значения из выбранной таблицы

автоматически применяется при сохранении кэша при помощи функции SaveCache
*/
func DeleteEX(table string) error {
	keys, err := CacheConn.HKeys(ctx, table).Result()
	if err != nil {
		return err
	}

	// удаляем все протухшие ключи из Redis
	for _, key := range keys {
		// Получаем время до истечения срока действия ключа
		ttl := CacheConn.TTL(ctx, key).Val()

		if ttl <= 0 {
			// Если TTL < 0, значит ключ уже истек и можно его удалить
			err := CacheConn.Del(ctx, key).Err()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

/*
функция для стирания кэша

нужна в основном для дэбага
*/
func ClearCache() {
	// Удаление всего кэша из Redis
	err := CacheConn.FlushAll(ctx).Err()
	if err != nil {
		return
	}
}
