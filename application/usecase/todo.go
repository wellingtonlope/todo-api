package usecase

import (
	"errors"
	"log"

	"github.com/wellingtonlope/todo-api/application/dto"
	"github.com/wellingtonlope/todo-api/application/myerrors"
	"github.com/wellingtonlope/todo-api/application/repositories"
	"github.com/wellingtonlope/todo-api/domain"
)

type TodoUseCase struct {
	TodoRepository repositories.TodoRepository
}

func (uc *TodoUseCase) Insert(todoDTO dto.TodoNewDTO) (*domain.Todo, *myerrors.Error) {
	todo, err := domain.NewTodo(todoDTO.Title, todoDTO.Description)
	if err != nil {
		log.Printf("Error during a creation of a todo: %v", err)
		return nil, myerrors.NewError(err, myerrors.DOMAIN)
	}

	todoSaved, myerr := uc.TodoRepository.Insert(todo)
	if myerr != nil {
		log.Printf("Error during a creation of a todo: %v", myerr)
		return nil, myerr
	}

	return todoSaved, nil
}

func (uc *TodoUseCase) Update(id string, todoDTO dto.TodoNewDTO) (*domain.Todo, *myerrors.Error) {
	todoGet, myerr := uc.GetById(id)
	if myerr != nil {
		log.Printf("Error during a update of a todo: %v", myerr)
		return nil, myerr
	}

	todoToUpdate, err := domain.UpdateTodo(todoDTO.Title, todoDTO.Description, *todoGet)
	if err != nil {
		log.Printf("Error during a update of a todo: %v", err)
		return nil, myerrors.NewError(err, myerrors.DOMAIN)
	}

	todoUpdated, myerr := uc.TodoRepository.Update(todoToUpdate)
	if myerr != nil {
		log.Printf("Error during a update of a todo: %v", myerr)
		return nil, myerr
	}

	return todoUpdated, nil
}

func (uc *TodoUseCase) GetById(id string) (*domain.Todo, *myerrors.Error) {
	if id == "" {
		return nil, myerrors.NewError(errors.New("ID muns't be empty"), myerrors.REGISTER_NOT_FOUND)
	}

	todoGet, err := uc.TodoRepository.GetById(id)
	if err != nil {
		log.Printf("Error during a get todo by ID: %v", err)
		return nil, err
	}

	return todoGet, nil
}

func (uc *TodoUseCase) GetAll() (*[]domain.Todo, *myerrors.Error) {
	todos, err := uc.TodoRepository.GetAll()
	if err != nil {
		log.Printf("Error during get all todos: %v", err)
		return nil, err
	}

	return todos, nil
}

func (uc *TodoUseCase) Delete(id string) *myerrors.Error {
	err := uc.TodoRepository.Delete(id)
	if err != nil {
		log.Printf("Error during delete a todo: %v", err)
		return err
	}

	return nil
}
