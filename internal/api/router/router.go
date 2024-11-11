package router

import (
	"OlxScraper/internal/api/handler"
	"OlxScraper/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(service *service.Service) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	h := handler.New(service)

	e.GET("/health", h.HandleHealth)
	e.POST("/register", h.HandleRegister)

	return e
}
