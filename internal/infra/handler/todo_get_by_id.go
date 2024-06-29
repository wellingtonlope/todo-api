package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

type (
	todoGetByIDOutput struct {
		ID          string    `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
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

func (h *TodoGetByID) usecaseToHandlerOutput(todo todo.GetByIDOutput) todoGetByIDOutput {
	return todoGetByIDOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}
