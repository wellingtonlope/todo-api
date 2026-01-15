package gorm

import (
	"time"

	"github.com/wellingtonlope/todo-api/internal/domain"
)

type TodoModel struct {
	ID          string `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (TodoModel) TableName() string {
	return "todos"
}

func toDomain(m TodoModel) domain.Todo {
	return domain.Todo{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func fromDomain(t domain.Todo) TodoModel {
	return TodoModel{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
