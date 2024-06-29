package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

type (
	todoCreateInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	todoCreateOutput struct {
		ID          string    `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
	TodoCreate struct {
		create todo.Create
	}
)

func NewTodoCreate(create todo.Create) *TodoCreate {
	return &TodoCreate{create: create}
}

func (h *TodoCreate) Handle(c echo.Context) error {
	var input todoCreateInput
	if err := c.Bind(&input); err != nil {
		return usecase.NewError("invalid JSON input", err, usecase.ErrorTypeBadRequest)
	}
	output, err := h.create.Handle(c.Request().Context(), todo.CreateInput{
		Title:       input.Title,
		Description: input.Description,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, h.usecaseToHandlerOutput(output))
}

func (h *TodoCreate) Path() string {
	return "/todos"
}

func (h *TodoCreate) Method() string {
	return http.MethodPost
}

func (h *TodoCreate) usecaseToHandlerOutput(todo todo.CreateOutput) todoCreateOutput {
	return todoCreateOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}
