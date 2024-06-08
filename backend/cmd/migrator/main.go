package main

import (
	"errors"
	"fmt"
	"log"
	"moscowhack/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Ошибка при инициализации конфига: %s", err)
	}

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable", cfg.DB.DBUser, cfg.DB.DBPassword, cfg.DB.DBHost, cfg.DB.DBName)

	m, err := migrate.New(
		"file://migrations",
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
