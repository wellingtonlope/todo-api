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

func TestGetAll_Handle(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	testCases := []struct {
		name   string
		store  *getAllStoreMock
		ctx    context.Context
		result []todo.TodoOutput
		err    error
	}{
		{
			name: "should fail when store fails",
			store: func() *getAllStoreMock {
				m := new(getAllStoreMock)
				m.On("GetAll", context.TODO()).
					Return([]domain.Todo{}, assert.AnError).Once()
				return m
			}(),
			ctx:    context.TODO(),
			result: []todo.TodoOutput{},
			err: usecase.NewError("fail to get all todos",
				assert.AnError, usecase.ErrorTypeInternalError),
		},
		{
			name: "should get all todos",
			store: func() *getAllStoreMock {
				m := new(getAllStoreMock)
				m.On("GetAll", context.TODO()).
					Return([]domain.Todo{
						{
							ID:          "123",
							Title:       "title",
							Description: "description",
							CreatedAt:   exampleDate,
							UpdatedAt:   exampleDate,
						},
					}, nil).Once()
				return m
			}(),
			ctx: context.TODO(),
			result: []todo.TodoOutput{
				{
					ID:          "123",
					Title:       "title",
					Description: "description",
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDate,
				},
			},
			err: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uc := todo.NewGetAll(tc.store)
			result, err := uc.Handle(tc.ctx)
			assert.Equal(t, tc.result, result)
			assert.Equal(t, tc.err, err)
			tc.store.AssertExpectations(t)
		})
	}
}

type getAllStoreMock struct {
	mock.Mock
}

func (m *getAllStoreMock) GetAll(ctx context.Context) ([]domain.Todo, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Todo), args.Error(1)
}
