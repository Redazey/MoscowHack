package app

import (
	"log"
	"moscowhack/config"
	pbAuth "moscowhack/gen/go/auth"
	pbNews "moscowhack/gen/go/news"
	"moscowhack/internal/app/endpoint/grpcAuth"
	"moscowhack/internal/app/endpoint/grpcNews"
	"moscowhack/internal/app/lib/cacher"
	"moscowhack/internal/app/service/auth"
	"moscowhack/internal/app/service/getNews"
	"moscowhack/pkg/cache"
	"moscowhack/pkg/db"
	"moscowhack/pkg/logger"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type App struct {
	grpcAuth *grpcAuth.Endpoint
	grpcNews *grpcNews.Endpoint

	auth    *auth.Service
	getNews *getNews.Service

	server *grpc.Server
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

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			// Логируем информацию о панике с уровнем Error
			logger.Error("Recovered from panic: ", zap.Error(err))

			// Можете либо честно вернуть клиенту содержимое паники
			// Либо ответить - "internal error", если не хотим делиться внутренностями
			return status.Errorf(codes.Internal, err.Error())
		}),
	}

	a.server = grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
	))

	// обьявляем сервисы
	a.auth = auth.New()
	a.getNews = getNews.New()

	// обьявляем эндпоинты
	a.grpcAuth = grpcAuth.New(a.auth)
	a.grpcNews = grpcNews.New(a.getNews)

	serviceAuth := &grpcAuth.Endpoint{}
	pbAuth.RegisterAuthServiceServer(a.server, serviceAuth)

	serviceNews := &grpcNews.Endpoint{}
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
