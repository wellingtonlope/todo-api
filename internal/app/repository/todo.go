package repository

import (
	"context"
	"errors"

	"github.com/wellingtonlope/todo-api/internal/domain"
)

var (
	ErrTodoNotFound      = errors.New("todo not found")
	ErrTodoAlreadyExists = errors.New("todo already exists")
)

type Todo interface {
	GetAll(ctx context.Context) ([]domain.Todo, error)
	GetByID(ctx context.Context, id string) (domain.Todo, error)
	Create(ctx context.Context, todo domain.Todo) (domain.Todo, error)
	Update(ctx context.Context, todo domain.Todo) error
	DeleteByID(ctx context.Context, id string) error
}
