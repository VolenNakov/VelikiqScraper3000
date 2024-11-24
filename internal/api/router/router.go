package router

import (
	"OlxScraper/internal/api/handler"
	"OlxScraper/internal/middleware"
	"OlxScraper/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validate *validator.Validate
}

func (c CustomValidator) Validate(i interface{}) error {
	if err := c.validate.Struct(i); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	return nil
}

// NewValidator creates a new validator
func NewValidator() *CustomValidator {
	return &CustomValidator{
		validate: validator.New(),
	}
}

func New(service *service.Service) *echo.Echo {
	e := echo.New()

	e.Use(echomw.RemoveTrailingSlash())
	e.Use(echomw.Recover())
	e.Use(echomw.CORS())
	e.Use(echomw.Gzip())
	e.Use(middleware.EnhancedLogger())
	e.Validator = NewValidator()

	h := handler.New(service)

	e.GET("/health", h.HandleHealth)
	e.POST("/register", h.HandleRegister)
	e.POST("/login", h.HandleLogin)

	return e
}
