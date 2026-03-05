package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// ErrTodoInvalidInput is returned when the todo input is invalid.
var ErrTodoInvalidInput = errors.New("todo invalid input")

// TodoStatus represents the current status of a todo.
type TodoStatus string

const (
	// TodoStatusPending indicates the todo has not been completed yet.
	TodoStatusPending TodoStatus = "pending"
	// TodoStatusCompleted indicates the todo has been completed.
	TodoStatusCompleted TodoStatus = "completed"
)

// Todo represents a task or item to be done.
type Todo struct {
	ID          string
	Title       string
	Description string
	Status      TodoStatus
	DueDate     *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewTodo creates a new Todo with the given parameters.
// It validates that title is not empty and date is not zero.
// If dueDate is provided, it must be after date.
//
// Parameters:
//   - title: the todo title (required)
//   - description: the todo description (optional)
//   - date: the current timestamp (required, must not be zero)
//   - dueDate: optional deadline, must be after date if provided
//
// Returns:
//   - Todo: the created todo instance
//   - error: ErrTodoInvalidInput if validation fails
func NewTodo(title, description string, date time.Time, dueDate *time.Time) (Todo, error) {
	title = strings.TrimSpace(title)
	description = strings.TrimSpace(description)
	if title == "" {
		return Todo{}, fmt.Errorf("%w: title", ErrTodoInvalidInput)
	}
	if date.IsZero() {
		return Todo{}, fmt.Errorf("%w: date", ErrTodoInvalidInput)
	}
	if dueDate != nil && dueDate.Before(date) {
		return Todo{}, fmt.Errorf("%w: due date must be in the future", ErrTodoInvalidInput)
	}
	return Todo{
		Title:       title,
		Description: description,
		Status:      TodoStatusPending,
		DueDate:     dueDate,
		CreatedAt:   date,
		UpdatedAt:   date,
	}, nil
}

// Update modifies the todo with new values.
// It validates that title is not empty and date is not zero.
// If dueDate is provided, it must be after date.
//
// Parameters:
//   - title: the new todo title (required)
//   - description: the new todo description (optional)
//   - date: the current timestamp (required, must not be zero)
//   - dueDate: optional deadline, must be after date if provided
//
// Returns:
//   - Todo: the updated todo instance
//   - error: ErrTodoInvalidInput if validation fails
func (t Todo) Update(title, description string, date time.Time, dueDate *time.Time) (Todo, error) {
	title = strings.TrimSpace(title)
	description = strings.TrimSpace(description)
	if title == "" {
		return Todo{}, fmt.Errorf("%w: title", ErrTodoInvalidInput)
	}
	if date.IsZero() {
		return Todo{}, fmt.Errorf("%w: date", ErrTodoInvalidInput)
	}
	if dueDate != nil && dueDate.Before(date) {
		return Todo{}, fmt.Errorf("%w: due date must be in the future", ErrTodoInvalidInput)
	}
	t.Title = title
	t.Description = description
	t.DueDate = dueDate
	t.UpdatedAt = date
	return t, nil
}

// MarkAsCompleted marks the todo as completed with the given date.
//
// Parameters:
//   - date: the current timestamp
//
// Returns:
//   - Todo: the updated todo with status set to completed
func (t Todo) MarkAsCompleted(date time.Time) Todo {
	t.Status = TodoStatusCompleted
	t.UpdatedAt = date
	return t
}

// MarkAsPending marks the todo as pending with the given date.
//
// Parameters:
//   - date: the current timestamp
//
// Returns:
//   - Todo: the updated todo with status set to pending
func (t Todo) MarkAsPending(date time.Time) Todo {
	t.Status = TodoStatusPending
	t.UpdatedAt = date
	return t
}
