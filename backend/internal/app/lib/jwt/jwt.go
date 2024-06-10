package jwt

import (
	"time"

	"moscowhack/internal/app/errorz"

	"github.com/golang-jwt/jwt"
)

func Keygen(username string, pwd string, secretKey []byte) (string, error) {
	// Создаем новый JWT токен
	token := jwt.New(jwt.SigningMethodHS256)

	// Указываем параметры для токена
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["password"] = pwd
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Подписываем токен с помощью указанного секретного ключа
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func TokenAuth(tokenString string, secretKey []byte) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {

		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)

		if time.Now().After(expirationTime) {
			return false, errorz.ErrTokenExpired
		}

	} else {
		return false, errorz.ErrValidation
	}

	return true, nil
}

// С помощью этой функции мы получаем данные о пользователе из jwt токена
func UserDataFromJwt(tokenString string, secretKey []byte) (map[string]string, error) {
	UserData := make(map[string]string)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for key, value := range claims {
			UserData[key] = value.(string)
		}
	} else {
		return nil, errorz.ErrValidation
	}

	return UserData, nil
}
