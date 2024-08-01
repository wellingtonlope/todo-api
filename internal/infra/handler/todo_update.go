package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

type (
	todoUpdateInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	todoUpdateOutput struct {
		ID          string    `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
	TodoUpdate struct {
		update todo.Update
	}
)

func NewTodoUpdate(update todo.Update) *TodoUpdate {
	return &TodoUpdate{update: update}
}

func (h *TodoUpdate) Handle(c echo.Context) error {
	id := c.Param("id")
	var input todoUpdateInput
	if err := c.Bind(&input); err != nil {
		return usecase.NewError("invalid JSON input", err, usecase.ErrorTypeBadRequest)
	}
	output, err := h.update.Handle(c.Request().Context(), todo.UpdateInput{
		ID:          id,
		Title:       input.Title,
		Description: input.Description,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h.usecaseToHandlerOutput(output))
}

func (h *TodoUpdate) Path() string {
	return "/todos/:id"
}

func (h *TodoUpdate) Method() string {
	return http.MethodPut
}

func (h *TodoUpdate) usecaseToHandlerOutput(todo todo.UpdateOutput) todoUpdateOutput {
	return todoUpdateOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}
