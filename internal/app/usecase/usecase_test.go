package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
)

func TestNewError(t *testing.T) {
	testCases := []struct {
		name      string
		message   string
		cause     error
		errorType usecase.ErrorType
		result    usecase.Error
	}{
		{
			name:      "should create error",
			message:   "example message",
			cause:     assert.AnError,
			errorType: usecase.ErrorTypeInternalError,
			result: usecase.Error{
				Message: "example message",
				Cause:   assert.AnError,
				Type:    usecase.ErrorTypeInternalError,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := usecase.NewError(tc.message, tc.cause, tc.errorType)
			assert.Equal(t, tc.result, result)
		})
	}
}

func TestError_Error(t *testing.T) {
	testCases := []struct {
		name   string
		error  usecase.Error
		result string
	}{
		{
			name:   "should return message",
			error:  usecase.NewError("example message", nil, usecase.ErrorTypeInternalError),
			result: "example message",
		},
		{
			name:   "should return cause",
			error:  usecase.NewError("example message", assert.AnError, usecase.ErrorTypeInternalError),
			result: assert.AnError.Error(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.error.Error()
			assert.Equal(t, tc.result, result)
		})
	}
}
