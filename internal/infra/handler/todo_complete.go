package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

type (
	TodoComplete struct {
		complete todo.Complete
	}
)

func NewTodoComplete(complete todo.Complete) *TodoComplete {
	return &TodoComplete{complete: complete}
}

// @Summary Mark a todo as completed
// @Description Mark an existing todo item as completed
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} todoOutput
// @Failure 404 {object} ErrorResponse
// @Router /todos/{id}/complete [post]
func (h *TodoComplete) Handle(c echo.Context) error {
	id := c.Param("id")
	output, err := h.complete.Handle(c.Request().Context(), todo.CompleteInput{
		ID: id,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todoOutputFromUsecase(output))
}

func (h *TodoComplete) Path() string {
	return "/todos/:id/complete"
}

func (h *TodoComplete) Method() string {
	return http.MethodPost
}
