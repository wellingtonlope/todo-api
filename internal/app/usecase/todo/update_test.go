package todo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/todo-api/internal/app/repository"
	"github.com/wellingtonlope/todo-api/internal/domain"
	"github.com/wellingtonlope/todo-api/internal/infra/memory"
)

func TestUpdate(t *testing.T) {
	t.Run("should update a todo", func(t *testing.T) {
		repo := memory.TodoRepository{}
		usecase := Update{
			todoRepository: &repo,
		}

		expectedCreatedDate := time.Now()

		todo, _ := domain.NewTodo("title", "description", &expectedCreatedDate)
		todo, _ = repo.Insert(*todo)

		expectedTitle := "title_updated"
		expectedDescription := "description_updated"
		expectedUpdatedDate := time.Now()

		input := UpdateInput{
			ID:          todo.ID,
			Title:       expectedTitle,
			Description: expectedDescription,
			UpdatedDate: &expectedUpdatedDate,
		}

		output, err := usecase.Handle(input)

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, todo.ID, output.ID)
		assert.Equal(t, expectedTitle, output.Title)
		assert.Equal(t, expectedDescription, output.Description)
		assert.Equal(t, expectedCreatedDate, *output.CreatedDate)
		assert.Equal(t, expectedUpdatedDate, *output.UpdatedDate)

		todo, _ = repo.GetByID(todo.ID)
		assert.Equal(t, output.ID, todo.ID)
		assert.Equal(t, expectedTitle, todo.Title)
		assert.Equal(t, expectedDescription, todo.Description)
		assert.Equal(t, expectedCreatedDate, *todo.CreatedDate)
		assert.Equal(t, expectedUpdatedDate, *todo.UpdatedDate)
	})

	t.Run("shouldn't update a todo because todo not exists", func(t *testing.T) {
		repo := memory.TodoRepository{}
		usecase := Update{
			todoRepository: &repo,
		}

		input := UpdateInput{
			ID: "invalid_id",
		}

		output, err := usecase.Handle(input)

		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Equal(t, repository.ErrTodoNotFound, err)
	})

	t.Run("shouldn't update a todo because invalid input", func(t *testing.T) {
		repo := memory.TodoRepository{}
		usecase := Update{
			todoRepository: &repo,
		}

		todo, _ := domain.NewTodo("title", "description", nil)
		todo, _ = repo.Insert(*todo)

		input := UpdateInput{
			ID: todo.ID,
		}

		output, err := usecase.Handle(input)

		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Equal(t, domain.ErrTitleRequired, err)
	})
}
