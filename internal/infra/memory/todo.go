package memory

import (
	"github.com/google/uuid"
	"github.com/wellingtonlope/todo-api/internal/app/repository"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

type TodoRepository struct {
	todos []domain.Todo
}

func (r *TodoRepository) GetAll() ([]domain.Todo, error) {
	return r.todos, nil
}

func (r *TodoRepository) Insert(todo domain.Todo) (domain.Todo, error) {
	todo.ID = uuid.New().String()

	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *TodoRepository) GetByID(id string) (domain.Todo, error) {
	for _, todo := range r.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return domain.Todo{}, repository.ErrTodoNotFound
}

func (r *TodoRepository) Update(todo domain.Todo) error {
	_, err := r.GetByID(todo.ID)
	if err != nil {
		return repository.ErrTodoNotFound
	}

	for index, item := range r.todos {
		if item.ID == todo.ID {
			r.todos[index] = todo
			break
		}
	}

	return nil
}

func (r *TodoRepository) DeleteByID(id string) error {
	_, err := r.GetByID(id)
	if err != nil {
		return repository.ErrTodoNotFound
	}

	for index, item := range r.todos {
		if item.ID == id {
			r.todos = removeIndex(r.todos, index)
		}
	}

	return nil
}

func removeIndex(s []domain.Todo, index int) []domain.Todo {
	return append(s[:index], s[index+1:]...)
}
