package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/repository"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

type DeleteTodoByIdHandle struct {
	useCase todo.DeleteById
}

func (h *DeleteTodoByIdHandle) Handle(c echo.Context) error {
	id := c.Param("id")
	_, err := h.useCase.Handle(todo.DeleteByIdInput{ID: id})

	if err != nil {
		if err == repository.ErrTodoNotFound {
			return c.JSON(http.StatusNotFound, wrapError(err))
		}
		return c.JSON(http.StatusInternalServerError, wrapError(err))
	}

	return c.NoContent(http.StatusNoContent)
}
