package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/repository"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

type GetTodoByIdHandle struct {
	useCase todo.GetById
}

func (h *GetTodoByIdHandle) Handle(c echo.Context) error {
	id := c.Param("id")
	todoGet, err := h.useCase.Handle(todo.GetByIdInput{ID: id})

	if err != nil {
		if err == repository.ErrTodoNotFound {
			return c.JSON(http.StatusNotFound, wrapError(err))
		}
		return c.JSON(http.StatusInternalServerError, wrapError(err))
	}

	return c.JSON(http.StatusOK, Todo{
		Id:          todoGet.ID,
		Title:       todoGet.Title,
		Description: todoGet.Description,
		CreatedDate: todoGet.CreatedDate,
		UpdatedDate: todoGet.UpdatedDate,
	})
}
