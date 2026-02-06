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
	"github.com/wellingtonlope/todo-api/internal/domain"
	"github.com/wellingtonlope/todo-api/internal/infra/handler"
)

func TestTodoList_Handle(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	pendingStatus := domain.TodoStatusPending
	completedStatus := domain.TodoStatusCompleted

	testCases := []struct {
		name           string
		list           *todoListMock
		queryParams    string
		responseBody   string
		responseStatus int
		err            error
	}{
		{
			name: "should fail when list use case fails",
			list: func() *todoListMock {
				m := new(todoListMock)
				m.On("Handle", mock.Anything, todo.ListInput{}).Return([]todo.TodoOutput{}, usecase.AnError).Once()
				return m
			}(),
			queryParams:    "",
			responseBody:   "",
			responseStatus: http.StatusOK,
			err:            usecase.AnError,
		},
		{
			name: "should list all todos without filter",
			list: func() *todoListMock {
				m := new(todoListMock)
				m.On("Handle", mock.Anything, todo.ListInput{}).Return([]todo.TodoOutput{
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
			queryParams:    "",
			responseBody:   `[{"id":"123","title":"example title","description":"example description","status":"pending","created_at":"2024-01-01T00:00:00Z","updated_at":"2024-01-01T00:00:00Z"}]`,
			responseStatus: http.StatusOK,
			err:            nil,
		},
		{
			name: "should list todos filtered by pending status",
			list: func() *todoListMock {
				m := new(todoListMock)
				m.On("Handle", mock.Anything, todo.ListInput{Status: &pendingStatus}).Return([]todo.TodoOutput{
					{
						ID:        "123",
						Title:     "pending todo",
						Status:    "pending",
						CreatedAt: exampleDate,
						UpdatedAt: exampleDate,
					},
				}, nil).Once()
				return m
			}(),
			queryParams:    "?status=pending",
			responseBody:   `[{"id":"123","title":"pending todo","description":"","status":"pending","created_at":"2024-01-01T00:00:00Z","updated_at":"2024-01-01T00:00:00Z"}]`,
			responseStatus: http.StatusOK,
			err:            nil,
		},
		{
			name: "should list todos filtered by completed status",
			list: func() *todoListMock {
				m := new(todoListMock)
				m.On("Handle", mock.Anything, todo.ListInput{Status: &completedStatus}).Return([]todo.TodoOutput{
					{
						ID:        "456",
						Title:     "completed todo",
						Status:    "completed",
						CreatedAt: exampleDate,
						UpdatedAt: exampleDate,
					},
				}, nil).Once()
				return m
			}(),
			queryParams:    "?status=completed",
			responseBody:   `[{"id":"456","title":"completed todo","description":"","status":"completed","created_at":"2024-01-01T00:00:00Z","updated_at":"2024-01-01T00:00:00Z"}]`,
			responseStatus: http.StatusOK,
			err:            nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/todos"+tc.queryParams, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := handler.NewTodoList(tc.list)
			err := h.Handle(c)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.responseBody, strings.Trim(rec.Body.String(), "\n"))
			assert.Equal(t, tc.responseStatus, rec.Result().StatusCode)
		})
	}
}

func TestTodoList_Path(t *testing.T) {
	h := handler.NewTodoList(new(todoListMock))
	assert.Equal(t, "/todos", h.Path())
}

func TestTodoList_Method(t *testing.T) {
	h := handler.NewTodoList(new(todoListMock))
	assert.Equal(t, http.MethodGet, h.Method())
}

type todoListMock struct {
	mock.Mock
}

func (m *todoListMock) Handle(ctx context.Context, input todo.ListInput) ([]todo.TodoOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).([]todo.TodoOutput), args.Error(1)
}
