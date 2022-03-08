package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

type InsertTodoHandle struct {
	useCase todo.Insert
}

func (h *InsertTodoHandle) Handle(c echo.Context) error {
	var dto Todo
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, wrapError(err))
	}
	today := time.Now()

	todoGet, err := h.useCase.Handle(todo.InsertInput{
		Title:       dto.Title,
		Description: dto.Description,
		CreatedDate: &today,
	})

	if err != nil {
		if err == domain.ErrTitleRequired {
			return c.JSON(http.StatusBadRequest, wrapError(err))
		}
		return c.JSON(http.StatusInternalServerError, wrapError(err))
	}

	return c.JSON(http.StatusCreated, Todo{
		Id:          todoGet.ID,
		Title:       todoGet.Title,
		Description: todoGet.Description,
		CreatedDate: todoGet.CreatedDate,
	})
}
