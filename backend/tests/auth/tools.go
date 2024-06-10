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

	// SQL запрос для вставки данных
	sqlStatement := `INSERT INTO Users (surname, name, patronymic, birthdate, photourl, push, email, password, roleId) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	// Выполнение SQL запроса
	_, err := db.Conn.Exec(sqlStatement, UserData["surname"], UserData["name"], UserData["patronymic"],
		UserData["birthdate"], UserData["photourl"], UserData["push"],
		UserData["email"], UserData["password"], UserData["roleId"])
	if err != nil {
		return nil, err
	}

	fmt.Println("Данные успешно добавлены в таблицу Users")

	return UserData, nil
}

// Очищает таблицу в бд
func ClearTable(table string) error {
	// SQL запрос для очистки
	sqlStatement := fmt.Sprintf(`DELETE FROM %s`, table)

	// Выполнение SQL запроса
	_, err := db.Conn.Exec(sqlStatement)
	if err != nil {
		return err
	}

	fmt.Println("Таблица очищена")

	return nil
}
