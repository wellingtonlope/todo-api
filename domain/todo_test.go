package domain

import (
	"testing"
	"time"
)

func TestNewTodo(t *testing.T) {
	messageTemplate := "Error: expected: %q, got: %q"
	expectedTitle, expectedDescription := "title 1", "description 1"

	t.Run("With title and description", func(t *testing.T) {
		todo, _ := NewTodo(expectedTitle, expectedDescription)

		if expectedTitle != todo.Title {
			t.Errorf(messageTemplate, expectedTitle, todo.Title)
		}

		if expectedDescription != todo.Description {
			t.Errorf(messageTemplate, expectedDescription, todo.Description)
		}

		if todo.ID == "" {
			t.Error("Error: expected a value, but got a empty string")
		}

		if todo.CreatedAt.IsZero() {
			t.Error("Error: expected the current date, but got a zero value")
		}
	})

	t.Run("With description empty", func(t *testing.T) {
		todo, _ := NewTodo(expectedTitle, "")

		if expectedTitle != todo.Title {
			t.Errorf(messageTemplate, expectedTitle, todo.Title)
		}

		if todo.Description != "" {
			t.Errorf("Error: expected a empty string, but got a %q", todo.Description)
		}

		if todo.ID == "" {
			t.Error("Error: expected a value, but got a empty string")
		}

		if todo.CreatedAt.IsZero() {
			t.Error("Error: expected the current date, but got a zero value")
		}
	})

	t.Run("With title empty", func(t *testing.T) {
		_, err := NewTodo("", "")
		if err == nil {
			t.Error("Error: expected a error, but got nothing")
		}
	})
}

func TestUpdateTodo(t *testing.T) {
	messageTemplate := "Error: expected: %q, got: %q"
	expectedTitle, expectedDescription := "title 1", "description 1"
	todoTemplate := Todo{
		Base: Base{
			ID:        "123",
			CreatedAt: time.Now(),
		},
		Title:       "title",
		Description: "description",
	}

	t.Run("Update title and description", func(t *testing.T) {
		got, _ := UpdateTodo(expectedTitle, expectedDescription, todoTemplate)

		if expectedTitle != got.Title {
			t.Errorf(messageTemplate, expectedTitle, got.Title)
		}

		if expectedDescription != got.Description {
			t.Errorf(messageTemplate, expectedDescription, got.Description)
		}

		if todoTemplate.ID != got.ID {
			t.Errorf(messageTemplate, todoTemplate.ID, got.ID)
		}

		if todoTemplate.CreatedAt != got.CreatedAt {
			t.Errorf(messageTemplate, todoTemplate.CreatedAt, got.CreatedAt)
		}

		if got.UpdatedAt.IsZero() {
			t.Error("Error: expected the current date, but got a zero value")
		}
	})

	t.Run("Update with description empty", func(t *testing.T) {
		got, _ := UpdateTodo(expectedTitle, "", todoTemplate)

		if expectedTitle != got.Title {
			t.Errorf(messageTemplate, expectedTitle, got.Title)
		}

		if got.Description != "" {
			t.Errorf("Error: expected a empty string, but got a %q", got.Description)
		}

		if todoTemplate.ID != got.ID {
			t.Errorf(messageTemplate, todoTemplate.ID, got.ID)
		}

		if todoTemplate.CreatedAt != got.CreatedAt {
			t.Errorf(messageTemplate, todoTemplate.CreatedAt, got.CreatedAt)
		}

		if got.UpdatedAt.IsZero() {
			t.Error("Error: expected the current date, but got a zero value")
		}
	})

	t.Run("Update with title empty", func(t *testing.T) {
		_, err := UpdateTodo("", "", todoTemplate)

		if err == nil {
			t.Error("Error: expected a error, but got nothing")
		}
	})

}
