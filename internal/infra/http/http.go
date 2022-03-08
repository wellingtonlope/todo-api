package http

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
)

func InitHandlers(e *echo.Echo, useCases *usecase.AllUseCases) {
	insertTodoHandle := InsertTodoHandle{
		useCase: *useCases.InsertTodo,
	}
	getTodoByIdHandle := GetTodoByIdHandle{
		useCase: *useCases.GetTodoById,
	}
	getAllTodoHandle := GetAllTodoHandle{
		useCase: *useCases.GetAllTodos,
	}
	updateTodoHandle := UpdateTodoHandle{
		useCase: *useCases.UpdateTodo,
	}
	deleteTodoById := DeleteTodoByIdHandle{
		useCase: *useCases.DeleteTodoById,
	}

	e.POST("/todos", insertTodoHandle.Handle)
	e.GET("/todos/:id", getTodoByIdHandle.Handle)
	e.DELETE("/todos/:id", deleteTodoById.Handle)
	e.PUT("/todos/:id", updateTodoHandle.Handle)
	e.GET("/todos", getAllTodoHandle.Handle)
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func wrapError(err error) ErrorResponse {
	return ErrorResponse{Message: err.Error()}
}

type Todo struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedDate *time.Time `json:"created_date"`
	UpdatedDate *time.Time `json:"updated_date,omitempty"`
}
