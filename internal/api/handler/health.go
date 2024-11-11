package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthStatus struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (h *Handler) HandleHealth(c echo.Context) error {
	health := &HealthStatus{
		ID:     "system",
		Status: "ok",
	}
	return c.JSON(http.StatusOK, health)
}
