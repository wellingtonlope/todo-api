package repository

import (
	"errors"

	"github.com/wellingtonlope/todo-api/internal/domain"
)

var (
	ErrTodoNotFound      = errors.New("todo not found")
	ErrTodoAlreadyExists = errors.New("todo already exists")
)

type TodoRepository interface {
	GetAll() (*[]domain.Todo, error)
	GetByID(id string) (*domain.Todo, error)
	Insert(todo domain.Todo) (*domain.Todo, error)
	Update(todo domain.Todo) (*domain.Todo, error)
	DeleteByID(id string) error
}
