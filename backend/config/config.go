package config

import (
	"log"

	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Port        string        `env:"ADDRESS" envDefault:"4000"`
	LoggerLevel string        `env:"LOGGER_LEVEL" envDefault:"debug"`
	JwtSecret   string        `env:"JWT_SECRET,required"`
	IsDebug     bool          `env:"DEBUG" envDefault:"false"`
	GRPCTimeout time.Duration `env:"GRPC_TIMEOUT" envDefault:"10h"`
	DB          DB
	Redis       Redis
	Cache       Cache
}

type Cache struct {
	CacheInterval string `env:"CACHE_CREATE_INTERVAL" envDefault:"15"`
	CacheEXTime   string `env:"CACHE_EX_TIME" envDefault:"15"`
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
	RedisUsername string `env:"REDIS_USERNAME,required"`
	RedisPassword string `env:"REDIS_PASSWORD,required"`
	RedisDBId     int    `env:"REDIS_DB_ID,required"`
}

/*
Структура файла конфигурации

	-------GENERAL------
	Port          string
	gRPCTimeout   time.Duration
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
	RedisDBid     int
	-------CACHE--------
	CacheInterval string
	CacheEXTime   int
*/
func NewConfig(files ...string) (*Configuration, error) {
	err := godotenv.Load(files...)
	if err != nil {
		log.Fatalf("Файл .env не найден: %s", err)
	}

	cfg := Configuration{}
	err = env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	err = env.Parse(&cfg.Redis)
	if err != nil {
		return nil, err
	}
	err = env.Parse(&cfg.DB)
	if err != nil {
		return nil, err
	}
	err = env.Parse(&cfg.Cache)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
