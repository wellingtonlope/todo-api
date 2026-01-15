package todo

import (
	"context"
	"errors"
	"fmt"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

var ErrUpdateStoreNotFound = errors.New("todo not found by ID")

type (
	UpdateInput struct {
		ID          string
		Title       string
		Description string
	}
	UpdateStore interface {
		GetByID(ctx context.Context, id string) (domain.Todo, error)
		Update(ctx context.Context, todo domain.Todo) (domain.Todo, error)
	}
	Update interface {
		Handle(context.Context, UpdateInput) (TodoOutput, error)
	}
	update struct {
		store UpdateStore
		clock usecase.Clock
	}
)

func NewUpdate(store UpdateStore, clock usecase.Clock) *update {
	return &update{
		store: store,
		clock: clock,
	}
}

func (uc *update) Handle(ctx context.Context, input UpdateInput) (TodoOutput, error) {
	todo, err := uc.store.GetByID(ctx, input.ID)
	if err != nil {
		if errors.Is(err, ErrGetByIDStoreNotFound) {
			return TodoOutput{}, usecase.NewError(fmt.Sprintf("todo not found with id %s", input.ID),
				err, usecase.ErrorTypeNotFound)
		}
		return TodoOutput{}, usecase.NewError("fail to get a todo by id",
			err, usecase.ErrorTypeInternalError)

	}
	todo, err = todo.Update(input.Title, input.Description, uc.clock.Now())
	if err != nil {
		return TodoOutput{}, usecase.NewError(err.Error(), err, usecase.ErrorTypeBadRequest)
	}
	todo, err = uc.store.Update(ctx, todo)
	if err != nil {
		if errors.Is(err, ErrUpdateStoreNotFound) {
			return TodoOutput{}, usecase.NewError(fmt.Sprintf("todo not found with id %s", input.ID),
				err, usecase.ErrorTypeNotFound)
		}
		return TodoOutput{}, usecase.NewError("fail to update a todo in the store", err,
			usecase.ErrorTypeInternalError)
	}
	return TodoOutputFromDomain(todo), nil
}
