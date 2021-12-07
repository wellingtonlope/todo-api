package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/application/usecase"
)

func InitHandlers(e *echo.Echo, useCases *usecase.AllUseCases) {
	initTodoHandler(e, useCases.TodoUseCase)
}
