package handler

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	Handle(echo.Context) error
	Path() string
	Method() string
}

type todoOutput struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
