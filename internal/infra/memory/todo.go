package memory

import (
	"context"

	"github.com/google/uuid"
	todoUC "github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

type todo struct {
	todos map[string]domain.Todo
}

func NewTodoRepository() *todo {
	return &todo{make(map[string]domain.Todo)}
}

func (r *todo) Create(_ context.Context, todo domain.Todo) (domain.Todo, error) {
	todo.ID = uuid.New().String()

	r.todos[todo.ID] = todo
	return todo, nil
}

func (r *todo) GetAll(_ context.Context) ([]domain.Todo, error) {
	todos := make([]domain.Todo, 0, len(r.todos))
	for _, item := range r.todos {
		todos = append(todos, item)
	}
	return todos, nil
}

func (r *todo) GetByID(_ context.Context, id string) (domain.Todo, error) {
	if item, ok := r.todos[id]; ok {
		return item, nil
	}
	return domain.Todo{}, todoUC.ErrGetByIDStoreNotFound
}

func (r *todo) DeleteByID(_ context.Context, id string) error {
	if _, ok := r.todos[id]; !ok {
		return todoUC.ErrDeleteByIDStoreNotFound
	}
	delete(r.todos, id)
	return nil
}

func (r *todo) Update(_ context.Context, todo domain.Todo) (domain.Todo, error) {
	if _, ok := r.todos[todo.ID]; ok {
		r.todos[todo.ID] = todo
		return todo, nil
	}
	return domain.Todo{}, todoUC.ErrUpdateStoreNotFound
}
