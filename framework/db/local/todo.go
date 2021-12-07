package local

import (
	"errors"

	"github.com/wellingtonlope/todo-api/application/myerrors"
	"github.com/wellingtonlope/todo-api/domain"
)

type TodoRepositoryLocal struct {
	Todos []domain.Todo
}

func (r *TodoRepositoryLocal) GetAll() (*[]domain.Todo, *myerrors.Error) {
	return &r.Todos, nil
}

func (r *TodoRepositoryLocal) GetById(id string) (*domain.Todo, *myerrors.Error) {
	for _, todo := range r.Todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, myerrors.NewError(errors.New("todo not found"), myerrors.REGISTER_NOT_FOUND)
}

func (r *TodoRepositoryLocal) Insert(todo *domain.Todo) (*domain.Todo, *myerrors.Error) {
	todoGet, _ := r.GetById(todo.ID)
	if todoGet != nil {
		return nil, myerrors.NewError(errors.New("todo is already exists"), myerrors.REGISTER_ALREADY_EXISTS)
	}
	r.Todos = append(r.Todos, *todo)
	return todo, nil
}

func (r *TodoRepositoryLocal) Update(todo *domain.Todo) (*domain.Todo, *myerrors.Error) {
	_, err := r.GetById(todo.ID)
	if err != nil {
		return nil, myerrors.NewError(errors.New("todo not found"), myerrors.REGISTER_NOT_FOUND)
	}

	for index, item := range r.Todos {
		if item.ID == todo.ID {
			r.Todos[index] = *todo
			break
		}
	}

	return todo, nil
}

func (r *TodoRepositoryLocal) Delete(id string) *myerrors.Error {
	_, err := r.GetById(id)
	if err != nil {
		return myerrors.NewError(errors.New("todo not found"), myerrors.REGISTER_NOT_FOUND)
	}

	for index, item := range r.Todos {
		if item.ID == id {
			r.Todos = removeIndex(r.Todos, index)
		}
	}

	return nil
}

func removeIndex(s []domain.Todo, index int) []domain.Todo {
	return append(s[:index], s[index+1:]...)
}
