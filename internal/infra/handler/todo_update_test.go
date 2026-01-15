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

func TestTodoUpdate_Handle(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	testCases := []struct {
		name           string
		update         *todoUpdateMock
		pathID         string
		requestBody    string
		responseBody   string
		responseStatus int
		err            error
	}{
		{
			name:           "should fail when JSON invalid",
			update:         new(todoUpdateMock),
			pathID:         "123",
			requestBody:    "{",
			responseBody:   "",
			responseStatus: http.StatusOK,
			err: usecase.NewError("invalid JSON input", func() error {
				e := echo.New()
				req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader("{"))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				var aux any
				return c.Bind(&aux)
			}(), usecase.ErrorTypeBadRequest),
		},
		{
			name: "should fail when update use case fails",
			update: func() *todoUpdateMock {
				m := new(todoUpdateMock)
				m.On("Handle", mock.Anything, todo.UpdateInput{
					ID:          "123",
					Title:       "example title",
					Description: "example description",
				}).Return(todo.TodoOutput{}, usecase.AnError).Once()
				return m
			}(),
			pathID:         "123",
			requestBody:    `{"title":"example title","description":"example description"}`,
			responseBody:   "",
			responseStatus: http.StatusOK,
			err:            usecase.AnError,
		},
		{
			name: "should update a todo",
			update: func() *todoUpdateMock {
				m := new(todoUpdateMock)
				m.On("Handle", mock.Anything, todo.UpdateInput{
					ID:          "123",
					Title:       "example title",
					Description: "example description",
				}).Return(todo.TodoOutput{
					ID:          "123",
					Title:       "example title",
					Description: "example description",
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDate,
				}, nil).Once()
				return m
			}(),
			pathID:         "123",
			requestBody:    `{"title":"example title","description":"example description"}`,
			responseBody:   `{"id":"123","title":"example title","description":"example description","created_at":"2024-01-01T00:00:00Z","updated_at":"2024-01-01T00:00:00Z"}`,
			responseStatus: http.StatusOK,
			err:            nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(tc.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/todos/:id")
			c.SetParamNames("id")
			c.SetParamValues(tc.pathID)
			h := handler.NewTodoUpdate(tc.update)
			err := h.Handle(c)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.responseBody, strings.Trim(rec.Body.String(), "\n"))
			assert.Equal(t, tc.responseStatus, rec.Result().StatusCode)
		})
	}
}

func TestTodoUpdate_Path(t *testing.T) {
	h := handler.NewTodoUpdate(new(todoUpdateMock))
	assert.Equal(t, "/todos/:id", h.Path())
}

func TestTodoUpdate_Method(t *testing.T) {
	h := handler.NewTodoUpdate(new(todoUpdateMock))
	assert.Equal(t, http.MethodPut, h.Method())
}

type todoUpdateMock struct {
	mock.Mock
}

func (m *todoUpdateMock) Handle(ctx context.Context, input todo.UpdateInput) (todo.TodoOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(todo.TodoOutput), args.Error(1)
}
