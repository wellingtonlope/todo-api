package repositories

import (
	"github.com/wellingtonlope/todo-api/application/myerrors"
	"github.com/wellingtonlope/todo-api/domain"
)

type TodoRepository interface {
	GetAll() (*[]domain.Todo, *myerrors.Error)
	GetById(id string) (*domain.Todo, *myerrors.Error)
	Insert(todo *domain.Todo) (*domain.Todo, *myerrors.Error)
	Update(todo *domain.Todo) (*domain.Todo, *myerrors.Error)
	Delete(id string) *myerrors.Error
}
