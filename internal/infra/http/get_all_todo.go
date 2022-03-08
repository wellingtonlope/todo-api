package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

type GetAllTodoHandle struct {
	useCase todo.GetAll
}

func (h *GetAllTodoHandle) Handle(c echo.Context) error {
	todos, err := h.useCase.Handle(todo.GetAllInput{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, wrapError(err))
	}

	output := make([]Todo, 0, len(*todos))
	for _, item := range *todos {
		output = append(output, Todo{
			Id:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			CreatedDate: item.CreatedDate,
			UpdatedDate: item.UpdatedDate,
		})
	}

	return c.JSON(http.StatusOK, output)
}
