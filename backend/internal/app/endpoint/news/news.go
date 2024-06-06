package news

import (
	"github.com/labstack/echo"
)

type Message struct {
	news map[string]map[string]string
}

type Service interface {
	GetNews() (map[string]map[string]string, error)
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) News(ctx echo.Context) error {
	return nil
}
