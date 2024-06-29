package domain

import (
	"errors"
	"fmt"
	"time"
)

var ErrTodoInvalidInput = errors.New("todo invalid input")

type Todo struct {
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTodo(title, description string, date time.Time) (Todo, error) {
	if title == "" {
		return Todo{}, fmt.Errorf("%w: title", ErrTodoInvalidInput)
	}
	if date.IsZero() {
		return Todo{}, fmt.Errorf("%w: date", ErrTodoInvalidInput)
	}
	return Todo{
		Title:       title,
		Description: description,
		CreatedAt:   date,
		UpdatedAt:   date,
	}, nil
}

func (t Todo) Update(title, description string, date time.Time) (Todo, error) {
	if title == "" {
		return Todo{}, fmt.Errorf("%w: title", ErrTodoInvalidInput)
	}
	if date.IsZero() {
		return Todo{}, fmt.Errorf("%w: date", ErrTodoInvalidInput)
	}
	t.Title = title
	t.Description = description
	t.UpdatedAt = date
	return t, nil
}
