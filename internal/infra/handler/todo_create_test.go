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

func TestTodoCreate_Handle(t *testing.T) {
	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	testCases := []struct {
		name           string
		create         *todoCreateMock
		requestBody    string
		responseBody   string
		responseStatus int
		err            error
	}{
		{
			name:           "should fail when JSON invalid",
			create:         new(todoCreateMock),
			requestBody:    "{",
			responseBody:   "",
			responseStatus: http.StatusOK,
			err: usecase.NewError("invalid JSON input", func() error {
				e := echo.New()
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{"))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				var aux any
				return c.Bind(&aux)
			}(), usecase.ErrorTypeBadRequest),
		},
		{
			name: "should fail when create use case fails",
			create: func() *todoCreateMock {
				m := new(todoCreateMock)
				m.On("Handle", mock.Anything, todo.CreateInput{
					Title:       "example title",
					Description: "example description",
				}).Return(todo.TodoOutput{}, usecase.AnError).Once()
				return m
			}(),
			requestBody:    `{"title":"example title","description":"example description"}`,
			responseBody:   "",
			responseStatus: http.StatusOK,
			err:            usecase.AnError,
		},
		{
			name: "should create a todo",
			create: func() *todoCreateMock {
				m := new(todoCreateMock)
				m.On("Handle", mock.Anything, todo.CreateInput{
					Title:       "example title",
					Description: "example description",
				}).Return(todo.TodoOutput{
					ID:          "123",
					Title:       "example title",
					Description: "example description",
					Status:      "pending",
					CreatedAt:   exampleDate,
					UpdatedAt:   exampleDate,
				}, nil).Once()
				return m
			}(),
			requestBody:    `{"title":"example title","description":"example description"}`,
			responseBody:   `{"id":"123","title":"example title","description":"example description","status":"pending","created_at":"2024-01-01T00:00:00Z","updated_at":"2024-01-01T00:00:00Z"}`,
			responseStatus: http.StatusCreated,
			err:            nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/todos")
			h := handler.NewTodoCreate(tc.create)
			err := h.Handle(c)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.responseBody, strings.Trim(rec.Body.String(), "\n"))
			assert.Equal(t, tc.responseStatus, rec.Result().StatusCode)
		})
	}
}

func TestTodoCreate_Path(t *testing.T) {
	h := handler.NewTodoCreate(new(todoCreateMock))
	assert.Equal(t, "/todos", h.Path())
}

func TestTodoCreate_Method(t *testing.T) {
	h := handler.NewTodoCreate(new(todoCreateMock))
	assert.Equal(t, http.MethodPost, h.Method())
}

type todoCreateMock struct {
	mock.Mock
}

func (m *todoCreateMock) Handle(ctx context.Context, input todo.CreateInput) (todo.TodoOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(todo.TodoOutput), args.Error(1)
}
