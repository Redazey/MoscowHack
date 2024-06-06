package auth

import (
	"database/sql"
	"moscowhack/internal/app/errorz"
	"moscowhack/internal/app/service/db"
	"moscowhack/pkg/cache"
	"moscowhack/pkg/logger"
	"net/http"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type Message struct {
	Text string `json:"text"`
}

type Service interface {
	Keygen(string, string) (string, error)
	TokenAuth(string) (bool, error)
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) UserLogin(ctx echo.Context) error {
	message := ctx.Request().Header
	username := message.Get("username")
	password := message.Get("password")

	userData, hashKey := cache.ConvertMap(message, "username", "password")

	cachePwd, err := cache.IsDataInCache("users", hashKey, "password")
	if err != nil {
		logger.Error("ошибка при поиске данных в кэше Redis: ", zap.Error(err))
		return err
	}

	// генерируем jwt токен и данных юзера для использования в дальнейшем
	key, err := e.s.Keygen(username, password)
	if err != nil {
		logger.Error("ошибка при генерации токена: %s", zap.Error(err))
		return err
	}

	returnMsg := Message{
		Text: key,
	}

	if cachePwd != nil && cachePwd == message.Get("password") {

		return ctx.JSON(http.StatusOK, returnMsg)
	} else if cachePwd == nil {
		dbMap, err := db.FetchUserData(username)
		if err != nil {
			return err
		}
		if dbMap != nil && dbMap["password"] == password {
			// сохраняем залогиненого юзера в кэш
			cache.SaveCache("users", userData)

			// авторизуем его
			return ctx.JSON(http.StatusOK, returnMsg)
		}
	}

	return errorz.ErrUserNotFound
}

// передаем в эту функцию username и password
func (e *Endpoint) NewUserRegistration(ctx echo.Context) error {
	message := ctx.Request().Header
	username := message.Get("username")
	password := message.Get("password")
	userData, hashKey := cache.ConvertMap(message, "username", "password")

	cachePwd, err := cache.IsDataInCache("users", hashKey, "password")
	if err != nil {
		return err
	}

	// если пароль у юзера есть, значит и юзер существует
	if cachePwd != "" {
		dbMap, err := db.FetchUserData(username)
		if err != sql.ErrNoRows && err != nil {
			return err
		}

		if len(dbMap) != 0 {
			logger.Info("такой пользователь уже существует")
			return errorz.ErrUserExists
		}
	}

	err = cache.SaveCache("users", userData)
	if err != nil {
		logger.Error("ошибка при регистрации пользователя: %s", zap.Error(err))
		return err
	}

	key, err := e.s.Keygen(username, password)
	if err != nil {
		logger.Error("ошибка при генерации токена: %s", zap.Error(err))
		return err
	}

	returnMsg := Message{
		Text: key,
	}
	return ctx.JSON(http.StatusOK, returnMsg)
}
