package todo

import (
	"context"
	"errors"
	"fmt"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
)

var ErrDeleteByIDStoreNotFound = errors.New("todo not found by ID")

type (
	DeleteByIDStore interface {
		DeleteByID(ctx context.Context, id string) error
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
		if errors.Is(err, ErrDeleteByIDStoreNotFound) {
			return usecase.NewError(fmt.Sprintf("todo not found with id %s", id),
				err, usecase.ErrorTypeNotFound)
		}
		return usecase.NewError("fail to delete a todo by id", err,
			usecase.ErrorTypeInternalError)
	}
	return nil
}
