package todo

import (
	"time"

	"github.com/wellingtonlope/todo-api/internal/app/repository"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

type Insert struct {
	todoRepository repository.TodoRepository
}

func NewInsert(todoRepository repository.TodoRepository) *Insert {
	return &Insert{
		todoRepository: todoRepository,
	}
}

type InsertInput struct {
	Title       string
	Description string
	CreatedDate *time.Time
}

type InsertOutput struct {
	ID          string
	Title       string
	Description string
	CreatedDate *time.Time
}

func (u *Insert) Handle(input InsertInput) (*InsertOutput, error) {
	todo, err := domain.NewTodo(input.Title, input.Description, input.CreatedDate)
	if err != nil {
		return nil, err
	}

	todo, err = u.todoRepository.Insert(*todo)
	if err != nil {
		return nil, err
	}

	return &InsertOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedDate: todo.CreatedDate,
	}, nil
}
