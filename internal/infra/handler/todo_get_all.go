package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

type (
	TodoGetAll struct {
		getAll todo.GetAll
	}
)

func NewTodoGetAll(getAll todo.GetAll) *TodoGetAll {
	return &TodoGetAll{getAll: getAll}
}

// @Summary Get all todos
// @Description Retrieve all todo items
// @Tags todos
// @Produce json
// @Success 200 {array} todoOutput
// @Router /todos [get]
func (h *TodoGetAll) Handle(c echo.Context) error {
	outputs, err := h.getAll.Handle(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todoOutputsFromUsecase(outputs))
}

func (h *TodoGetAll) Path() string {
	return "/todos"
}

func (h *TodoGetAll) Method() string {
	return http.MethodGet
}
