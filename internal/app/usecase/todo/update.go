package todo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/domain"
	"github.com/wellingtonlope/todo-api/pkg/clock"
)

var ErrUpdateStoreNotFound = errors.New("todo not found by ID")

type (
	UpdateInput struct {
		ID          string
		Title       string
		Description string
	}
	UpdateOutput struct {
		ID          string
		Title       string
		Description string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	UpdateStore interface {
		GetByID(ctx context.Context, id string) (domain.Todo, error)
		Update(ctx context.Context, todo domain.Todo) error
	}
	Update interface {
		Handle(context.Context, UpdateInput) (UpdateOutput, error)
	}
	update struct {
		store UpdateStore
		clock clock.Client
	}
)

func NewUpdate(store UpdateStore, clock clock.Client) *update {
	return &update{
		store: store,
		clock: clock,
	}
}

func (uc *update) Handle(ctx context.Context, input UpdateInput) (UpdateOutput, error) {
	todo, err := uc.store.GetByID(ctx, input.ID)
	if err != nil {
		if errors.Is(err, ErrGetByIDStoreNotFound) {
			return UpdateOutput{}, usecase.NewError(fmt.Sprintf("todo not found with id %s", input.ID),
				err, usecase.ErrorTypeNotFound)
		}
		return UpdateOutput{}, usecase.NewError("fail to get a todo by id",
			err, usecase.ErrorTypeInternalError)

	}
	todo, err = todo.Update(input.Title, input.Description, uc.clock.Now())
	if err != nil {
		return UpdateOutput{}, usecase.NewError(err.Error(), err, usecase.ErrorTypeBadRequest)
	}
	err = uc.store.Update(ctx, todo)
	if err != nil {
		if errors.Is(err, ErrUpdateStoreNotFound) {
			return UpdateOutput{}, usecase.NewError(fmt.Sprintf("todo not found with id %s", input.ID),
				err, usecase.ErrorTypeNotFound)
		}
		return UpdateOutput{}, usecase.NewError("fail to update a todo in the store", err,
			usecase.ErrorTypeInternalError)
	}
	return uc.domainTodoToOutput(todo), nil
}

func (uc *update) domainTodoToOutput(todo domain.Todo) UpdateOutput {
	return UpdateOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}
