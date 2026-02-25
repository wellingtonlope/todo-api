package todo

import (
	"context"
	"errors"
	"fmt"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

var ErrMarkPendingStoreNotFound = errors.New("todo not found by ID")

type (
	MarkAsPendingInput struct {
		ID string
	}
	MarkAsPendingStore interface {
		GetByID(context.Context, string) (domain.Todo, error)
		Update(context.Context, domain.Todo) (domain.Todo, error)
	}
	MarkAsPending interface {
		Handle(context.Context, MarkAsPendingInput) (TodoOutput, error)
	}
	markAsPending struct {
		store MarkAsPendingStore
		clock usecase.Clock
	}
)

func NewMarkAsPending(store MarkAsPendingStore, clock usecase.Clock) *markAsPending {
	return &markAsPending{
		store: store,
		clock: clock,
	}
}

func (uc *markAsPending) Handle(ctx context.Context, input MarkAsPendingInput) (TodoOutput, error) {
	todo, err := uc.store.GetByID(ctx, input.ID)
	if err != nil {
		if errors.Is(err, ErrGetByIDStoreNotFound) {
			return TodoOutput{}, usecase.NewError(fmt.Sprintf("todo not found with id %s", input.ID),
				err, usecase.ErrorTypeNotFound)
		}
		return TodoOutput{}, usecase.NewError("fail to get a todo by id",
			err, usecase.ErrorTypeInternalError)
	}
	todo = todo.MarkAsPending(uc.clock.Now())
	todo, err = uc.store.Update(ctx, todo)
	if err != nil {
		if errors.Is(err, ErrMarkPendingStoreNotFound) {
			return TodoOutput{}, usecase.NewError(fmt.Sprintf("todo not found with id %s", input.ID),
				err, usecase.ErrorTypeNotFound)
		}
		return TodoOutput{}, usecase.NewError("fail to update a todo in the store", err,
			usecase.ErrorTypeInternalError)
	}
	return TodoOutputFromDomain(todo), nil
}
