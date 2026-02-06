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

func TestList(t *testing.T) {
	repo := NewTodoRepository()

	// Test empty list
	todos, err := repo.List(context.Background(), nil)
	assert.Nil(t, err)
	assert.Len(t, todos, 0)

	// Create test todos with different statuses
	todo1 := domain.Todo{ID: "1", Title: "Todo 1", Status: domain.TodoStatusPending}
	todo2 := domain.Todo{ID: "2", Title: "Todo 2", Status: domain.TodoStatusCompleted}
	todo3 := domain.Todo{ID: "3", Title: "Todo 3", Status: domain.TodoStatusPending}
	repo.todos["1"] = todo1
	repo.todos["2"] = todo2
	repo.todos["3"] = todo3

	// Test list all
	todos, err = repo.List(context.Background(), nil)
	assert.Nil(t, err)
	assert.Len(t, todos, 3)
	assert.Contains(t, todos, todo1)
	assert.Contains(t, todos, todo2)
	assert.Contains(t, todos, todo3)

	// Test filter by pending status
	pendingStatus := domain.TodoStatusPending
	pendingTodos, err := repo.List(context.Background(), &pendingStatus)
	assert.Nil(t, err)
	assert.Len(t, pendingTodos, 2)
	assert.Contains(t, pendingTodos, todo1)
	assert.Contains(t, pendingTodos, todo3)

	// Test filter by completed status
	completedStatus := domain.TodoStatusCompleted
	completedTodos, err := repo.List(context.Background(), &completedStatus)
	assert.Nil(t, err)
	assert.Len(t, completedTodos, 1)
	assert.Contains(t, completedTodos, todo2)
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
	assert.Equal(t, todoUC.ErrDeleteByIDStoreNotFound, err)
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
