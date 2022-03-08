package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/repository"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

type UpdateTodoHandle struct {
	useCase todo.Update
}

func (h *UpdateTodoHandle) Handle(c echo.Context) error {
	id := c.Param("id")
	var dto Todo
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, wrapError(err))
	}
	today := time.Now()

	todoGet, err := h.useCase.Handle(todo.UpdateInput{
		ID:          id,
		Title:       dto.Title,
		Description: dto.Description,
		UpdatedDate: &today,
	})

	if err != nil {
		if err == domain.ErrTitleRequired {
			return c.JSON(http.StatusBadRequest, wrapError(err))
		}
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
