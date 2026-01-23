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

func TestTodoGetAll_Handle(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	testCases := []struct {
		name           string
		getAll         *todoGetAllMock
		responseBody   string
		responseStatus int
		err            error
	}{
		{
			name: "should fail when get use case fails",
			getAll: func() *todoGetAllMock {
				m := new(todoGetAllMock)
				m.On("Handle", mock.Anything).Return([]todo.TodoOutput{}, usecase.AnError).Once()
				return m
			}(),
			responseBody:   "",
			responseStatus: http.StatusOK,
			err:            usecase.AnError,
		},
		{
			name: "should get all todos",
			getAll: func() *todoGetAllMock {
				m := new(todoGetAllMock)
				m.On("Handle", mock.Anything).Return([]todo.TodoOutput{
					{
						ID:          "123",
						Title:       "example title",
						Description: "example description",
						Status:      "pending",
						CreatedAt:   exampleDate,
						UpdatedAt:   exampleDate,
					},
				}, nil).Once()
				return m
			}(),
			responseBody:   `[{"id":"123","title":"example title","description":"example description","status":"pending","created_at":"2024-01-01T00:00:00Z","updated_at":"2024-01-01T00:00:00Z"}]`,
			responseStatus: http.StatusOK,
			err:            nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/todos", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := handler.NewTodoGetAll(tc.getAll)
			err := h.Handle(c)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.responseBody, strings.Trim(rec.Body.String(), "\n"))
			assert.Equal(t, tc.responseStatus, rec.Result().StatusCode)
		})
	}
}

func TestTodoGetAll_Path(t *testing.T) {
	h := handler.NewTodoGetAll(new(todoGetAllMock))
	assert.Equal(t, "/todos", h.Path())
}

func TestTodoGetAll_Method(t *testing.T) {
	h := handler.NewTodoGetAll(new(todoGetAllMock))
	assert.Equal(t, http.MethodGet, h.Method())
}

type todoGetAllMock struct {
	mock.Mock
}

func (m *todoGetAllMock) Handle(ctx context.Context) ([]todo.TodoOutput, error) {
	args := m.Called(ctx)
	return args.Get(0).([]todo.TodoOutput), args.Error(1)
}
