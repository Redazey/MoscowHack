package config

import (
	"kinogo/pkg/logger"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Configuration struct {
	Port          string `env:"ADDRESS" envDefault:"4000"`
	LoggerLevel   string `env:"LOGGER_LEVEL" envDefault:"debug"`
	DBUser        string `env:"DB_USER,required"`
	DBPassword    string `env:"DB_PASSWORD,required"`
	DBName        string `env:"DB_NAME,required"`
	DBHost        string `env:"DB_HOST,required"`
	RedisAddr     string `env:"REDIS_ADDR,required"`
	RedisPort     string `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword string `env:"REDIS_PASSWORD,required"`
	IsDebug       bool   `env:"DEBUG" envDefault:"false"`
}

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
