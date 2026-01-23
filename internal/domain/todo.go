package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var ErrTodoInvalidInput = errors.New("todo invalid input")

type Todo struct {
	ID          string
	Title       string
	Description string
	DueDate     *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

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
		DueDate:     dueDate,
		CreatedAt:   date,
		UpdatedAt:   date,
	}, nil
}

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
