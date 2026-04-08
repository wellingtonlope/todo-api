package todo

import (
	"context"

	"github.com/wellingtonlope/todo-api/internal/domain"
)

type TodoUpdater interface {
	GetByID(context.Context, string) (domain.Todo, error)
	Update(context.Context, domain.Todo) (domain.Todo, error)
}
