package steps

import (
	"net/http/httptest"
	"time"

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
	// Convert map[string]interface{} to TodoCreateInput
	input := &asserts.TodoCreateInput{
		Title:       getStringFromMap(expectedInput, "title"),
		Description: getStringFromMap(expectedInput, "description"),
	}

	if dueDate, hasDueDate := expectedInput["due_date"]; hasDueDate && dueDate != nil {
		if dueDateStr, ok := dueDate.(string); ok && dueDateStr != "" {
			if parsedDueDate, err := time.Parse(time.RFC3339, dueDateStr); err == nil {
				input.DueDate = &parsedDueDate
			}
		}
	}

	return asserts.ShouldHaveCreatedTodo(response, input)
}

func validateRetrievedTodoResponse(response *httptest.ResponseRecorder, expectedID, expectedTitle, expectedDesc, expectedDueDate string) error {
	// Convert string parameters to TodoUpdateInput
	input := &asserts.TodoUpdateInput{
		Title:       expectedTitle,
		Description: expectedDesc,
	}

	if expectedDueDate != "" {
		if parsedDueDate, err := time.Parse(time.RFC3339, expectedDueDate); err == nil {
			input.DueDate = &parsedDueDate
		}
	}

	return asserts.ShouldHaveRetrievedTodo(response, expectedID, input)
}

// Helper function to safely get string from map
func getStringFromMap(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func validateErrorResponse(response *httptest.ResponseRecorder, expectedStatus int, expectedMessageContains string) error {
	return asserts.ShouldReturnError(response, expectedStatus, expectedMessageContains)
}
