package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

type (
	TodoMarkPending struct {
		markAsPending todo.MarkAsPending
	}
)

func NewTodoMarkPending(markAsPending todo.MarkAsPending) *TodoMarkPending {
	return &TodoMarkPending{markAsPending: markAsPending}
}

// @Summary Mark a todo as pending
// @Description Mark an existing todo item as pending
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} todoOutput
// @Failure 404 {object} ErrorResponse
// @Router /todos/{id}/pending [post]
func (h *TodoMarkPending) Handle(c echo.Context) error {
	id := c.Param("id")
	output, err := h.markAsPending.Handle(c.Request().Context(), todo.MarkAsPendingInput{
		ID: id,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todoOutputFromUsecase(output))
}

func (h *TodoMarkPending) Path() string {
	return "/todos/:id/pending"
}

func (h *TodoMarkPending) Method() string {
	return http.MethodPost
}
