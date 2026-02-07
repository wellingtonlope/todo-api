package todo_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

func TestTodoOutputFromDomain(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	exampleDueDate, _ := time.Parse(time.DateOnly, "2024-12-31")

	testCases := []struct {
		name   string
		input  domain.Todo
		output todo.TodoOutput
	}{
		{
			name: "should convert domain.Todo with all fields",
			input: domain.Todo{
				ID:          "123",
				Title:       "Test Title",
				Description: "Test Description",
				Status:      domain.TodoStatusCompleted,
				DueDate:     &exampleDueDate,
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
			output: todo.TodoOutput{
				ID:          "123",
				Title:       "Test Title",
				Description: "Test Description",
				Status:      "completed",
				DueDate:     &exampleDueDate,
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
		},
		{
			name: "should convert domain.Todo without due date",
			input: domain.Todo{
				ID:          "456",
				Title:       "Another Title",
				Description: "Another Description",
				Status:      domain.TodoStatusPending,
				DueDate:     nil,
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
			output: todo.TodoOutput{
				ID:          "456",
				Title:       "Another Title",
				Description: "Another Description",
				Status:      "pending",
				DueDate:     nil,
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
		},
		{
			name: "should convert empty domain.Todo",
			input: domain.Todo{
				ID:          "",
				Title:       "",
				Description: "",
				Status:      domain.TodoStatusPending,
				DueDate:     nil,
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
			output: todo.TodoOutput{
				ID:          "",
				Title:       "",
				Description: "",
				Status:      "pending",
				DueDate:     nil,
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := todo.TodoOutputFromDomain(tc.input)
			assert.Equal(t, tc.output, result)
		})
	}
}

func TestTodoOutputsFromDomain(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	exampleDueDate, _ := time.Parse(time.DateOnly, "2024-12-31")

	testCases := []struct {
		name   string
		input  []domain.Todo
		output []todo.TodoOutput
	}{
		{
			name:   "should convert empty slice",
			input:  []domain.Todo{},
			output: []todo.TodoOutput{},
		},
		{
			name: "should convert single item",
			input: []domain.Todo{
				{
					ID:          "123",
					Title:       "Test Title",
					Description: "Test Description",
					Status:      domain.TodoStatusCompleted,
					DueDate:     &exampleDueDate,
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDate,
				},
			},
			output: []todo.TodoOutput{
				{
					ID:          "123",
					Title:       "Test Title",
					Description: "Test Description",
					Status:      "completed",
					DueDate:     &exampleDueDate,
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDate,
				},
			},
		},
		{
			name: "should convert multiple items",
			input: []domain.Todo{
				{
					ID:          "1",
					Title:       "First Todo",
					Description: "First Description",
					Status:      domain.TodoStatusPending,
					DueDate:     nil,
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDate,
				},
				{
					ID:          "2",
					Title:       "Second Todo",
					Description: "Second Description",
					Status:      domain.TodoStatusCompleted,
					DueDate:     &exampleDueDate,
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDate,
				},
			},
			output: []todo.TodoOutput{
				{
					ID:          "1",
					Title:       "First Todo",
					Description: "First Description",
					Status:      "pending",
					DueDate:     nil,
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDate,
				},
				{
					ID:          "2",
					Title:       "Second Todo",
					Description: "Second Description",
					Status:      "completed",
					DueDate:     &exampleDueDate,
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDate,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := todo.TodoOutputsFromDomain(tc.input)
			assert.Equal(t, tc.output, result)
		})
	}
}
