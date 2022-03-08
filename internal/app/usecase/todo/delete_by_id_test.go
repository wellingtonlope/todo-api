package todo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/todo-api/internal/app/repository"
	"github.com/wellingtonlope/todo-api/internal/domain"
	"github.com/wellingtonlope/todo-api/internal/infra/memory"
)

func TestDeleteByID(t *testing.T) {
	t.Run("should delete a todo by id", func(t *testing.T) {
		repo := memory.TodoRepository{}
		usecase := DeleteById{
			todoRepository: &repo,
		}

		todo, _ := domain.NewTodo("title", "description", nil)
		todo, _ = repo.Insert(*todo)

		input := DeleteByIdInput{
			ID: todo.ID,
		}

		output, err := usecase.Handle(input)
		assert.Nil(t, err)
		assert.NotNil(t, output)

		todo, err = repo.GetByID(todo.ID)
		assert.Nil(t, todo)
		assert.NotNil(t, err)
		assert.Equal(t, repository.ErrTodoNotFound, err)
	})

	t.Run("shouldn't delete a todo by id", func(t *testing.T) {
		repo := memory.TodoRepository{}
		usecase := DeleteById{
			todoRepository: &repo,
		}

		input := DeleteByIdInput{
			ID: "invalid_id",
		}

		output, err := usecase.Handle(input)
		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Equal(t, repository.ErrTodoNotFound, err)
	})
}
