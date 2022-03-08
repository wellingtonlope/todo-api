package todo

import (
	"time"

	"github.com/wellingtonlope/todo-api/internal/app/repository"
)

type GetAll struct {
	todoRepository repository.TodoRepository
}

func NewGetAll(todoRepository repository.TodoRepository) *GetAll {
	return &GetAll{
		todoRepository: todoRepository,
	}
}

type GetAllInput struct{}

type GetAllOutput struct {
	ID          string
	Title       string
	Description string
	CreatedDate *time.Time
	UpdatedDate *time.Time
}

func (u *GetAll) Handle(input GetAllInput) (*[]GetAllOutput, error) {
	todos, err := u.todoRepository.GetAll()
	if err != nil {
		return nil, err
	}

	output := make([]GetAllOutput, 0, len(*todos))

	for _, todo := range *todos {
		output = append(output, GetAllOutput{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			CreatedDate: todo.CreatedDate,
			UpdatedDate: todo.UpdatedDate,
		})
	}

	return &output, nil
}
