package app

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"moscowhack/config"
	pbAuth "moscowhack/gen/go/auth"
	pbNews "moscowhack/gen/go/news"
	"moscowhack/internal/app/endpoint/grpcAuth"
	"moscowhack/internal/app/endpoint/grpcNews"
	"moscowhack/internal/app/lib/cacher"
	"moscowhack/internal/app/service/auth"
	"moscowhack/internal/app/service/news"
	"moscowhack/pkg/cache"
	"moscowhack/pkg/db"
	"moscowhack/pkg/logger"
	"net"
	"net/http"
	_ "net/http/pprof"
)

type App struct {
	auth *auth.Service
	news *news.Service

	server *grpc.Server
}

func New() (*App, error) {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

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
	a.auth = auth.New()
	a.news = news.New()

	// регистрируем эндпоинты
	serviceAuth := &grpcAuth.Endpoint{
		Auth: a.auth,
	}
	pbAuth.RegisterAuthServiceServer(a.server, serviceAuth)

	serviceNews := &grpcNews.Endpoint{
		News: a.news,
	}
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

	return nil
}

func (a *App) Stop() {
	logger.Info("закрытие gRPC сервера")

	a.server.GracefulStop()
}
