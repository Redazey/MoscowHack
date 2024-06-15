package cache

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx         = context.Background()
	Rdb         *redis.Client
	CacheEXTime = 15
)

func Init(Addr string, Username string, Password string, DB int) error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Username: Username,
		Password: Password,
		DB:       DB,
	})

	err := Rdb.Ping(Ctx).Err()
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
func ConvertMap(inputMap map[string]interface{}, columns ...string) (map[string]map[string]interface{}, string) {
	var mainKey string

	hash := md5.Sum([]byte(strings.Join(columns, "")))
	mainKey = hex.EncodeToString(hash[:])

	outputMap := make(map[string]map[string]interface{})
	outputMap[mainKey] = make(map[string]interface{})

	for key, value := range inputMap {
		outputMap[mainKey][key] = value
	}

	return outputMap, mainKey
}

// Функция для получения хэш значения из вводимых полей
func GetHash(columns ...string) string {
	var hashKey string

	hash := md5.Sum([]byte(strings.Join(columns, "")))
	hashKey = hex.EncodeToString(hash[:])

	return hashKey
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
функция для проверки существования таблицы в кэше

принимает:

	table - имя таблицы

возвращает:

	bool - true, если таблица существует, иначе false
	error - ошибка, если возникла
*/
func IsExistInCache(table string) (bool, error) {
	exists, err := Rdb.Exists(Ctx, table).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
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

		err = Rdb.HSet(Ctx, table, key, jsonMap).Err()
		if err != nil {
			return err
		}

		// Устанавливаем время жизни ключа
		err = Rdb.Expire(Ctx, key, time.Minute*time.Duration(CacheEXTime)).Err()
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
	result, err := Rdb.HGetAll(Ctx, table).Result()
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
	keys, err := Rdb.HKeys(Ctx, table).Result()
	if err != nil {
		return err
	}

	// удаляем все протухшие ключи из Redis
	for _, key := range keys {
		// Получаем время до истечения срока действия ключа
		ttl := Rdb.TTL(Ctx, key).Val()

		if ttl <= 0 {
			// Если TTL < 0, значит ключ уже истек и можно его удалить
			err := Rdb.Del(Ctx, key).Err()
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
func ClearCache(Rdb *redis.Client) error {
	// Удаление всего кэша из Redis
	err := Rdb.FlushAll(Ctx).Err()
	if err != nil {
		return err
	}
	return nil
}
