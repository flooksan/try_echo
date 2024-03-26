package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler
type Handler struct {
}

func (h *Handler) Hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, Kloof!")
}

func (h *Handler) ThrowErr(ctx echo.Context) error {
	return ctx.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
}
