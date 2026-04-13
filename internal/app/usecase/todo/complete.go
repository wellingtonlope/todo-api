package todo

import (
	"context"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
)

type (
	CompleteInput struct {
		ID string
	}
	CompleteStore = TodoUpdater
	Complete      interface {
		Handle(context.Context, CompleteInput) (TodoOutput, error)
	}
	complete struct {
		store CompleteStore
		clock usecase.Clock
	}
)

func NewComplete(store CompleteStore, clock usecase.Clock) *complete {
	return &complete{
		store: store,
		clock: clock,
	}
}

func (uc *complete) Handle(ctx context.Context, input CompleteInput) (TodoOutput, error) {
	todo, err := uc.store.GetByID(ctx, input.ID)
	if err != nil {
		if isNotFound(err) {
			return TodoOutput{}, notFoundError(input.ID, err)
		}
		return TodoOutput{}, internalError("fail to get a todo by id", err)
	}
	todo = todo.MarkAsCompleted(uc.clock.Now())
	todo, err = uc.store.Update(ctx, todo)
	if err != nil {
		if isNotFound(err) {
			return TodoOutput{}, notFoundError(input.ID, err)
		}
		return TodoOutput{}, internalError("fail to update a todo in the store", err)
	}
	return TodoOutputFromDomain(todo), nil
}
