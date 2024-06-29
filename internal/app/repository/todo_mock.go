package repository

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/wellingtonlope/todo-api/internal/domain"
)

type TodoMock struct {
	mock.Mock
}

func NewTodoMock() *TodoMock {
	return new(TodoMock)
}

func (m *TodoMock) GetAll(ctx context.Context) ([]domain.Todo, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Todo), args.Error(1)
}

func (m *TodoMock) GetByID(ctx context.Context, id string) (domain.Todo, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Todo), args.Error(1)
}

func (m *TodoMock) Create(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	args := m.Called(ctx, todo)
	return args.Get(0).(domain.Todo), args.Error(1)
}

func (m *TodoMock) Update(ctx context.Context, todo domain.Todo) error {
	args := m.Called(ctx, todo)
	return args.Error(0)
}

func (m *TodoMock) DeleteByID(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
