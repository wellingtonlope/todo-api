package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

type (
	todoGetAllOutput struct {
		ID          string    `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
	TodoGetAll struct {
		getAll todo.GetAll
	}
)

func NewTodoGetAll(getAll todo.GetAll) *TodoGetAll {
	return &TodoGetAll{getAll: getAll}
}

func (h *TodoGetAll) Handle(c echo.Context) error {
	outputs, err := h.getAll.Handle(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h.usecaseToHandlerOutput(outputs))
}

func (h *TodoGetAll) Path() string {
	return "/todos"
}

func (h *TodoGetAll) Method() string {
	return http.MethodGet
}

func (h *TodoGetAll) usecaseToHandlerOutput(todos []todo.GetAllOutput) []todoGetAllOutput {
	outputs := make([]todoGetAllOutput, 0, len(todos))
	for _, todo := range todos {
		outputs = append(outputs, todoGetAllOutput{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			CreatedAt:   todo.CreatedAt,
			UpdatedAt:   todo.UpdatedAt,
		})
	}
	return outputs
}
