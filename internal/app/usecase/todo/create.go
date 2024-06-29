package todo

import (
	"context"
	"time"

	"github.com/wellingtonlope/todo-api/internal/app/repository"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/domain"
	"github.com/wellingtonlope/todo-api/pkg/clock"
)

type (
	CreateInput struct {
		Title       string
		Description string
	}
	CreateOutput struct {
		ID          string
		Title       string
		Description string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	Create interface {
		Handle(context.Context, CreateInput) (CreateOutput, error)
	}
	create struct {
		repository repository.Todo
		clock      clock.Client
	}
)

func NewCreate(repository repository.Todo, clock clock.Client) *create {
	return &create{
		repository: repository,
		clock:      clock,
	}
}

func (uc *create) Handle(ctx context.Context, input CreateInput) (CreateOutput, error) {
	todo, err := domain.NewTodo(input.Title, input.Description, uc.clock.Now())
	if err != nil {
		return CreateOutput{}, usecase.NewError(err.Error(), err, usecase.ErrorTypeBadRequest)
	}
	todo, err = uc.repository.Create(ctx, todo)
	if err != nil {
		return CreateOutput{}, usecase.NewError("fail to create a todo in the repository", err,
			usecase.ErrorTypeInternalError)
	}
	return uc.domainTodoToOutput(todo), nil
}

func (uc *create) domainTodoToOutput(todo domain.Todo) CreateOutput {
	return CreateOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}
