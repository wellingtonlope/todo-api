package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	"github.com/wellingtonlope/todo-api/internal/infra/handler"
)

func TestTodoGetByID_Handle(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	testCases := []struct {
		name           string
		getByID        *todoGetByIDMock
		pathID         string
		responseBody   string
		responseStatus int
		err            error
	}{
		{
			name: "should fail when get use case fails",
			getByID: func() *todoGetByIDMock {
				m := new(todoGetByIDMock)
				m.On("Handle", mock.Anything, "123").Return(todo.GetByIDOutput{}, usecase.AnError).Once()
				return m
			}(),
			pathID:         "123",
			responseBody:   "",
			responseStatus: http.StatusOK,
			err:            usecase.AnError,
		},
		{
			name: "should get a todo by id",
			getByID: func() *todoGetByIDMock {
				m := new(todoGetByIDMock)
				m.On("Handle", mock.Anything, "123").Return(todo.GetByIDOutput{
					ID:          "123",
					Title:       "example title",
					Description: "example description",
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDate,
				}, nil).Once()
				return m
			}(),
			pathID:         "123",
			responseBody:   `{"id":"123","title":"example title","description":"example description","created_at":"2024-01-01T00:00:00Z","updated_at":"2024-01-01T00:00:00Z"}`,
			responseStatus: http.StatusOK,
			err:            nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/todos/:id")
			c.SetParamNames("id")
			c.SetParamValues(tc.pathID)
			h := handler.NewTodoGetByID(tc.getByID)
			err := h.Handle(c)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.responseBody, strings.Trim(rec.Body.String(), "\n"))
			assert.Equal(t, tc.responseStatus, rec.Result().StatusCode)
		})
	}
}

func TestTodoGetByID_Path(t *testing.T) {
	h := handler.NewTodoGetByID(new(todoGetByIDMock))
	assert.Equal(t, "/todos/:id", h.Path())
}

func TestTodoGetByID_Method(t *testing.T) {
	h := handler.NewTodoGetByID(new(todoGetByIDMock))
	assert.Equal(t, http.MethodGet, h.Method())
}

type todoGetByIDMock struct {
	mock.Mock
}

func (m *todoGetByIDMock) Handle(ctx context.Context, id string) (todo.GetByIDOutput, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(todo.GetByIDOutput), args.Error(1)
}
