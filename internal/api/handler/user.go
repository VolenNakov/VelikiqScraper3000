package handler

import (
	"OlxScraper/internal/model"
	"OlxScraper/internal/repository"
	"OlxScraper/internal/response"
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func (h *Handler) HandleRegister(c echo.Context) error {
	var req model.RegisterRequest
	if err := h.ValidateRequest(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()

	resp, err := h.service.User.Register(ctx, &req)
	if err != nil {
		log.Println(err)
		if errors.Is(err, repository.ErrDuplicateEmail) {
			return c.JSON(http.StatusConflict, response.Error("Email already exists", nil))
		}

		return c.JSON(http.StatusInternalServerError, response.Error("Internal server error", nil))
	}
	return c.JSON(http.StatusCreated, response.Success(resp))
}
func (h *Handler) HandleLogin(c echo.Context) error {
	var req model.LoginRequest
	if err := h.ValidateRequest(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()

	resp, err := h.service.Auth.Login(ctx, &req)
	if err != nil {
		log.Println(err)
		if errors.Is(err, repository.ErrUserNotFound) {
			return c.JSON(http.StatusUnauthorized, response.Error("Invalid credentials", nil))
		}
		if errors.Is(err, repository.ErrInvalidPassword) {
			return c.JSON(http.StatusUnauthorized, response.Error("Invalid credentials", nil))
		}
		if errors.Is(err, repository.ErrUnverifiedUser) {
			return c.JSON(http.StatusUnauthorized, response.Error("Unverified user", nil))
		}
		return c.JSON(http.StatusInternalServerError, response.Error("Internal server error", nil))
	}
	return c.JSON(http.StatusOK, response.Success(resp))
}
