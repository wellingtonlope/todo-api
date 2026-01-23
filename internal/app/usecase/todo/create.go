package todo

import (
	"context"
	"time"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

type (
	CreateInput struct {
		Title       string
		Description string
		DueDate     *time.Time
	}
	CreateStore interface {
		Create(ctx context.Context, todo domain.Todo) (domain.Todo, error)
	}
	Create interface {
		Handle(context.Context, CreateInput) (TodoOutput, error)
	}
	create struct {
		store CreateStore
		clock usecase.Clock
	}
)

func NewCreate(store CreateStore, clock usecase.Clock) *create {
	return &create{
		store: store,
		clock: clock,
	}
}

func (uc *create) Handle(ctx context.Context, input CreateInput) (TodoOutput, error) {
	todo, err := domain.NewTodo(input.Title, input.Description, uc.clock.Now(), input.DueDate)
	if err != nil {
		return TodoOutput{}, usecase.NewError(err.Error(), err, usecase.ErrorTypeBadRequest)
	}
	todo, err = uc.store.Create(ctx, todo)
	if err != nil {
		return TodoOutput{}, usecase.NewError("fail to create a todo in the repository", err,
			usecase.ErrorTypeInternalError)
	}
	return TodoOutputFromDomain(todo), nil
}
