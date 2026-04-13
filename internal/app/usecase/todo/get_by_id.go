package todo

import (
	"context"

	"github.com/wellingtonlope/todo-api/internal/domain"
)

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
		if isNotFound(err) {
			return TodoOutput{}, notFoundError(id, err)
		}
		return TodoOutput{}, internalError("fail to get a todo by id", err)
	}
	return TodoOutputFromDomain(todo), nil
}
