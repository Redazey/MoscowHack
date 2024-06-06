package app

import (
	"log"
	"moscowhack/internal/app/config"
	"moscowhack/internal/app/endpoint/auth"
	"moscowhack/internal/app/service/cacher"
	"moscowhack/internal/app/service/jwtAuth"
	"moscowhack/pkg/cache"
	"moscowhack/pkg/db"
	"moscowhack/pkg/logger"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type App struct {
	auth *auth.Endpoint
	jwt  *jwtAuth.Service
	echo *echo.Echo
}

func New() (*App, error) {
	// инициализируем конфиг, логгер и кэш
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Ошибка при попытке спарсить .env файл в структуру: %v", err)
	}

	logger.Init(cfg.LoggerLevel)
	cacher.Init(cfg.Cache.CacheEXTime)

	a := &App{}
	// обьявляем сервисы
	a.jwt = jwtAuth.New()

	// обьявляем эндпоинты
	a.auth = auth.New(a.jwt)

	a.echo = echo.New()

	a.echo.GET("/UserLogin", a.auth.UserLogin)
	a.echo.GET("/NewUserRegistration", a.auth.NewUserRegistration)

	err = cache.Init(cfg.Redis.RedisAddr+":"+cfg.Redis.RedisPort, cfg.Redis.RedisPassword, cfg.Redis.RedisDBId)
	if err != nil {
		logger.Error("ошибка при инициализации кэша: ", zap.Error(err))
		return nil, err
	}

	err = db.Init(cfg.DB.DBUser, cfg.DB.DBPassword, cfg.DB.DBHost, cfg.DB.DBName)
	if err != nil {
		logger.Fatal("ошибка при инициализации БД: ", zap.Error(err))
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	err := a.echo.Start(":8080")
	if err != nil {
		logger.Fatal("Ошибка при инициализации сервера: ", zap.Error(err))
		return err
	}

	return nil
}

func (a *App) Stop() error {
	err := a.echo.Close()
	if err != nil {
		logger.Fatal("Ошибка при инициализации сервера: ", zap.Error(err))
		return err
	}

	return nil
}
