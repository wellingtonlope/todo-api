package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

type (
	TodoList struct {
		list todo.List
	}
)

func NewTodoList(list todo.List) *TodoList {
	return &TodoList{list: list}
}

// @Summary List todos
// @Description Retrieve todo items with optional status filter
// @Tags todos
// @Produce json
// @Param status query string false "Filter by status (pending or completed)"
// @Success 200 {array} todoOutput
// @Router /todos [get]
func (h *TodoList) Handle(c echo.Context) error {
	statusParam := c.QueryParam("status")

	var status *domain.TodoStatus
	if statusParam != "" {
		s := domain.TodoStatus(statusParam)
		status = &s
	}

	input := todo.ListInput{Status: status}
	outputs, err := h.list.Handle(c.Request().Context(), input)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todoOutputsFromUsecase(outputs))
}

func (h *TodoList) Path() string {
	return "/todos"
}

func (h *TodoList) Method() string {
	return http.MethodGet
}
