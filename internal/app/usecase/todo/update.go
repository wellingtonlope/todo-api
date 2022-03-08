package todo

import (
	"time"

	"github.com/wellingtonlope/todo-api/internal/app/repository"
)

type Update struct {
	todoRepository repository.TodoRepository
}

func NewUpdate(todoRepository repository.TodoRepository) *Update {
	return &Update{
		todoRepository: todoRepository,
	}
}

type UpdateInput struct {
	ID          string
	Title       string
	Description string
	UpdatedDate *time.Time
}

type UpdateOutput struct {
	ID          string
	Title       string
	Description string
	CreatedDate *time.Time
	UpdatedDate *time.Time
}

func (u *Update) Handle(input UpdateInput) (*UpdateOutput, error) {
	todo, err := u.todoRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	err = todo.Update(input.Title, input.Description, input.UpdatedDate)
	if err != nil {
		return nil, err
	}

	todo, err = u.todoRepository.Update(*todo)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedDate: todo.CreatedDate,
		UpdatedDate: todo.UpdatedDate,
	}, nil
}
