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

func TestList_Handle(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	pendingStatus := domain.TodoStatusPending
	completedStatus := domain.TodoStatusCompleted

	testCases := []struct {
		name   string
		store  *listStoreMock
		input  todo.ListInput
		result []todo.TodoOutput
		err    error
	}{
		{
			name: "should fail when store fails",
			store: func() *listStoreMock {
				m := new(listStoreMock)
				m.On("List", context.TODO(), mock.Anything).
					Return([]domain.Todo{}, assert.AnError).Once()
				return m
			}(),
			input:  todo.ListInput{},
			result: []todo.TodoOutput{},
			err: usecase.NewError("fail to list todos",
				assert.AnError, usecase.ErrorTypeInternalError),
		},
		{
			name: "should list all todos without filter",
			store: func() *listStoreMock {
				m := new(listStoreMock)
				m.On("List", context.TODO(), mock.Anything).
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
			input: todo.ListInput{},
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
		{
			name: "should list todos filtered by pending status",
			store: func() *listStoreMock {
				m := new(listStoreMock)
				m.On("List", context.TODO(), &pendingStatus).
					Return([]domain.Todo{
						{
							ID:        "123",
							Title:     "pending todo",
							Status:    domain.TodoStatusPending,
							CreatedAt: exampleDate,
							UpdatedAt: exampleDate,
						},
					}, nil).Once()
				return m
			}(),
			input: todo.ListInput{Status: &pendingStatus},
			result: []todo.TodoOutput{
				{
					ID:        "123",
					Title:     "pending todo",
					Status:    "pending",
					CreatedAt: exampleDate,
					UpdatedAt: exampleDate,
				},
			},
			err: nil,
		},
		{
			name: "should list todos filtered by completed status",
			store: func() *listStoreMock {
				m := new(listStoreMock)
				m.On("List", context.TODO(), &completedStatus).
					Return([]domain.Todo{
						{
							ID:        "456",
							Title:     "completed todo",
							Status:    domain.TodoStatusCompleted,
							CreatedAt: exampleDate,
							UpdatedAt: exampleDate,
						},
					}, nil).Once()
				return m
			}(),
			input: todo.ListInput{Status: &completedStatus},
			result: []todo.TodoOutput{
				{
					ID:        "456",
					Title:     "completed todo",
					Status:    "completed",
					CreatedAt: exampleDate,
					UpdatedAt: exampleDate,
				},
			},
			err: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uc := todo.NewList(tc.store)
			result, err := uc.Handle(context.TODO(), tc.input)
			assert.Equal(t, tc.result, result)
			assert.Equal(t, tc.err, err)
			tc.store.AssertExpectations(t)
		})
	}
}

type listStoreMock struct {
	mock.Mock
}

func (m *listStoreMock) List(ctx context.Context, status *domain.TodoStatus) ([]domain.Todo, error) {
	args := m.Called(ctx, status)
	return args.Get(0).([]domain.Todo), args.Error(1)
}
