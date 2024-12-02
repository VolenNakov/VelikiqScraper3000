package router

import (
	"OlxScraper/internal/api/handler"
	"OlxScraper/internal/auth"
	"OlxScraper/internal/middleware"
	"OlxScraper/internal/service"
	"OlxScraper/internal/validation"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

func New(service *service.Service, jwtService auth.JWTService) *echo.Echo {
	e := echo.New()

	e.Use(echomw.RemoveTrailingSlash())
	e.Use(echomw.Recover())
	e.Use(echomw.CORS())
	e.Use(echomw.Gzip())
	e.Use(middleware.EnhancedLogger())
	e.Validator = validation.NewValidator()

	adminMiddleware := middleware.NewMiddleware(jwtService)

	h := handler.New(service)

	e.GET("/health", h.HandleHealth)
	e.POST("/register", h.HandleRegister)
	e.POST("/login", h.HandleLogin)

	adminGroup := e.Group("/admin")
	adminGroup.Use(adminMiddleware.AdminGuard)
	e.POST("/verify", h.VerifyUser)
	e.GET("/getUnverified", h.GetUnverifiedUsers)

	return e
}
