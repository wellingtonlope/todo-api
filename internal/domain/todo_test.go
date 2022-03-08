package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTodo(t *testing.T) {
	t.Run("should create a todo", func(t *testing.T) {
		expectedTitle := "title"
		expectedDescription := "description"
		expectedCreatedDate := time.Now()

		todo, err := NewTodo(expectedTitle, expectedDescription, &expectedCreatedDate)

		assert.Nil(t, err)
		assert.NotNil(t, todo)
		assert.Equal(t, expectedTitle, todo.Title)
		assert.Equal(t, expectedDescription, todo.Description)
		assert.Equal(t, expectedCreatedDate, *todo.CreatedDate)
	})

	t.Run("shouln't create a todo", func(t *testing.T) {
		expectedTitle := ""
		expectedDescription := ""
		expectedCreatedDate := time.Now()

		todo, err := NewTodo(expectedTitle, expectedDescription, &expectedCreatedDate)

		assert.Nil(t, todo)
		assert.NotNil(t, err)
		assert.Equal(t, ErrTitleRequired.Error(), err.Error())
	})
}

func TestUpdate(t *testing.T) {
	t.Run("should update a todo", func(t *testing.T) {
		expectedTitle := "title"
		expectedDescription := "description"
		expectedCreatedDate := time.Now()
		expectedUpdatedDate := time.Now()

		todo, _ := NewTodo(expectedTitle, expectedDescription, &expectedCreatedDate)
		err := todo.Update(expectedTitle, expectedDescription, &expectedUpdatedDate)

		assert.Nil(t, err)
		assert.Equal(t, expectedTitle, todo.Title)
		assert.Equal(t, expectedDescription, todo.Description)
		assert.Equal(t, expectedCreatedDate, *todo.CreatedDate)
		assert.Equal(t, expectedUpdatedDate, *todo.UpdatedDate)
	})

	t.Run("shouldn't update a todo", func(t *testing.T) {
		expectedTitle := ""
		expectedDescription := ""
		expectedCreatedDate := time.Now()
		expectedUpdatedDate := time.Now()

		todo, _ := NewTodo(expectedTitle, expectedDescription, &expectedCreatedDate)
		err := todo.Update(expectedTitle, expectedDescription, &expectedUpdatedDate)

		assert.Nil(t, todo)
		assert.NotNil(t, err)
		assert.Equal(t, ErrTitleRequired.Error(), err.Error())
	})
}
