package todo_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

func TestComplete_Handle(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	exampleDateUpdated, _ := time.Parse(time.DateOnly, "2024-01-02")
	testCases := []struct {
		name          string
		completeStore *completeStoreMock
		clock         *usecase.ClockMock
		ctx           context.Context
		input         todo.CompleteInput
		result        todo.TodoOutput
		err           error
	}{
		{
			name: "should fail when todo not found",
			completeStore: func() *completeStoreMock {
				m := new(completeStoreMock)
				m.On("GetByID", context.TODO(), "123").
					Return(domain.Todo{}, todo.ErrGetByIDStoreNotFound).Once()
				return m
			}(),
			clock: func() *usecase.ClockMock {
				m := usecase.NewClockMock()
				return m
			}(),
			ctx: context.TODO(),
			input: todo.CompleteInput{
				ID: "123",
			},
			result: todo.TodoOutput{},
			err: usecase.NewError("todo not found with id 123",
				todo.ErrGetByIDStoreNotFound, usecase.ErrorTypeNotFound),
		},
		{
			name: "should fail when repository fails",
			completeStore: func() *completeStoreMock {
				m := new(completeStoreMock)
				m.On("GetByID", context.TODO(), "123").
					Return(domain.Todo{}, assert.AnError).Once()
				return m
			}(),
			clock: func() *usecase.ClockMock {
				m := usecase.NewClockMock()
				return m
			}(),
			ctx: context.TODO(),
			input: todo.CompleteInput{
				ID: "123",
			},
			result: todo.TodoOutput{},
			err: usecase.NewError("fail to get a todo by id",
				assert.AnError, usecase.ErrorTypeInternalError),
		},
		{
			name: "should fail when update fails",
			completeStore: func() *completeStoreMock {
				m := new(completeStoreMock)
				m.On("GetByID", context.TODO(), "123").
					Return(domain.Todo{
						ID:          "123",
						Title:       "example title",
						Description: "example description",
						Status:      domain.TodoStatusPending,
						CreatedAt:   exampleDate,
						UpdatedAt:   exampleDate,
					}, nil).Once()
				m.On("Update", context.TODO(), domain.Todo{
					ID:          "123",
					Title:       "example title",
					Description: "example description",
					Status:      domain.TodoStatusCompleted,
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDateUpdated,
				}).Return(domain.Todo{}, assert.AnError).Once()
				return m
			}(),
			clock: func() *usecase.ClockMock {
				m := usecase.NewClockMock()
				m.On("Now").Return(exampleDateUpdated).Once()
				return m
			}(),
			ctx: context.TODO(),
			input: todo.CompleteInput{
				ID: "123",
			},
			result: todo.TodoOutput{},
			err: usecase.NewError("fail to update a todo in the store", assert.AnError,
				usecase.ErrorTypeInternalError),
		},
		{
			name: "should complete a todo",
			completeStore: func() *completeStoreMock {
				m := new(completeStoreMock)
				m.On("GetByID", context.TODO(), "123").
					Return(domain.Todo{
						ID:          "123",
						Title:       "example title",
						Description: "example description",
						Status:      domain.TodoStatusPending,
						CreatedAt:   exampleDate,
						UpdatedAt:   exampleDate,
					}, nil).Once()
				m.On("Update", context.TODO(), domain.Todo{
					ID:          "123",
					Title:       "example title",
					Description: "example description",
					Status:      domain.TodoStatusCompleted,
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDateUpdated,
				}).Return(domain.Todo{
					ID:          "123",
					Title:       "example title",
					Description: "example description",
					Status:      domain.TodoStatusCompleted,
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDateUpdated,
				}, nil).Once()
				return m
			}(),
			clock: func() *usecase.ClockMock {
				m := usecase.NewClockMock()
				m.On("Now").Return(exampleDateUpdated).Once()
				return m
			}(),
			ctx: context.TODO(),
			input: todo.CompleteInput{
				ID: "123",
			},
			result: todo.TodoOutput{
				ID:          "123",
				Title:       "example title",
				Description: "example description",
				Status:      "completed",
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDateUpdated,
			},
			err: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uc := todo.NewComplete(tc.completeStore, tc.clock)
			result, err := uc.Handle(tc.ctx, tc.input)
			assert.Equal(t, tc.result, result)
			assert.Equal(t, tc.err, err)
		})
	}
}

type completeStoreMock struct {
	mock.Mock
}

func (m *completeStoreMock) GetByID(ctx context.Context, id string) (domain.Todo, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Todo), args.Error(1)
}

func (m *completeStoreMock) Update(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	args := m.Called(ctx, todo)
	return args.Get(0).(domain.Todo), args.Error(1)
}
