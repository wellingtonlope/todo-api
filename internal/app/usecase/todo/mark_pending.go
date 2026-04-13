package todo

import (
	"context"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
)

type (
	MarkAsPendingInput struct {
		ID string
	}
	MarkAsPendingStore = TodoUpdater
	MarkAsPending      interface {
		Handle(context.Context, MarkAsPendingInput) (TodoOutput, error)
	}
	markAsPending struct {
		store MarkAsPendingStore
		clock usecase.Clock
	}
)

func NewMarkAsPending(store MarkAsPendingStore, clock usecase.Clock) *markAsPending {
	return &markAsPending{
		store: store,
		clock: clock,
	}
}

func (uc *markAsPending) Handle(ctx context.Context, input MarkAsPendingInput) (TodoOutput, error) {
	todo, err := uc.store.GetByID(ctx, input.ID)
	if err != nil {
		if isNotFound(err) {
			return TodoOutput{}, notFoundError(input.ID, err)
		}
		return TodoOutput{}, internalError("fail to get a todo by id", err)
	}
	todo = todo.MarkAsPending(uc.clock.Now())
	todo, err = uc.store.Update(ctx, todo)
	if err != nil {
		if isNotFound(err) {
			return TodoOutput{}, notFoundError(input.ID, err)
		}
		return TodoOutput{}, internalError("fail to update a todo in the store", err)
	}
	return TodoOutputFromDomain(todo), nil
}
