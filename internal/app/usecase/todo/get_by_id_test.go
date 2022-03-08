package todo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/todo-api/internal/app/repository"
	"github.com/wellingtonlope/todo-api/internal/domain"
	"github.com/wellingtonlope/todo-api/internal/infra/memory"
)

func TestGetById(t *testing.T) {
	t.Run("should get todo by id", func(t *testing.T) {
		repo := memory.TodoRepository{}
		usecase := GetById{
			todoRepository: &repo,
		}

		expectedTitle := "title"
		expectedDescription := "description"
		expectedCreatedDate := time.Now()
		expectedUpdatedDate := time.Now()

		todo, _ := domain.NewTodo(expectedTitle, expectedDescription, &expectedCreatedDate)
		todo, _ = repo.Insert(*todo)
		todo.Update(expectedTitle, expectedDescription, &expectedUpdatedDate)
		todo, _ = repo.Update(*todo)

		input := GetByIdInput{
			ID: todo.ID,
		}

		output, err := usecase.Handle(input)

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, todo.ID, output.ID)
		assert.Equal(t, expectedTitle, output.Title)
		assert.Equal(t, expectedDescription, output.Description)
		assert.Equal(t, expectedCreatedDate, *output.CreatedDate)
		assert.Equal(t, expectedUpdatedDate, *output.UpdatedDate)
	})

	t.Run("shouldn't get todo by id", func(t *testing.T) {
		repo := memory.TodoRepository{}
		usecase := GetById{
			todoRepository: &repo,
		}

		input := GetByIdInput{
			ID: "invalid_id",
		}

		output, err := usecase.Handle(input)

		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Equal(t, repository.ErrTodoNotFound, err)
	})
}
