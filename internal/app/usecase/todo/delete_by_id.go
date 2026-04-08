package todo

import (
	"context"
	"errors"
	"fmt"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

type (
	DeleteByIDStore interface {
		DeleteByID(context.Context, string) error
	}
	DeleteByID interface {
		Handle(ctx context.Context, id string) error
	}
	deleteByID struct {
		store DeleteByIDStore
	}
)

func NewDeleteByID(store DeleteByIDStore) *deleteByID {
	return &deleteByID{store}
}

func (uc *deleteByID) Handle(ctx context.Context, id string) error {
	err := uc.store.DeleteByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrTodoNotFound) {
			return usecase.NewError(fmt.Sprintf("todo not found with id %s", id),
				err, usecase.ErrorTypeNotFound)
		}
		return usecase.NewError("fail to delete a todo by id", err,
			usecase.ErrorTypeInternalError)
	}
	return nil
}
