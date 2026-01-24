package steps

import (
	"net/http/httptest"

	"github.com/wellingtonlope/todo-api/test/asserts"
)

// Re-export constants for backward compatibility
const (
	StatusCreated    = asserts.StatusCreated
	StatusOK         = asserts.StatusOK
	StatusNoContent  = asserts.StatusNoContent
	StatusBadRequest = asserts.StatusBadRequest
	StatusNotFound   = asserts.StatusNotFound
	ContentTypeJSON  = asserts.ContentTypeJSON
)

// Re-export types for backward compatibility
type TodoResponse = asserts.TodoResponse

// Backward compatibility functions that delegate to the new asserts library
// These maintain the existing API while internally using the improved asserts

func validateResponseHeaders(response *httptest.ResponseRecorder, expectedStatus int) error {
	return asserts.ShouldHaveValidResponseHeaders(response, expectedStatus)
}

func validateTodoResponse(response *httptest.ResponseRecorder, expectedInput map[string]interface{}) error {
	return asserts.ShouldHaveCreatedTodo(response, expectedInput)
}

func validateRetrievedTodoResponse(response *httptest.ResponseRecorder, expectedID, expectedTitle, expectedDesc, expectedDueDate string) error {
	return asserts.ShouldHaveRetrievedTodo(response, expectedID, expectedTitle, expectedDesc, expectedDueDate)
}

func validateErrorResponse(response *httptest.ResponseRecorder, expectedStatus int, expectedMessageContains string) error {
	return asserts.ShouldReturnError(response, expectedStatus, expectedMessageContains)
}
