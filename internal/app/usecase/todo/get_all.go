package todo

import (
	"context"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

type (
	GetAllStore interface {
		GetAll(ctx context.Context) ([]domain.Todo, error)
	}
	GetAll interface {
		Handle(ctx context.Context) ([]TodoOutput, error)
	}
	getAll struct {
		store GetAllStore
	}
)

func NewGetAll(store GetAllStore) *getAll {
	return &getAll{store}
}

func (uc *getAll) Handle(ctx context.Context) ([]TodoOutput, error) {
	todos, err := uc.store.GetAll(ctx)
	if err != nil {
		return []TodoOutput{}, usecase.NewError("fail to get all todos",
			err, usecase.ErrorTypeInternalError)
	}
	return TodoOutputsFromDomain(todos), nil
}
