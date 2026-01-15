package memory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	todoUC "github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

func TestNewTodoRepository(t *testing.T) {
	repo := NewTodoRepository()
	assert.NotNil(t, repo)
	assert.NotNil(t, repo.todos)
	assert.Len(t, repo.todos, 0)
}

func TestCreate(t *testing.T) {
	repo := NewTodoRepository()
	todo := domain.Todo{Title: "Test Todo", Description: "Test Description"}
	created, err := repo.Create(context.Background(), todo)
	assert.Nil(t, err)
	assert.NotEqual(t, "", created.ID)
	assert.Equal(t, todo.Title, created.Title)
	assert.Equal(t, todo.Description, created.Description)
	retrieved, _ := repo.GetByID(context.Background(), created.ID)
	assert.Equal(t, created, retrieved)
}

func TestGetAll(t *testing.T) {
	repo := NewTodoRepository()
	todos, err := repo.GetAll(context.Background())
	assert.Nil(t, err)
	assert.Len(t, todos, 0)

	todo1 := domain.Todo{ID: "1", Title: "Todo 1"}
	todo2 := domain.Todo{ID: "2", Title: "Todo 2"}
	repo.todos["1"] = todo1
	repo.todos["2"] = todo2

	todos, err = repo.GetAll(context.Background())
	assert.Nil(t, err)
	assert.Len(t, todos, 2)
	// Note: Order may vary, so check if both are present
	assert.Contains(t, todos, todo1)
	assert.Contains(t, todos, todo2)
}

func TestGetByID(t *testing.T) {
	repo := NewTodoRepository()
	todo := domain.Todo{ID: "123", Title: "Test"}
	repo.todos["123"] = todo

	tests := []struct {
		name     string
		id       string
		expected domain.Todo
		err      error
	}{
		{"existing ID", "123", todo, nil},
		{"non-existing ID", "999", domain.Todo{}, todoUC.ErrGetByIDStoreNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := repo.GetByID(context.Background(), tt.id)
			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestDeleteByID(t *testing.T) {
	repo := NewTodoRepository()
	todo := domain.Todo{ID: "123", Title: "Test"}
	repo.todos["123"] = todo

	err := repo.DeleteByID(context.Background(), "123")
	assert.Nil(t, err)
	assert.Len(t, repo.todos, 0)

	err = repo.DeleteByID(context.Background(), "999") // non-existing
	assert.Nil(t, err)                                 // Delete is idempotent, no error
}

func TestUpdate(t *testing.T) {
	repo := NewTodoRepository()
	todo := domain.Todo{ID: "123", Title: "Original"}
	repo.todos["123"] = todo

	updatedTodo := domain.Todo{ID: "123", Title: "Updated", Description: "New Desc"}
	result, err := repo.Update(context.Background(), updatedTodo)
	assert.Nil(t, err)
	assert.Equal(t, updatedTodo, result)
	retrieved, _ := repo.GetByID(context.Background(), "123")
	assert.Equal(t, updatedTodo, retrieved)

	_, err = repo.Update(context.Background(), domain.Todo{ID: "999", Title: "Non-existing"})
	assert.Equal(t, todoUC.ErrUpdateStoreNotFound, err)
}
