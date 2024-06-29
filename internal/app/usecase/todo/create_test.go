package todo_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/todo-api/internal/app/repository"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	"github.com/wellingtonlope/todo-api/internal/domain"
	"github.com/wellingtonlope/todo-api/pkg/clock"
)

func TestCreate_Handle(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	testCases := []struct {
		name       string
		repository *repository.TodoMock
		clock      *clock.ClientMock
		ctx        context.Context
		input      todo.CreateInput
		result     todo.CreateOutput
		err        error
	}{
		{
			name:       "should fail when input is invalid",
			repository: repository.NewTodoMock(),
			clock: func() *clock.ClientMock {
				m := clock.NewClientMock()
				m.On("Now").Return(exampleDate).Once()
				return m
			}(),
			ctx: context.TODO(),
			input: todo.CreateInput{
				Title:       "",
				Description: "example description",
			},
			result: todo.CreateOutput{},
			err: usecase.NewError(fmt.Errorf("%w: title", domain.ErrTodoInvalidInput).Error(),
				fmt.Errorf("%w: title", domain.ErrTodoInvalidInput), usecase.ErrorTypeBadRequest),
		},
		{
			name: "should fail when repository fails",
			repository: func() *repository.TodoMock {
				m := repository.NewTodoMock()
				m.On("Create", context.TODO(), domain.Todo{
					Title:       "example title",
					Description: "example description",
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDate,
				}).Return(domain.Todo{}, assert.AnError).Once()
				return m
			}(),
			clock: func() *clock.ClientMock {
				m := clock.NewClientMock()
				m.On("Now").Return(exampleDate).Once()
				return m
			}(),
			ctx: context.TODO(),
			input: todo.CreateInput{
				Title:       "example title",
				Description: "example description",
			},
			result: todo.CreateOutput{},
			err: usecase.NewError("fail to create a todo in the repository", assert.AnError,
				usecase.ErrorTypeInternalError),
		},
		{
			name: "should create a todo",
			repository: func() *repository.TodoMock {
				m := repository.NewTodoMock()
				m.On("Create", context.TODO(), domain.Todo{
					Title:       "example title",
					Description: "example description",
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDate,
				}).Return(domain.Todo{
					ID:          "123",
					Title:       "example title",
					Description: "example description",
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDate,
				}, nil).Once()
				return m
			}(),
			clock: func() *clock.ClientMock {
				m := clock.NewClientMock()
				m.On("Now").Return(exampleDate).Once()
				return m
			}(),
			ctx: context.TODO(),
			input: todo.CreateInput{
				Title:       "example title",
				Description: "example description",
			},
			result: todo.CreateOutput{
				ID:          "123",
				Title:       "example title",
				Description: "example description",
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
			err: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uc := todo.NewCreate(tc.repository, tc.clock)
			result, err := uc.Handle(tc.ctx, tc.input)
			assert.Equal(t, tc.result, result)
			assert.Equal(t, tc.err, err)
		})
	}
}
