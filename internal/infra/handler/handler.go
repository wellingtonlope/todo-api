package handler

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

type Handler interface {
	Handle(echo.Context) error
	Path() string
	Method() string
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type todoOutput struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// todoOutputFromUsecase converts a usecase TodoOutput to handler todoOutput
func todoOutputFromUsecase(usecaseOutput todo.TodoOutput) todoOutput {
	return todoOutput{
		ID:          usecaseOutput.ID,
		Title:       usecaseOutput.Title,
		Description: usecaseOutput.Description,
		DueDate:     usecaseOutput.DueDate,
		CreatedAt:   usecaseOutput.CreatedAt,
		UpdatedAt:   usecaseOutput.UpdatedAt,
	}
}

// todoOutputsFromUsecase converts a slice of usecase TodoOutput to []todoOutput
func todoOutputsFromUsecase(usecaseOutputs []todo.TodoOutput) []todoOutput {
	outputs := make([]todoOutput, 0, len(usecaseOutputs))
	for _, usecaseOutput := range usecaseOutputs {
		outputs = append(outputs, todoOutputFromUsecase(usecaseOutput))
	}
	return outputs
}
