package todo_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	"github.com/wellingtonlope/todo-api/internal/domain"
	"github.com/wellingtonlope/todo-api/pkg/clock"
)

func TestUpdate_Handle(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	exampleDateUpdated, _ := time.Parse(time.DateOnly, "2024-01-02")
	testCases := []struct {
		name        string
		updateStore *updateStoreMock
		clock       *clock.ClientMock
		ctx         context.Context
		input       todo.UpdateInput
		result      todo.UpdateOutput
		err         error
	}{
		{
			name: "should fail when todo not found",
			updateStore: func() *updateStoreMock {
				m := new(updateStoreMock)
				m.On("GetByID", context.TODO(), "123").
					Return(domain.Todo{}, todo.ErrGetByIDStoreNotFound).Once()
				return m
			}(),
			clock: clock.NewClientMock(),
			ctx:   context.TODO(),
			input: todo.UpdateInput{
				ID:          "123",
				Title:       "example title updated",
				Description: "example description updated",
			},
			result: todo.UpdateOutput{},
			err: usecase.NewError("todo not found with id 123",
				todo.ErrGetByIDStoreNotFound, usecase.ErrorTypeNotFound),
		},
		{
			name: "should fail when get by id fails",
			updateStore: func() *updateStoreMock {
				m := new(updateStoreMock)
				m.On("GetByID", context.TODO(), "123").
					Return(domain.Todo{}, assert.AnError).Once()
				return m
			}(),
			clock: clock.NewClientMock(),
			ctx:   context.TODO(),
			input: todo.UpdateInput{
				ID:          "123",
				Title:       "example title updated",
				Description: "example description updated",
			},
			result: todo.UpdateOutput{},
			err: usecase.NewError("fail to get a todo by id",
				assert.AnError, usecase.ErrorTypeInternalError),
		},
		{
			name: "should fail when input is invalid",
			updateStore: func() *updateStoreMock {
				m := new(updateStoreMock)
				m.On("GetByID", context.TODO(), "123").
					Return(domain.Todo{
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
				m.On("Now").Return(exampleDateUpdated).Once()
				return m
			}(),
			ctx: context.TODO(),
			input: todo.UpdateInput{
				ID:          "123",
				Title:       "",
				Description: "example description updated",
			},
			result: todo.UpdateOutput{},
			err: usecase.NewError(fmt.Errorf("%w: title", domain.ErrTodoInvalidInput).Error(),
				fmt.Errorf("%w: title", domain.ErrTodoInvalidInput), usecase.ErrorTypeBadRequest),
		},
		{
			name: "should fail when update todo not found",
			updateStore: func() *updateStoreMock {
				m := new(updateStoreMock)
				m.On("GetByID", context.TODO(), "123").
					Return(domain.Todo{
						ID:          "123",
						Title:       "example title",
						Description: "example description",
						CreatedAt:   exampleDate,
						UpdatedAt:   exampleDate,
					}, nil).Once()
				m.On("Update", context.TODO(), domain.Todo{
					ID:          "123",
					Title:       "example title updated",
					Description: "example description updated",
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDateUpdated,
				}).Return(todo.ErrUpdateStoreNotFound).Once()
				return m
			}(),
			clock: func() *clock.ClientMock {
				m := clock.NewClientMock()
				m.On("Now").Return(exampleDateUpdated).Once()
				return m
			}(),
			ctx: context.TODO(),
			input: todo.UpdateInput{
				ID:          "123",
				Title:       "example title updated",
				Description: "example description updated",
			},
			result: todo.UpdateOutput{},
			err: usecase.NewError("todo not found with id 123",
				todo.ErrUpdateStoreNotFound, usecase.ErrorTypeNotFound),
		},
		{
			name: "should fail when update store fails",
			updateStore: func() *updateStoreMock {
				m := new(updateStoreMock)
				m.On("GetByID", context.TODO(), "123").
					Return(domain.Todo{
						ID:          "123",
						Title:       "example title",
						Description: "example description",
						CreatedAt:   exampleDate,
						UpdatedAt:   exampleDate,
					}, nil).Once()
				m.On("Update", context.TODO(), domain.Todo{
					ID:          "123",
					Title:       "example title updated",
					Description: "example description updated",
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDateUpdated,
				}).Return(assert.AnError).Once()
				return m
			}(),
			clock: func() *clock.ClientMock {
				m := clock.NewClientMock()
				m.On("Now").Return(exampleDateUpdated).Once()
				return m
			}(),
			ctx: context.TODO(),
			input: todo.UpdateInput{
				ID:          "123",
				Title:       "example title updated",
				Description: "example description updated",
			},
			result: todo.UpdateOutput{},
			err: usecase.NewError("fail to update a todo in the store", assert.AnError,
				usecase.ErrorTypeInternalError),
		},
		{
			name: "should update todo",
			updateStore: func() *updateStoreMock {
				m := new(updateStoreMock)
				m.On("GetByID", context.TODO(), "123").
					Return(domain.Todo{
						ID:          "123",
						Title:       "example title",
						Description: "example description",
						CreatedAt:   exampleDate,
						UpdatedAt:   exampleDate,
					}, nil).Once()
				m.On("Update", context.TODO(), domain.Todo{
					ID:          "123",
					Title:       "example title updated",
					Description: "example description updated",
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDateUpdated,
				}).Return(nil).Once()
				return m
			}(),
			clock: func() *clock.ClientMock {
				m := clock.NewClientMock()
				m.On("Now").Return(exampleDateUpdated).Once()
				return m
			}(),
			ctx: context.TODO(),
			input: todo.UpdateInput{
				ID:          "123",
				Title:       "example title updated",
				Description: "example description updated",
			},
			result: todo.UpdateOutput{
				ID:          "123",
				Title:       "example title updated",
				Description: "example description updated",
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDateUpdated,
			},
			err: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uc := todo.NewUpdate(tc.updateStore, tc.clock)
			result, err := uc.Handle(tc.ctx, tc.input)
			assert.Equal(t, tc.result, result)
			assert.Equal(t, tc.err, err)
			tc.updateStore.AssertExpectations(t)
			tc.clock.AssertExpectations(t)
		})
	}
}

type updateStoreMock struct {
	mock.Mock
}

func (m *updateStoreMock) GetByID(ctx context.Context, id string) (domain.Todo, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Todo), args.Error(1)
}

func (m *updateStoreMock) Update(ctx context.Context, todo domain.Todo) error {
	args := m.Called(ctx, todo)
	return args.Error(0)
}
