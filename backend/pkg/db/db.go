package db

import (
	"fmt"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var Conn *sqlx.DB

func Init(DBUser string, DBPassword string, DBHost string, DBName string) error {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s", DBUser, DBPassword, DBHost, DBName)

	var err error
	Conn, err = sqlx.Open("pgx", connStr)
	if err != nil {
		return err
	}

	// Настройка пула соединений
	Conn.SetMaxOpenConns(100)
	Conn.SetMaxIdleConns(50)
	Conn.SetConnMaxLifetime(time.Hour)

	// Проверка подключения к базе данных
	err = Conn.Ping()
	if err != nil {
		return err
	}

	return nil
}

// принимает map - значения, которые нужно внести в БД и string - таблицу, в которую будем вносить значения
func PullData(table string, data map[string]map[string]interface{}) error {
	for _, keyData := range data {
		var (
			columns []string
			values  []string
		)

		for key, value := range keyData {
			columns = append(columns, key)
			values = append(values, fmt.Sprintf("%s", value))
		}
		cmdStr := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (?)`, table, strings.Join(columns, ", "))
		query, args, err := sqlx.In(cmdStr, values)

		if err != nil {
			return err
		}

		query = Conn.Rebind(query)
		_, err = Conn.Query(query, args...)
		if err != nil {
			return err
		}
	}

	return nil
}
