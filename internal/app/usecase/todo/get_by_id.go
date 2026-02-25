package todo

import (
	"context"
	"errors"
	"fmt"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

var ErrGetByIDStoreNotFound = errors.New("todo not found by ID")

type (
	GetByIDStore interface {
		GetByID(context.Context, string) (domain.Todo, error)
	}
	GetByID interface {
		Handle(ctx context.Context, id string) (TodoOutput, error)
	}
	getByID struct {
		store GetByIDStore
	}
)

func NewGetByID(store GetByIDStore) *getByID {
	return &getByID{store}
}

func (uc *getByID) Handle(ctx context.Context, id string) (TodoOutput, error) {
	todo, err := uc.store.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrGetByIDStoreNotFound) {
			return TodoOutput{}, usecase.NewError(fmt.Sprintf("todo not found with id %s", id),
				err, usecase.ErrorTypeNotFound)
		}
		return TodoOutput{}, usecase.NewError("fail to get a todo by id",
			err, usecase.ErrorTypeInternalError)
	}
	return TodoOutputFromDomain(todo), nil
}
