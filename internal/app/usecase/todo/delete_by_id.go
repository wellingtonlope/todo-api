package todo

import (
	"context"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
)

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
		return usecase.NewError("fail to detele a todo by id", err,
			usecase.ErrorTypeInternalError)
	}
	return nil
}
