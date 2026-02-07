package handler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

func TestTodoOutputFromUsecase(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	exampleDueDate, _ := time.Parse(time.DateOnly, "2024-12-31")

	testCases := []struct {
		name   string
		input  todo.TodoOutput
		output todoOutput
	}{
		{
			name: "should convert usecase.TodoOutput with all fields",
			input: todo.TodoOutput{
				ID:          "123",
				Title:       "Test Title",
				Description: "Test Description",
				Status:      "completed",
				DueDate:     &exampleDueDate,
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
			output: todoOutput{
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
			name: "should convert usecase.TodoOutput without due date",
			input: todo.TodoOutput{
				ID:          "456",
				Title:       "Another Title",
				Description: "Another Description",
				Status:      "pending",
				DueDate:     nil,
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
			output: todoOutput{
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
			name: "should convert empty usecase.TodoOutput",
			input: todo.TodoOutput{
				ID:          "",
				Title:       "",
				Description: "",
				Status:      "pending",
				DueDate:     nil,
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
			output: todoOutput{
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
			result := todoOutputFromUsecase(tc.input)
			assert.Equal(t, tc.output, result)
		})
	}
}

func TestTodoOutputsFromUsecase(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	exampleDueDate, _ := time.Parse(time.DateOnly, "2024-12-31")

	testCases := []struct {
		name   string
		input  []todo.TodoOutput
		output []todoOutput
	}{
		{
			name:   "should convert empty slice",
			input:  []todo.TodoOutput{},
			output: []todoOutput{},
		},
		{
			name: "should convert single item",
			input: []todo.TodoOutput{
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
			output: []todoOutput{
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
			input: []todo.TodoOutput{
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
			output: []todoOutput{
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
			result := todoOutputsFromUsecase(tc.input)
			assert.Equal(t, tc.output, result)
		})
	}
}
