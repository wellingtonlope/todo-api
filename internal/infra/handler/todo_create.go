package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

type (
	todoCreateInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
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
	return c.JSON(http.StatusCreated, todoOutputFromUsecase(output))
}

func (h *TodoCreate) Path() string {
	return "/todos"
}

func (h *TodoCreate) Method() string {
	return http.MethodPost
}
