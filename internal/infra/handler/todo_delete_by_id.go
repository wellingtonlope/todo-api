package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

type (
	TodoDeleteByID struct {
		deleteByID todo.DeleteByID
	}
)

func NewTodoDeleteByID(deleteByID todo.DeleteByID) *TodoDeleteByID {
	return &TodoDeleteByID{deleteByID: deleteByID}
}

// @Summary Delete a todo by ID
// @Description Delete a todo item by its ID
// @Tags todos
// @Param id path string true "Todo ID"
// @Success 204 "No Content"
// @Failure 404 {object} ErrorResponse
// @Router /todos/{id} [delete]
func (h *TodoDeleteByID) Handle(c echo.Context) error {
	id := c.Param("id")
	err := h.deleteByID.Handle(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *TodoDeleteByID) Path() string {
	return "/todos/:id"
}

func (h *TodoDeleteByID) Method() string {
	return http.MethodDelete
}
