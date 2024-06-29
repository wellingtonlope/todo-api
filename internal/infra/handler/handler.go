package handler

import "github.com/labstack/echo/v4"

type Handler interface {
	Handle(echo.Context) error
	Path() string
	Method() string
}
