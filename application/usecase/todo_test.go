package usecase

import (
	"testing"

	"github.com/wellingtonlope/todo-api/application/dto"
	"github.com/wellingtonlope/todo-api/application/myerrors"
	"github.com/wellingtonlope/todo-api/framework/db/local"
)

func TestNew(t *testing.T) {
	expectedTitle, expectedDescription := "title 1", "description 1"
	repo := &local.TodoRepositoryLocal{}
	useCase := TodoUseCase{repo}

	t.Run("With title and description", func(t *testing.T) {
		todoDTO := dto.TodoNewDTO{Title: expectedTitle, Description: expectedDescription}
		got, _ := useCase.Insert(todoDTO)
		if got == nil {
			t.Error("Error: expected a Todo, but got nothing")
		}
	})

	t.Run("With empty title and description", func(t *testing.T) {
		_, err := useCase.Insert(dto.TodoNewDTO{})
		if err == nil && err.Type == myerrors.DOMAIN {
			t.Error("Error: expected a error, but got nothing")
		}
		if err.Type != myerrors.DOMAIN {
			t.Errorf("Error: expected a error type %q, but got %q", myerrors.DOMAIN, err.Type)
		}
	})
}

func TestUpdate(t *testing.T) {
	repo := &local.TodoRepositoryLocal{}
	useCase := TodoUseCase{repo}
	expectedTitle, expectedDescription := "title 1", "description 1"
	todoSaved, _ := useCase.Insert(dto.TodoNewDTO{Title: "title", Description: "description"})

	t.Run("Update with title and description", func(t *testing.T) {
		todoDTO := dto.TodoNewDTO{Title: expectedTitle, Description: expectedDescription}
		got, _ := useCase.Update(todoSaved.ID, todoDTO)
		if got == nil {
			t.Error("Error: expected a Todo, but got nothing")
		}
	})

	t.Run("With empty title and description", func(t *testing.T) {
		_, err := useCase.Update(todoSaved.ID, dto.TodoNewDTO{})
		if err == nil {
			t.Error("Error: expected a error, but got nothing")
		}
		if err != nil && err.Type != myerrors.DOMAIN {
			t.Errorf("Error: expected a error type %q, but got %q", myerrors.DOMAIN, err.Type)
		}
	})

	t.Run("With not existed id", func(t *testing.T) {
		todoDTO := dto.TodoNewDTO{Title: expectedTitle, Description: expectedDescription}
		_, err := useCase.Update("123", todoDTO)
		if err == nil {
			t.Error("Error: expected a error, but got nothing")
		}
		if err != nil && err.Type != myerrors.REGISTER_NOT_FOUND {
			t.Errorf("Error: expected a error type %q, but got %q", myerrors.REGISTER_NOT_FOUND, err.Type)
		}
	})

}

func TestGetById(t *testing.T) {
	repo := &local.TodoRepositoryLocal{}
	useCase := TodoUseCase{repo}
	todoSaved, _ := useCase.Insert(dto.TodoNewDTO{Title: "title", Description: "description"})

	t.Run("Get by an existent ID", func(t *testing.T) {
		todoGet, _ := useCase.GetById(todoSaved.ID)
		if todoGet == nil {
			t.Error("Error: expected a todo, but got nothing")
		}
	})

	t.Run("Get by a non-existent ID", func(t *testing.T) {
		_, err := useCase.GetById("non-existent")
		if err == nil {
			t.Error("Error: expected a error, but got nothing")
		}
		if err != nil && err.Type != myerrors.REGISTER_NOT_FOUND {
			t.Errorf("Error: expected a error type %q, but got %q", myerrors.REGISTER_NOT_FOUND, err.Type)
		}
	})

	t.Run("Get by a empty ID", func(t *testing.T) {
		_, err := useCase.GetById("")
		if err == nil {
			t.Error("Error: expected a error, but got nothing")
		}
		if err != nil && err.Type != myerrors.REGISTER_NOT_FOUND {
			t.Errorf("Error: expected a error type %q, but got %q", myerrors.REGISTER_NOT_FOUND, err.Type)
		}
	})
}

func TestGetAll(t *testing.T) {
	repo := &local.TodoRepositoryLocal{}
	useCase := TodoUseCase{repo}

	t.Run("Get all todos when database is empty", func(t *testing.T) {
		todos, _ := useCase.GetAll()
		if todos == nil {
			t.Error("Error: expected a todo slice, but got nothing")
		}
		if todos != nil && len(*todos) != 0 {
			t.Errorf("Error: expected a todo slice with 0 lenth, but got %d", len(*todos))
		}
	})

	useCase.Insert(dto.TodoNewDTO{Title: "title", Description: "description"})
	t.Run("Get all todos", func(t *testing.T) {
		todos, _ := useCase.GetAll()
		if todos == nil {
			t.Error("Error: expected a todo slice, but got nothing")
		}
		if todos != nil && len(*todos) != 1 {
			t.Errorf("Error: expected a todo slice with 1 lenth, but got %d", len(*todos))
		}
	})
}

func TestDelete(t *testing.T) {
	repo := &local.TodoRepositoryLocal{}
	useCase := TodoUseCase{repo}
	todoSaved, _ := useCase.Insert(dto.TodoNewDTO{Title: "title", Description: "description"})

	t.Run("Delete a existent todo", func(t *testing.T) {
		useCase.Delete(todoSaved.ID)
		todosAfter, _ := useCase.GetAll()
		if len(*todosAfter) > 0 {
			t.Errorf("Error: expected a empty slice, but got a slice with %d length", len(*todosAfter))
		}
	})

	t.Run("Delete a non-existent todo", func(t *testing.T) {
		err := useCase.Delete("non-existent")
		if err == nil {
			t.Error("Error: expected a error, but got nothing")
		}
		if err != nil && err.Type != myerrors.REGISTER_NOT_FOUND {
			t.Errorf("Error: expected a error type %q, but got %q", myerrors.REGISTER_NOT_FOUND, err.Type)
		}
	})

	t.Run("Delete a todo with empty ID", func(t *testing.T) {
		err := useCase.Delete("")
		if err == nil {
			t.Error("Error: expected a error, but got nothing")
		}
		if err != nil && err.Type != myerrors.REGISTER_NOT_FOUND {
			t.Errorf("Error: expected a error type %q, but got %q", myerrors.REGISTER_NOT_FOUND, err.Type)
		}
	})
}
