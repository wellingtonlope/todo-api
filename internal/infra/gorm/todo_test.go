package gorm

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	todoUC "github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	"github.com/wellingtonlope/todo-api/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&TodoModel{})
	assert.NoError(t, err)
	return db
}

func TestNewTodoRepository(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTodoRepository(db)
	assert.NotNil(t, repo)
	assert.Equal(t, db, repo.db)
}

func TestCreate(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTodoRepository(db)
	date := time.Now().UTC()
	todo, _ := domain.NewTodo("Test Todo", "Test Description", date, nil)
	created, err := repo.Create(context.Background(), todo)
	assert.Nil(t, err)
	assert.NotEqual(t, "", created.ID)
	assert.Equal(t, todo.Title, created.Title)
	assert.Equal(t, todo.Description, created.Description)
	retrieved, _ := repo.GetByID(context.Background(), created.ID)
	assert.Equal(t, created, retrieved)
}

func TestGetAll(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTodoRepository(db)
	todos, err := repo.GetAll(context.Background())
	assert.Nil(t, err)
	assert.Len(t, todos, 0)

	date := time.Now().UTC()
	todo1, _ := domain.NewTodo("Todo 1", "", date, nil)
	todo2, _ := domain.NewTodo("Todo 2", "", date, nil)
	_, err = repo.Create(context.Background(), todo1)
	assert.Nil(t, err)
	_, err = repo.Create(context.Background(), todo2)
	assert.Nil(t, err)

	todos, err = repo.GetAll(context.Background())
	assert.Nil(t, err)
	assert.Len(t, todos, 2)
	// Check if both are present (order may vary)
	titles := make([]string, len(todos))
	for i, td := range todos {
		titles[i] = td.Title
	}
	assert.Contains(t, titles, "Todo 1")
	assert.Contains(t, titles, "Todo 2")
}

func TestGetByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTodoRepository(db)
	date := time.Now().UTC()
	todo, _ := domain.NewTodo("Test", "", date, nil)
	_, err := repo.Create(context.Background(), todo)
	assert.Nil(t, err)
	created, err := repo.Create(context.Background(), todo)
	assert.Nil(t, err)

	tests := []struct {
		name     string
		id       string
		expected domain.Todo
		err      error
	}{
		{"existing ID", created.ID, created, nil},
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
	db := setupTestDB(t)
	repo := NewTodoRepository(db)
	date := time.Now().UTC()
	todo, _ := domain.NewTodo("Test", "", date, nil)
	created, _ := repo.Create(context.Background(), todo)

	err := repo.DeleteByID(context.Background(), created.ID)
	assert.Nil(t, err)
	_, err = repo.GetByID(context.Background(), created.ID)
	assert.Equal(t, todoUC.ErrGetByIDStoreNotFound, err)

	err = repo.DeleteByID(context.Background(), "999") // non-existing
	assert.Nil(t, err)                                 // Delete is idempotent
}

func TestUpdate(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTodoRepository(db)
	date := time.Now().UTC()
	todo, _ := domain.NewTodo("Original", "", date, nil)
	created, _ := repo.Create(context.Background(), todo)

	updatedTodo := created
	updatedTodo.Title = "Updated"
	updatedTodo.Description = "New Desc"
	updatedTodo.UpdatedAt = time.Now().UTC()

	result, err := repo.Update(context.Background(), updatedTodo)
	assert.Nil(t, err)
	assert.Equal(t, updatedTodo.Title, result.Title)
	assert.Equal(t, updatedTodo.Description, result.Description)
	retrieved, _ := repo.GetByID(context.Background(), created.ID)
	assert.Equal(t, updatedTodo.Title, retrieved.Title)
	assert.Equal(t, updatedTodo.Description, retrieved.Description)

	_, err = repo.Update(context.Background(), domain.Todo{ID: "999", Title: "Non-existing"})
	assert.Equal(t, todoUC.ErrUpdateStoreNotFound, err)
}
