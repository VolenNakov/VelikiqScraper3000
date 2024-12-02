package middleware

import (
	"OlxScraper/internal/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Middleware struct {
	jwtService auth.JWTService
}

func NewMiddleware(jwtService auth.JWTService) *Middleware {
	return &Middleware{jwtService: jwtService}
}

func (m *Middleware) AdminGuard(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, err := m.jwtService.ValidateToken(c.Request())
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
		}

		if claims.Role != "admin" {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": "Admin access required",
			})
		}

		return next(c)
	}
}
