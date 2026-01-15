package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

type (
	todoUpdateInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
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
	return c.JSON(http.StatusOK, todoOutputFromUsecase(output))
}

func (h *TodoUpdate) Path() string {
	return "/todos/:id"
}

func (h *TodoUpdate) Method() string {
	return http.MethodPut
}
