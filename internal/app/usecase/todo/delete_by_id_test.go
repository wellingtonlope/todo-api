package todo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)

func TestDeleteByID_Handle(t *testing.T) {
	testCases := []struct {
		name  string
		store *deleteByIDStoreMock
		ctx   context.Context
		id    string
		err   error
	}{
		{
			name: "should fail when store fails",
			store: func() *deleteByIDStoreMock {
				m := new(deleteByIDStoreMock)
				m.On("DeleteByID", context.TODO(), "123").
					Return(assert.AnError).Once()
				return m
			}(),
			ctx: context.TODO(),
			id:  "123",
			err: usecase.NewError("fail to delete a todo by id", assert.AnError,
				usecase.ErrorTypeInternalError),
		},
		{
			name: "should delete trade by id",
			store: func() *deleteByIDStoreMock {
				m := new(deleteByIDStoreMock)
				m.On("DeleteByID", context.TODO(), "123").
					Return(nil).Once()
				return m
			}(),
			ctx: context.TODO(),
			id:  "123",
			err: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uc := todo.NewDeleteByID(tc.store)
			err := uc.Handle(tc.ctx, tc.id)
			assert.Equal(t, tc.err, err)
			tc.store.AssertExpectations(t)
		})
	}
}

type deleteByIDStoreMock struct {
	mock.Mock
}

func (m *deleteByIDStoreMock) DeleteByID(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
