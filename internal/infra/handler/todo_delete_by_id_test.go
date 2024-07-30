package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/infra/handler"
)

func TestTodoDeleteByID_Handle(t *testing.T) {
	testCases := []struct {
		name           string
		deleteByID     *todoDeleteByIDMock
		pathID         string
		responseStatus int
		err            error
	}{
		{
			name: "should fail when delete use case fails",
			deleteByID: func() *todoDeleteByIDMock {
				m := new(todoDeleteByIDMock)
				m.On("Handle", mock.Anything, "123").Return(usecase.AnError).Once()
				return m
			}(),
			pathID:         "123",
			responseStatus: http.StatusOK,
			err:            usecase.AnError,
		},
		{
			name: "should delete a todo by id",
			deleteByID: func() *todoDeleteByIDMock {
				m := new(todoDeleteByIDMock)
				m.On("Handle", mock.Anything, "123").Return(nil).Once()
				return m
			}(),
			pathID:         "123",
			responseStatus: http.StatusNoContent,
			err:            nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/todos/:id")
			c.SetParamNames("id")
			c.SetParamValues(tc.pathID)
			h := handler.NewTodoDeleteByID(tc.deleteByID)
			err := h.Handle(c)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.responseStatus, rec.Result().StatusCode)
		})
	}
}

func TestTodoDeleteByID_Path(t *testing.T) {
	h := handler.NewTodoDeleteByID(new(todoDeleteByIDMock))
	assert.Equal(t, "/todos/:id", h.Path())
}

func TestTodoDeleteByID_Method(t *testing.T) {
	h := handler.NewTodoDeleteByID(new(todoDeleteByIDMock))
	assert.Equal(t, http.MethodDelete, h.Method())
}

type todoDeleteByIDMock struct {
	mock.Mock
}

func (m *todoDeleteByIDMock) Handle(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
