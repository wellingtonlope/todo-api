package todo

import (
	"time"

	"github.com/wellingtonlope/todo-api/internal/app/repository"
)

type GetById struct {
	todoRepository repository.TodoRepository
}

func NewGetById(todoRepository repository.TodoRepository) *GetById {
	return &GetById{
		todoRepository: todoRepository,
	}
}

type GetByIdInput struct {
	ID string
}

type GetByIdOutput struct {
	ID          string
	Title       string
	Description string
	CreatedDate *time.Time
	UpdatedDate *time.Time
}

func (u *GetById) Handle(input GetByIdInput) (*GetByIdOutput, error) {
	todo, err := u.todoRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetByIdOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedDate: todo.CreatedDate,
		UpdatedDate: todo.UpdatedDate,
	}, nil
}
