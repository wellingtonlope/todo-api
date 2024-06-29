package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/infra/handler"
)

func TestError(t *testing.T) {
	testCases := []struct {
		name           string
		next           echo.HandlerFunc
		responseBody   string
		responseStatus int
	}{
		{
			name:           "should do nothing when err is nil",
			next:           func(c echo.Context) error { return nil },
			responseBody:   "",
			responseStatus: http.StatusOK,
		},
		{
			name:           "should fail with internal error when cast fails",
			next:           func(c echo.Context) error { return assert.AnError },
			responseBody:   `{"message":"internal server error"}`,
			responseStatus: http.StatusInternalServerError,
		},
		{
			name:           "should fail with internal error when error type is invalid",
			next:           func(c echo.Context) error { return usecase.NewError("test", assert.AnError, "invalid") },
			responseBody:   `{"message":"internal server error"}`,
			responseStatus: http.StatusInternalServerError,
		},
		{
			name: "should handle error",
			next: func(c echo.Context) error {
				return usecase.NewError("test message", assert.AnError, usecase.ErrorTypeBadRequest)
			},
			responseBody:   `{"message":"test message"}`,
			responseStatus: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := handler.Error(tc.next)(c)
			assert.Nil(t, err)
			assert.Equal(t, tc.responseBody, strings.Trim(rec.Body.String(), "\n"))
			assert.Equal(t, tc.responseStatus, rec.Result().StatusCode)
		})
	}
}
