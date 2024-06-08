package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"moscowhack/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationsPath, migrationsTable string

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Ошибка при инициализации конфига: %s", err)
	}

	flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	// Путь до папки с миграциями.
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	// Таблица, в которой будет храниться информация о миграциях. Она нужна
	// для того, чтобы понимать, какие миграции уже применены, а какие нет.
	// Дефолтное значение - 'migrations'.
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.Parse() // Выполняем парсинг флагов

	// Валидация параметров
	if storagePath == "" {
		panic("storage-path is required")
	}
	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s", cfg.DB.DBUser, cfg.DB.DBPassword, cfg.DB.DBHost, cfg.DB.DBName)

	m, err := migrate.New(
		"file://"+migrationsPath,
		connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer m.Close()

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("No migrations to apply")
		} else {
			log.Fatal(err)
		}
	} else {
		log.Println("Migrations applied successfully")
	}
}
