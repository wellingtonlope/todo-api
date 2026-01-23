package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	"github.com/wellingtonlope/todo-api/internal/infra/handler"
)

func TestTodoComplete_Handle(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	testCases := []struct {
		name           string
		complete       *todoCompleteMock
		responseBody   string
		responseStatus int
		err            error
	}{
		{
			name: "should fail when complete use case fails",
			complete: func() *todoCompleteMock {
				m := new(todoCompleteMock)
				m.On("Handle", mock.Anything, todo.CompleteInput{
					ID: "123",
				}).Return(todo.TodoOutput{}, usecase.NewError("todo not found", assert.AnError, usecase.ErrorTypeNotFound)).Once()
				return m
			}(),
			responseBody:   "",
			responseStatus: http.StatusOK,
			err:            usecase.NewError("todo not found", assert.AnError, usecase.ErrorTypeNotFound),
		},
		{
			name: "should complete a todo",
			complete: func() *todoCompleteMock {
				m := new(todoCompleteMock)
				m.On("Handle", mock.Anything, todo.CompleteInput{
					ID: "123",
				}).Return(todo.TodoOutput{
					ID:          "123",
					Title:       "example title",
					Description: "example description",
					Status:      "completed",
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDate,
				}, nil).Once()
				return m
			}(),
			responseBody:   `{"id":"123","title":"example title","description":"example description","status":"completed","created_at":"2024-01-01T00:00:00Z","updated_at":"2024-01-01T00:00:00Z"}`,
			responseStatus: http.StatusOK,
			err:            nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/todos/:id/complete")
			c.SetParamNames("id")
			c.SetParamValues("123")

			h := handler.NewTodoComplete(tc.complete)
			err := h.Handle(c)

			if tc.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.err, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.responseStatus, rec.Code)
				assert.JSONEq(t, tc.responseBody, rec.Body.String())
			}
		})
	}
}

type todoCompleteMock struct {
	mock.Mock
}

func (m *todoCompleteMock) Handle(ctx context.Context, input todo.CompleteInput) (todo.TodoOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(todo.TodoOutput), args.Error(1)
}
