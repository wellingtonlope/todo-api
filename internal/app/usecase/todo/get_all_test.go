package todo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/todo-api/internal/domain"
	"github.com/wellingtonlope/todo-api/internal/infra/memory"
)

func TestGetAll(t *testing.T) {
	t.Run("should get all todos", func(t *testing.T) {
		repo := memory.TodoRepository{}
		usecase := GetAll{
			todoRepository: &repo,
		}

		todo, _ := domain.NewTodo("title", "description", nil)
		_, _ = repo.Insert(*todo)
		_, _ = repo.Insert(*todo)

		input := GetAllInput{}

		output, err := usecase.Handle(input)

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, 2, len(*output))
	})
}
