package usecase

import (
	"github.com/wellingtonlope/todo-api/internal/app/repository"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

type UseCase interface {
	Handle(input interface{}) (output *interface{}, err error)
}

type AllUseCases struct {
	InsertTodo     *todo.Insert
	UpdateTodo     *todo.Update
	DeleteTodoById *todo.DeleteById
	GetAllTodos    *todo.GetAll
	GetTodoById    *todo.GetById
}

func GetUseCases(repositories repository.Repositories) (*AllUseCases, error) {
	repos, err := repositories.GetRepositories()
	if err != nil {
		return nil, err
	}

	return &AllUseCases{
		InsertTodo:     todo.NewInsert(repos.TodoRepository),
		UpdateTodo:     todo.NewUpdate(repos.TodoRepository),
		DeleteTodoById: todo.NewDeleteById(repos.TodoRepository),
		GetAllTodos:    todo.NewGetAll(repos.TodoRepository),
		GetTodoById:    todo.NewGetById(repos.TodoRepository),
	}, nil
}
