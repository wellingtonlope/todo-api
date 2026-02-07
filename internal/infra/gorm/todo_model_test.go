package gorm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

func TestTodoModel_TableName(t *testing.T) {
	model := TodoModel{}
	assert.Equal(t, "todos", model.TableName())
}

func TestToDomain(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	exampleDueDate, _ := time.Parse(time.DateOnly, "2024-12-31")

	testCases := []struct {
		name   string
		input  TodoModel
		output domain.Todo
	}{
		{
			name: "should convert TodoModel with all fields to domain.Todo",
			input: TodoModel{
				ID:          "123",
				Title:       "Test Title",
				Description: "Test Description",
				Status:      "completed",
				DueDate:     &exampleDueDate,
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
			output: domain.Todo{
				ID:          "123",
				Title:       "Test Title",
				Description: "Test Description",
				Status:      domain.TodoStatusCompleted,
				DueDate:     &exampleDueDate,
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
		},
		{
			name: "should convert TodoModel without due date",
			input: TodoModel{
				ID:          "456",
				Title:       "Another Title",
				Description: "Another Description",
				Status:      "pending",
				DueDate:     nil,
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
			output: domain.Todo{
				ID:          "456",
				Title:       "Another Title",
				Description: "Another Description",
				Status:      domain.TodoStatusPending,
				DueDate:     nil,
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
		},
		{
			name: "should convert empty TodoModel",
			input: TodoModel{
				ID:          "",
				Title:       "",
				Description: "",
				Status:      "pending",
				DueDate:     nil,
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
			output: domain.Todo{
				ID:          "",
				Title:       "",
				Description: "",
				Status:      domain.TodoStatusPending,
				DueDate:     nil,
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := toDomain(tc.input)
			assert.Equal(t, tc.output, result)
		})
	}
}

func TestFromDomain(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	exampleDueDate, _ := time.Parse(time.DateOnly, "2024-12-31")

	testCases := []struct {
		name   string
		input  domain.Todo
		output TodoModel
	}{
		{
			name: "should convert domain.Todo with all fields to TodoModel",
			input: domain.Todo{
				ID:          "123",
				Title:       "Test Title",
				Description: "Test Description",
				Status:      domain.TodoStatusCompleted,
				DueDate:     &exampleDueDate,
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
			output: TodoModel{
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
			output: TodoModel{
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
			output: TodoModel{
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
			result := fromDomain(tc.input)
			assert.Equal(t, tc.output, result)
		})
	}
}
