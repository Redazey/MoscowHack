package db

import (
	"fmt"
	"moscowhack/pkg/db"
)

// принимает таблицу как string и возвращает таблицу в виде map
func FetchUserData(username string) (map[string]string, error) {
	rows, err := db.Conn.Query(`SELECT username, password, roleid 
						   FROM users
						   WHERE username = $1`, username)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Получение информации о столбцах
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Инициализация именованного массива, который содержит структуру для сканирования
	values := make([]interface{}, len(columns))
	for i := range columns {
		values[i] = new(interface{})
	}

	// Инициализация мапы для хранения данных
	dataMap := make(map[string]string)

	// Чтение данных из таблицы и добавление их в map
	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		for i, colName := range columns {
			val := *(values[i].(*interface{}))
			if val == nil {
				dataMap[colName] = ""
			} else {
				dataMap[colName] = fmt.Sprintf("%v", val)
			}
		}
	}

	return dataMap, nil
}
