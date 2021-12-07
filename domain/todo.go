package domain

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Todo struct {
	Base        `bson:",inline"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewTodo(title, description string) (*Todo, error) {
	if title == "" {
		return nil, errors.New("title mustn't be empty")
	}

	return &Todo{
		Base: Base{
			ID:        uuid.NewV4().String(),
			CreatedAt: time.Now(),
		},
		Title:       title,
		Description: description,
	}, nil
}

func UpdateTodo(title, description string, todo Todo) (*Todo, error) {
	if title == "" {
		return nil, errors.New("title mustn't be empty")
	}

	todo.Title = title
	todo.Description = description
	todo.UpdatedAt = time.Now()
	return &todo, nil
}
