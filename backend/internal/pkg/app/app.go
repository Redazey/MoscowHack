package app

import (
	"github.com/gin-gonic/gin"
	"moscowhack/internal/app/endpoint"
	"moscowhack/internal/app/middleware"
	"moscowhack/internal/app/service"
)

type App struct {
	e *endpoint.Endpoint
	s *service.Service
}

func New() (*App, error) {
	a := &App{}

	router := gin.Default()
	v1 := router.Group("/api/v1")

	v1.Use(middleware.CORSPolicy())

	a.s = service.New()
	a.e = endpoint.New(a.s)

	//v1.GET("/contents", a.e.GetAllContents)

	err := router.Run(":4000")
	if err != nil {
		return a, err
	}

	return a, nil
}
