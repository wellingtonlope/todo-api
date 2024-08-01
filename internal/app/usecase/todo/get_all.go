package todo

import (
	"context"
	"time"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

type (
	GetAllOutput struct {
		ID          string
		Title       string
		Description string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	GetAllStore interface {
		GetAll(ctx context.Context) ([]domain.Todo, error)
	}
	GetAll interface {
		Handle(ctx context.Context) ([]GetAllOutput, error)
	}
	getAll struct {
		store GetAllStore
	}
)

func NewGetAll(store GetAllStore) *getAll {
	return &getAll{store}
}

func (uc *getAll) Handle(ctx context.Context) ([]GetAllOutput, error) {
	todos, err := uc.store.GetAll(ctx)
	if err != nil {
		return []GetAllOutput{}, usecase.NewError("fail to get all todos",
			err, usecase.ErrorTypeInternalError)
	}
	return uc.domainTodoToOutput(todos), nil
}

func (uc *getAll) domainTodoToOutput(todos []domain.Todo) []GetAllOutput {
	outputs := make([]GetAllOutput, 0, len(todos))
	for _, todo := range todos {
		outputs = append(outputs, GetAllOutput{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			CreatedAt:   todo.CreatedAt,
			UpdatedAt:   todo.UpdatedAt,
		})
	}
	return outputs
}
