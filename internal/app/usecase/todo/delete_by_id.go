package todo

import (
	"context"
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
		if isNotFound(err) {
			return notFoundError(id, err)
		}
		return internalError("fail to delete a todo by id", err)
	}
	return nil
}
