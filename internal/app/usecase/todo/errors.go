package todo

import (
	"errors"
	"fmt"

	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

func notFoundError(id string, cause error) error {
	return usecase.NewError(
		fmt.Sprintf("todo not found with id %s", id),
		cause,
		usecase.ErrorTypeNotFound,
	)
}

func internalError(msg string, cause error) error {
	return usecase.NewError(msg, cause, usecase.ErrorTypeInternalError)
}

func badRequestError(msg string, cause error) error {
	return usecase.NewError(msg, cause, usecase.ErrorTypeBadRequest)
}

func isNotFound(err error) bool {
	return errors.Is(err, domain.ErrTodoNotFound)
}
