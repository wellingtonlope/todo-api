package todo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

var ErrGetByIDStoreNotFound = errors.New("todo not found by ID")

type (
	GetByIDOutput struct {
		ID          string
		Title       string
		Description string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	GetByIDStore interface {
		GetByID(ctx context.Context, id string) (domain.Todo, error)
	}
	GetByID interface {
		Handle(ctx context.Context, id string) (GetByIDOutput, error)
	}
	getByID struct {
		store GetByIDStore
	}
)

func NewGetByID(store GetByIDStore) *getByID {
	return &getByID{store}
}

func (uc *getByID) Handle(ctx context.Context, id string) (GetByIDOutput, error) {
	todo, err := uc.store.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrGetByIDStoreNotFound) {
			return GetByIDOutput{}, usecase.NewError(fmt.Sprintf("todo not found with id %s", id),
				err, usecase.ErrorTypeNotFound)
		}
		return GetByIDOutput{}, usecase.NewError("fail to get a todo by id",
			err, usecase.ErrorTypeInternalError)
	}
	return uc.domainTodoToOutput(todo), nil
}

func (uc *getByID) domainTodoToOutput(todo domain.Todo) GetByIDOutput {
	return GetByIDOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}
