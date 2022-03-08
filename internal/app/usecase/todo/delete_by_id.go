package todo

import "github.com/wellingtonlope/todo-api/internal/app/repository"

type DeleteById struct {
	todoRepository repository.TodoRepository
}

func NewDeleteById(todoRepository repository.TodoRepository) *DeleteById {
	return &DeleteById{
		todoRepository: todoRepository,
	}
}

type DeleteByIdInput struct {
	ID string
}

type DeleteByIdOutput struct{}

func (u *DeleteById) Handle(input DeleteByIdInput) (*DeleteByIdOutput, error) {
	err := u.todoRepository.DeleteByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteByIdOutput{}, nil
}
