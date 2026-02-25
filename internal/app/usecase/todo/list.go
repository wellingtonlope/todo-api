package todo

import (
	"context"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

type (
	ListStore interface {
		List(context.Context, *domain.TodoStatus) ([]domain.Todo, error)
	}
	List interface {
		Handle(context.Context, ListInput) ([]TodoOutput, error)
	}
	ListInput struct {
		Status *domain.TodoStatus
	}
	list struct {
		store ListStore
	}
)

func NewList(store ListStore) *list {
	return &list{store}
}

func (uc *list) Handle(ctx context.Context, input ListInput) ([]TodoOutput, error) {
	todos, err := uc.store.List(ctx, input.Status)
	if err != nil {
		return []TodoOutput{}, usecase.NewError("fail to list todos",
			err, usecase.ErrorTypeInternalError)
	}
	return TodoOutputsFromDomain(todos), nil
}
