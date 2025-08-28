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
)

func TestCreate_Handle(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	testCases := []struct {
		name        string
		createStore *createStoreMock
		clock       *usecase.ClockMock
		ctx         context.Context
		input       todo.CreateInput
		result      todo.CreateOutput
		err         error
	}{
		{
			name:        "should fail when input is invalid",
			createStore: new(createStoreMock),
			clock: func() *usecase.ClockMock {
				m := usecase.NewClockMock()
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
			createStore: func() *createStoreMock {
				m := new(createStoreMock)
				m.On("Create", context.TODO(), domain.Todo{
					Title:       "example title",
					Description: "example description",
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDate,
				}).Return(domain.Todo{}, assert.AnError).Once()
				return m
			}(),
			clock: func() *usecase.ClockMock {
				m := usecase.NewClockMock()
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
			createStore: func() *createStoreMock {
				m := new(createStoreMock)
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
			clock: func() *usecase.ClockMock {
				m := usecase.NewClockMock()
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
			uc := todo.NewCreate(tc.createStore, tc.clock)
			result, err := uc.Handle(tc.ctx, tc.input)
			assert.Equal(t, tc.result, result)
			assert.Equal(t, tc.err, err)
			tc.createStore.AssertExpectations(t)
			tc.clock.AssertExpectations(t)
		})
	}
}

type createStoreMock struct {
	mock.Mock
}

func (m *createStoreMock) Create(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	args := m.Called(ctx, todo)
	return args.Get(0).(domain.Todo), args.Error(1)
}
