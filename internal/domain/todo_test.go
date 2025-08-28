package domain_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

func TestNewTodo(t *testing.T) {
	exampleTitle := "title example"
	exampleDescription := "description example"
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	exampleTodo := domain.Todo{
		Title:       exampleTitle,
		Description: exampleDescription,
		CreatedAt:   exampleDate,
		UpdatedAt:   exampleDate,
	}
	testCases := []struct {
		name        string
		title       string
		description string
		date        time.Time
		result      domain.Todo
		err         error
	}{
		{
			name:        "should fail when title is invalid",
			title:       "",
			description: exampleDescription,
			date:        exampleDate,
			result:      domain.Todo{},
			err:         fmt.Errorf("%w: title", domain.ErrTodoInvalidInput),
		},
		{
			name:        "should fail when title with space is invalid",
			title:       " ",
			description: exampleDescription,
			date:        exampleDate,
			result:      domain.Todo{},
			err:         fmt.Errorf("%w: title", domain.ErrTodoInvalidInput),
		},
		{
			name:        "should fail when date is invalid",
			title:       exampleTitle,
			description: exampleDescription,
			date:        time.Time{},
			result:      domain.Todo{},
			err:         fmt.Errorf("%w: date", domain.ErrTodoInvalidInput),
		},
		{
			name:        "should create todo",
			title:       exampleTitle,
			description: exampleDescription,
			date:        exampleDate,
			result:      exampleTodo,
			err:         nil,
		},
		{
			name:        "should create todo with title and description with spaces",
			title:       " title example ",
			description: " description example ",
			date:        exampleDate,
			result:      exampleTodo,
			err:         nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := domain.NewTodo(tc.title, tc.description, tc.date)
			assert.Equal(t, tc.result, result)
			assert.Equal(t, tc.err, err)
		})
	}
}

func TestTodo_Update(t *testing.T) {
	exampleTitle := "title example"
	exampleTitleUpdated := "title example updated"
	exampleDescription := "description example"
	exampleDescriptionUpdated := "description example updated"
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	exampleDateUpdated, _ := time.Parse(time.DateOnly, "2024-01-02")
	exampleTodo := domain.Todo{
		Title:       exampleTitle,
		Description: exampleDescription,
		CreatedAt:   exampleDate,
		UpdatedAt:   exampleDate,
	}
	exampleTodoUpdated := domain.Todo{
		Title:       exampleTitleUpdated,
		Description: exampleDescriptionUpdated,
		CreatedAt:   exampleDate,
		UpdatedAt:   exampleDateUpdated,
	}
	testCases := []struct {
		name        string
		title       string
		description string
		date        time.Time
		todo        domain.Todo
		result      domain.Todo
		err         error
	}{
		{
			name:        "should fail when title is invalid",
			title:       "",
			description: exampleDescriptionUpdated,
			date:        exampleDateUpdated,
			todo:        exampleTodo,
			result:      domain.Todo{},
			err:         fmt.Errorf("%w: title", domain.ErrTodoInvalidInput),
		},
		{
			name:        "should fail when title with space is invalid",
			title:       " ",
			description: exampleDescriptionUpdated,
			date:        exampleDateUpdated,
			todo:        exampleTodo,
			result:      domain.Todo{},
			err:         fmt.Errorf("%w: title", domain.ErrTodoInvalidInput),
		},
		{
			name:        "should fail when date is invalid",
			title:       exampleTitleUpdated,
			description: exampleDescriptionUpdated,
			date:        time.Time{},
			todo:        exampleTodo,
			result:      domain.Todo{},
			err:         fmt.Errorf("%w: date", domain.ErrTodoInvalidInput),
		},
		{
			name:        "should update todo",
			title:       exampleTitleUpdated,
			description: exampleDescriptionUpdated,
			date:        exampleDateUpdated,
			todo:        exampleTodo,
			result:      exampleTodoUpdated,
			err:         nil,
		},
		{
			name:        "should update todo with title and description with spaces",
			title:       " title example updated ",
			description: " description example updated ",
			date:        exampleDateUpdated,
			todo:        exampleTodo,
			result:      exampleTodoUpdated,
			err:         nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.todo.Update(tc.title, tc.description, tc.date)
			assert.Equal(t, tc.result, result)
			assert.Equal(t, tc.err, err)
		})
	}
}
