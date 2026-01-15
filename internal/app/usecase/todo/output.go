package todo

import (
	"time"

	"github.com/wellingtonlope/todo-api/internal/domain"
)

// TodoOutput represents the output structure for todo operations
type TodoOutput struct {
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// TodoOutputFromDomain converts a domain.Todo to TodoOutput
func TodoOutputFromDomain(todo domain.Todo) TodoOutput {
	return TodoOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}

// TodoOutputsFromDomain converts a slice of domain.Todo to []TodoOutput
func TodoOutputsFromDomain(todos []domain.Todo) []TodoOutput {
	outputs := make([]TodoOutput, 0, len(todos))
	for _, todo := range todos {
		outputs = append(outputs, TodoOutputFromDomain(todo))
	}
	return outputs
}
