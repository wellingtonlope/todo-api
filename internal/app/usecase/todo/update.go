package todo

import (
	"context"
	"time"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
)

type (
	UpdateInput struct {
		ID          string
		Title       string
		Description string
		DueDate     *time.Time
	}
	UpdateStore = TodoUpdater
	Update      interface {
		Handle(context.Context, UpdateInput) (TodoOutput, error)
	}
	update struct {
		store UpdateStore
		clock usecase.Clock
	}
)

func NewUpdate(store UpdateStore, clock usecase.Clock) *update {
	return &update{
		store: store,
		clock: clock,
	}
}

func (uc *update) Handle(ctx context.Context, input UpdateInput) (TodoOutput, error) {
	todo, err := uc.store.GetByID(ctx, input.ID)
	if err != nil {
		if isNotFound(err) {
			return TodoOutput{}, notFoundError(input.ID, err)
		}
		return TodoOutput{}, internalError("fail to get a todo by id", err)
	}
	todo, err = todo.Update(input.Title, input.Description, uc.clock.Now(), input.DueDate)
	if err != nil {
		return TodoOutput{}, badRequestError(err.Error(), err)
	}
	todo, err = uc.store.Update(ctx, todo)
	if err != nil {
		if isNotFound(err) {
			return TodoOutput{}, notFoundError(input.ID, err)
		}
		return TodoOutput{}, internalError("fail to update a todo in the store", err)
	}
	return TodoOutputFromDomain(todo), nil
}
