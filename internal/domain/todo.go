package domain

import (
	"errors"
	"time"
)

var (
	ErrTitleRequired = errors.New("title is required")
)

type Todo struct {
	ID          string
	Title       string
	Description string
	CreatedDate *time.Time
	UpdatedDate *time.Time
}

func NewTodo(title, description string, createdDate *time.Time) (*Todo, error) {
	if title == "" {
		return nil, ErrTitleRequired
	}
	if createdDate == nil {
		today := time.Now()
		createdDate = &today
	}

	return &Todo{
		Title:       title,
		Description: description,
		CreatedDate: createdDate,
	}, nil
}

func (todo *Todo) Update(title, description string, updatedDate *time.Time) error {
	if title == "" {
		return ErrTitleRequired
	}
	if updatedDate == nil {
		today := time.Now()
		updatedDate = &today
	}

	todo.Title = title
	todo.Description = description
	todo.UpdatedDate = updatedDate

	return nil
}
