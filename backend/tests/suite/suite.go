package suite

import (
	"context"
	"github.com/redis/go-redis/v9"
	"moscowhack/config"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pbAuth "moscowhack/gen/go/auth"
	pbNews "moscowhack/gen/go/news"
)

type Suite struct {
	*testing.T
	Cfg        *config.Configuration
	Rdb        *redis.Client
	NewsClient pbNews.NewsServiceClient
	AuthClient pbAuth.AuthServiceClient
}

// New creates new test suite.
func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()   // Функция будет восприниматься как вспомогательная для тестов
	t.Parallel() // Разрешаем параллельный запуск тестов

	// Читаем конфиг из файла
	cfg, err := config.NewConfig("../.env")
	if err != nil {
		t.Fatalf("ошибка при инициализации файла конфигурации: %s", err)
	}

	// Основной родительский контекст
	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPCTimeout)

	// Когда тесты пройдут, закрываем контекст
	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	// Создаем кеш
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.RedisAddr + ":" + cfg.Redis.RedisPort,
		Username: cfg.Redis.RedisUsername,
		Password: cfg.Redis.RedisPassword,
		DB:       cfg.Redis.RedisDBId,
	})
	err = rdb.Ping(ctx).Err()
	if err != nil {
		t.Fatalf("redis connection failed: %v", err)
	}

	// Создаем клиент
	cc, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	// gRPC-клиент сервера Auth
	newsClient := pbNews.NewNewsServiceClient(cc)
	authClient := pbAuth.NewAuthServiceClient(cc)

	return ctx, &Suite{
		T:          t,
		Cfg:        cfg,
		Rdb:        rdb,
		AuthClient: authClient,
		NewsClient: newsClient,
	}
}
