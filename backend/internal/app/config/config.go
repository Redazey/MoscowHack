package config

import (
	"moscowhack/pkg/logger"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Configuration struct {
	Port        string `env:"ADDRESS" envDefault:"4000"`
	LoggerLevel string `env:"LOGGER_LEVEL" envDefault:"debug"`
	IsDebug     bool   `env:"DEBUG" envDefault:"false"`
	DB          DB
	Redis       Redis
	Cache       Cache
}

type Cache struct {
	CacheInterval string `env:"CACHE_CREATE_INTERVAL" envDefault:"15"`
	CacheEXTime   string `env:"CacheEXTime" envDefault:"15"`
}

type DB struct {
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`
	DBHost     string `env:"DB_HOST,required"`
}

type Redis struct {
	RedisAddr     string `env:"REDIS_ADDR,required"`
	RedisPort     string `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword string `env:"REDIS_PASSWORD,required"`
}

/*
Структура файла конфигурации

	-------GENERAL------
	Port          string
	LoggerLevel   string
	IsDebug       bool
	---------DB---------
	DBUser        string
	DBPassword    string
	DBName        string
	DBHost        string
	-------REDIS--------
	RedisAddr     string
	RedisPort     string
	RedisPassword string
	-------CACHE--------
	CacheInterval string
	CacheEXTime   int
*/
func NewConfig(files ...string) (*Configuration, error) {
	err := godotenv.Load(files...)
	if err != nil {
		logger.Error("Файл .env не найден", zap.Error(err), zap.Any("Все файлы в директории", files))
	}

	cfg := Configuration{}
	err = env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
