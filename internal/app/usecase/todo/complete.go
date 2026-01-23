package todo

import (
	"context"
	"errors"
	"fmt"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

var ErrCompleteStoreNotFound = errors.New("todo not found by ID")

type (
	CompleteInput struct {
		ID string
	}
	CompleteStore interface {
		GetByID(ctx context.Context, id string) (domain.Todo, error)
		Update(ctx context.Context, todo domain.Todo) (domain.Todo, error)
	}
	Complete interface {
		Handle(context.Context, CompleteInput) (TodoOutput, error)
	}
	complete struct {
		store CompleteStore
		clock usecase.Clock
	}
)

func NewComplete(store CompleteStore, clock usecase.Clock) *complete {
	return &complete{
		store: store,
		clock: clock,
	}
}

func (uc *complete) Handle(ctx context.Context, input CompleteInput) (TodoOutput, error) {
	todo, err := uc.store.GetByID(ctx, input.ID)
	if err != nil {
		if errors.Is(err, ErrGetByIDStoreNotFound) {
			return TodoOutput{}, usecase.NewError(fmt.Sprintf("todo not found with id %s", input.ID),
				err, usecase.ErrorTypeNotFound)
		}
		return TodoOutput{}, usecase.NewError("fail to get a todo by id",
			err, usecase.ErrorTypeInternalError)
	}
	todo = todo.MarkAsCompleted(uc.clock.Now())
	todo, err = uc.store.Update(ctx, todo)
	if err != nil {
		if errors.Is(err, ErrCompleteStoreNotFound) {
			return TodoOutput{}, usecase.NewError(fmt.Sprintf("todo not found with id %s", input.ID),
				err, usecase.ErrorTypeNotFound)
		}
		return TodoOutput{}, usecase.NewError("fail to update a todo in the store", err,
			usecase.ErrorTypeInternalError)
	}
	return TodoOutputFromDomain(todo), nil
}
