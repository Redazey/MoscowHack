package auth_tests

import (
	"fmt"
	"moscowhack/pkg/db"

	"github.com/brianvoe/gofakeit/v6"
)

// Добавляет в бд случайного юзера с данным roleId и возвращает его данные
func MockUser(roleId int) (map[string]interface{}, error) {
	UserData := map[string]interface{}{
		"surname":    gofakeit.Name(),
		"name":       gofakeit.Name(),
		"patronymic": gofakeit.Name(),
		"birthdate":  gofakeit.Date().String(),
		"photourl":   "testimg",
		"push":       gofakeit.Bool(),
		"email":      gofakeit.Email(),
		"password":   gofakeit.Password(true, true, true, true, false, 10),
		"roleId":     roleId,
	}

	// Выполнение SQL запроса
	_, err := db.Conn.Exec(
		`INSERT IGNORE INTO users (surname, name, patronymic, birthdate, photourl, email, password, push) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		UserData["surname"], UserData["name"], UserData["patronymic"],
		UserData["birthdate"], UserData["photourl"], UserData["push"],
		UserData["email"], UserData["password"], UserData["roleId"],
	)
	if err != nil {
		return nil, err
	}

	// Добавляем роли
	// Выполнение SQL запроса
	_, err = db.Conn.Exec(
		`INSERT IGNORE INTO roles (name) VALUES ('admin'), ('user')`,
	)
	if err != nil {
		return nil, err
	}

	// Выполнение SQL скрипта для связывания пользователя с ролями
	_, err = db.Conn.Exec(`
		WITH new_user AS (
			SELECT id FROM users WHERE email = $1
		)
		INSERT IGNORE INTO userroles (userid, roleid) 
		SELECT new_user.id, roles.id 
		FROM new_user, roles 
		WHERE roles.name IN ('admin', 'user')`,
		UserData["email"],
	)

	if err != nil {
		return nil, err
	}

	fmt.Println("Данные о пользователе успешно добавлены")

	return UserData, nil
}

// Очищает таблицу в бд
func ClearTable(tables []string) error {
	for i := range tables {
		// SQL запрос для очистки
		sqlStatement := fmt.Sprintf(`DELETE FROM %s`, tables[i])

		// Выполнение SQL запроса
		_, err := db.Conn.Exec(sqlStatement)
		if err != nil {
			return err
		}
	}

	fmt.Println("Таблицы очищена")

	return nil
}
