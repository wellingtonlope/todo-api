package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

type (
	TodoGetByID struct {
		getByID todo.GetByID
	}
)

func NewTodoGetByID(getByID todo.GetByID) *TodoGetByID {
	return &TodoGetByID{getByID: getByID}
}

func (h *TodoGetByID) Handle(c echo.Context) error {
	id := c.Param("id")
	output, err := h.getByID.Handle(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h.usecaseToHandlerOutput(output))
}

func (h *TodoGetByID) Path() string {
	return "/todos/:id"
}

func (h *TodoGetByID) Method() string {
	return http.MethodGet
}

func (h *TodoGetByID) usecaseToHandlerOutput(todo todo.GetByIDOutput) todoOutput {
	return todoOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}
