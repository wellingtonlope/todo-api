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

func TestGetByID_Handle(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	testCases := []struct {
		name   string
		store  *getByIDStoreMock
		ctx    context.Context
		id     string
		result todo.GetByIDOutput
		err    error
	}{
		{
			name: "should fail when todo not found",
			store: func() *getByIDStoreMock {
				m := new(getByIDStoreMock)
				m.On("GetByID", context.TODO(), "123").
					Return(domain.Todo{}, todo.ErrGetByIDStoreNotFound).Once()
				return m
			}(),
			ctx:    context.TODO(),
			id:     "123",
			result: todo.GetByIDOutput{},
			err: usecase.NewError("todo not found with id 123",
				todo.ErrGetByIDStoreNotFound, usecase.ErrorTypeNotFound),
		},
		{
			name: "should fail when store fails",
			store: func() *getByIDStoreMock {
				m := new(getByIDStoreMock)
				m.On("GetByID", context.TODO(), "123").
					Return(domain.Todo{}, assert.AnError).Once()
				return m
			}(),
			ctx:    context.TODO(),
			id:     "123",
			result: todo.GetByIDOutput{},
			err: usecase.NewError("fail to get a todo by id",
				assert.AnError, usecase.ErrorTypeInternalError),
		},
		{
			name: "should get a todo by id",
			store: func() *getByIDStoreMock {
				m := new(getByIDStoreMock)
				m.On("GetByID", context.TODO(), "123").
					Return(domain.Todo{
						ID:          "123",
						Title:       "title",
						Description: "description",
						CreatedAt:   exampleDate,
						UpdatedAt:   exampleDate,
					}, nil).Once()
				return m
			}(),
			ctx: context.TODO(),
			id:  "123",
			result: todo.GetByIDOutput{
				ID:          "123",
				Title:       "title",
				Description: "description",
				CreatedAt:   exampleDate,
				UpdatedAt:   exampleDate,
			},
			err: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uc := todo.NewGetByID(tc.store)
			result, err := uc.Handle(tc.ctx, tc.id)
			assert.Equal(t, tc.result, result)
			assert.Equal(t, tc.err, err)
			tc.store.AssertExpectations(t)
		})
	}
}

type getByIDStoreMock struct {
	mock.Mock
}

func (m *getByIDStoreMock) GetByID(ctx context.Context, id string) (domain.Todo, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Todo), args.Error(1)
}
