package app

import (
	"log"
	"moscowhack/internal/app/config"
	"moscowhack/internal/app/endpoint/auth"
	"moscowhack/internal/app/endpoint/news"
	"moscowhack/internal/app/service"
	"moscowhack/internal/app/service/cacher"
	"moscowhack/internal/app/service/getNews"
	"moscowhack/internal/app/service/jwtAuth"
	"moscowhack/pkg/cache"
	"moscowhack/pkg/db"
	"moscowhack/pkg/logger"
	pbAuth "moscowhack/protos/auth"
	pbNews "moscowhack/protos/news"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	auth    *auth.Endpoint
	news    *news.Endpoint
	jwt     *jwtAuth.Service
	getNews *getNews.Service
	service *service.Service
	server  *grpc.Server
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

	a.server = grpc.NewServer()

	// обьявляем сервисы
	a.jwt = jwtAuth.New()
	a.getNews = getNews.New()

	// обьявляем эндпоинты
	a.auth = auth.New(a.jwt)
	a.news = news.New(a.getNews)

	serviceAuth := &auth.AuthServiceServer{}
	pbAuth.RegisterAuthServiceServer(a.server, serviceAuth)

	serviceNews := &news.NewsServiceServer{}
	pbNews.RegisterNewsServiceServer(a.server, serviceNews)

	err = cache.Init(cfg.Redis.RedisAddr+":"+cfg.Redis.RedisPort, cfg.Redis.RedisUsername, cfg.Redis.RedisPassword, cfg.Redis.RedisDBId)
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
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		logger.Fatal("Ошибка при открытии listener: ", zap.Error(err))
	}

	err = a.server.Serve(lis)
	if err != nil {
		logger.Fatal("Ошибка при инициализации сервера: ", zap.Error(err))
		return err
	}

	err = a.server.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) Stop() {
	a.server.Stop()
}
