package main

import (
	"moscowhack/internal/app/config"
	"moscowhack/internal/pkg/app"
	"moscowhack/pkg/db"
	"moscowhack/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	// Подгрузка конфигурации
	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// Инициализация логгепа
	logger.Init(configuration.LoggerLevel)

	// Подключение к БД
	err = db.Init(configuration.DBUser, configuration.DBPassword, configuration.DBHost, configuration.DBName)
	if err != nil {
		logger.Fatal("Ошибка при подключении к БД", zap.Error(err))
	}

	_, err = app.New()
	if err != nil {
		return
	}
}
