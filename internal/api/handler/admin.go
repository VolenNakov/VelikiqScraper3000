package handler

import (
	"OlxScraper/internal/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) VerifyUser(c echo.Context) error {
	var req model.VerifyRequest
	if err := h.ValidateRequest(c, &req); err != nil {
		fmt.Println(req)
		return c.JSON(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()

	resp, err := h.service.Admin.VerifyUser(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *Handler) GetUnverifiedUsers(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := h.service.Admin.GetUnverifiedUsers(ctx)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
