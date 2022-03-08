package todo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/todo-api/internal/domain"
	"github.com/wellingtonlope/todo-api/internal/infra/memory"
)

func TestInsert(t *testing.T) {
	t.Run("should insert a todo", func(t *testing.T) {
		repository := memory.TodoRepository{}
		usecase := Insert{
			todoRepository: &repository,
		}
		today := time.Now()

		input := InsertInput{
			Title:       "title",
			Description: "description",
			CreatedDate: &today,
		}

		output, err := usecase.Handle(input)

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.NotEmpty(t, output.ID)
		assert.Equal(t, input.Title, output.Title)
		assert.Equal(t, input.Description, output.Description)
		assert.Equal(t, *input.CreatedDate, *output.CreatedDate)

		todoGet, _ := repository.GetByID(output.ID)
		assert.NotNil(t, todoGet)
		assert.Equal(t, output.ID, todoGet.ID)
		assert.Equal(t, input.Title, todoGet.Title)
		assert.Equal(t, input.Description, todoGet.Description)
		assert.Equal(t, *input.CreatedDate, *todoGet.CreatedDate)
	})

	t.Run("shouldn't insert a todo because invalid input", func(t *testing.T) {
		repo := memory.TodoRepository{}
		usecase := Insert{
			todoRepository: &repo,
		}

		input := InsertInput{
			Title:       "",
			Description: "",
		}

		_, err := usecase.Handle(input)
		assert.NotNil(t, err)
		assert.Equal(t, domain.ErrTitleRequired, err)
	})
}
