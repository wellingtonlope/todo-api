package asserts

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"time"
)

const (
	StatusCreated    = 201
	StatusOK         = 200
	StatusNoContent  = 204
	StatusBadRequest = 400
	StatusNotFound   = 404
	ContentTypeJSON  = "application/json"
)

type TodoResponse struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// Response Validators

func ShouldHaveStatus(response *httptest.ResponseRecorder, expectedStatus int) error {
	if response.Code != expectedStatus {
		return fmt.Errorf("expected status %d, got %d, body: %s", expectedStatus, response.Code, response.Body.String())
	}
	return nil
}

func ShouldHaveJSONContentType(response *httptest.ResponseRecorder) error {
	if response.Header().Get("Content-Type") != ContentTypeJSON {
		return fmt.Errorf("expected Content-Type %s, got %s", ContentTypeJSON, response.Header().Get("Content-Type"))
	}
	return nil
}

func ShouldHaveValidResponseHeaders(response *httptest.ResponseRecorder, expectedStatus int) error {
	if err := ShouldHaveStatus(response, expectedStatus); err != nil {
		return err
	}

	if expectedStatus == StatusOK {
		if err := ShouldHaveJSONContentType(response); err != nil {
			return err
		}
	}

	return nil
}

// Todo Asserts

func ShouldHaveCreatedTodo(response *httptest.ResponseRecorder, expectedInput map[string]interface{}) error {
	if err := ShouldHaveValidResponseHeaders(response, StatusCreated); err != nil {
		return err
	}

	var resp TodoResponse
	if err := json.Unmarshal(response.Body.Bytes(), &resp); err != nil {
		return fmt.Errorf("failed to parse todo response: %w", err)
	}

	return validateTodoFields(&resp, expectedInput)
}

func ShouldHaveRetrievedTodo(response *httptest.ResponseRecorder, expectedID, expectedTitle, expectedDesc, expectedDueDate string) error {
	if err := ShouldHaveValidResponseHeaders(response, StatusOK); err != nil {
		return err
	}

	var resp TodoResponse
	if err := json.Unmarshal(response.Body.Bytes(), &resp); err != nil {
		return fmt.Errorf("failed to parse todo response: %w", err)
	}

	if resp.ID != expectedID {
		return fmt.Errorf("expected ID %s, got %s", expectedID, resp.ID)
	}

	return validateTodoFieldsWithStrings(&resp, expectedTitle, expectedDesc, expectedDueDate)
}

func ShouldHaveUpdatedTodo(response *httptest.ResponseRecorder, expectedID, expectedTitle, expectedDesc, expectedDueDate string) error {
	if err := ShouldHaveValidResponseHeaders(response, StatusOK); err != nil {
		return err
	}

	var resp TodoResponse
	if err := json.Unmarshal(response.Body.Bytes(), &resp); err != nil {
		return fmt.Errorf("failed to parse todo response: %w", err)
	}

	if resp.ID != expectedID {
		return fmt.Errorf("expected ID %s, got %s", expectedID, resp.ID)
	}

	return validateTodoFieldsWithStrings(&resp, expectedTitle, expectedDesc, expectedDueDate)
}

func ShouldHaveDeletedTodo(response *httptest.ResponseRecorder) error {
	return ShouldHaveValidResponseHeaders(response, StatusNoContent)
}

func ShouldHaveTodoList(response *httptest.ResponseRecorder, expectedCount int) error {
	if err := ShouldHaveValidResponseHeaders(response, StatusOK); err != nil {
		return err
	}

	var todos []TodoResponse
	if err := json.Unmarshal(response.Body.Bytes(), &todos); err != nil {
		return fmt.Errorf("failed to parse todo list response: %w", err)
	}

	if len(todos) != expectedCount {
		return fmt.Errorf("expected %d todos, got %d", expectedCount, len(todos))
	}

	return nil
}

// Error Asserts

func ShouldReturnValidationError(response *httptest.ResponseRecorder, expectedMessageContains string) error {
	return ShouldReturnError(response, StatusBadRequest, expectedMessageContains)
}

func ShouldReturnNotFoundError(response *httptest.ResponseRecorder, expectedMessageContains string) error {
	return ShouldReturnError(response, StatusNotFound, expectedMessageContains)
}

func ShouldReturnError(response *httptest.ResponseRecorder, expectedStatus int, expectedMessageContains string) error {
	if err := ShouldHaveStatus(response, expectedStatus); err != nil {
		return err
	}

	var resp ErrorResponse
	if err := json.Unmarshal(response.Body.Bytes(), &resp); err != nil {
		return fmt.Errorf("failed to parse error response: %w", err)
	}

	if !strings.Contains(resp.Message, expectedMessageContains) {
		return fmt.Errorf("expected message to contain '%s', got '%s'", expectedMessageContains, resp.Message)
	}

	return nil
}

// Field Validators

func validateTodoFields(resp *TodoResponse, expectedInput map[string]interface{}) error {
	if resp.ID == "" {
		return fmt.Errorf("todo ID should not be empty")
	}

	if resp.Title != expectedInput["title"] {
		return fmt.Errorf("expected title '%s', got '%s'", expectedInput["title"], resp.Title)
	}

	expectedDesc, hasDesc := expectedInput["description"]
	if hasDesc {
		if resp.Description != expectedDesc {
			return fmt.Errorf("expected description '%s', got '%s'", expectedDesc, resp.Description)
		}
	} else {
		if resp.Description != "" {
			return fmt.Errorf("expected description to be empty, got '%s'", resp.Description)
		}
	}

	if resp.Status != "pending" {
		return fmt.Errorf("expected status 'pending', got '%s'", resp.Status)
	}

	if resp.CreatedAt.IsZero() {
		return fmt.Errorf("CreatedAt should not be zero")
	}

	if resp.UpdatedAt.IsZero() {
		return fmt.Errorf("UpdatedAt should not be zero")
	}

	_, hasDueDate := expectedInput["due_date"]
	if hasDueDate {
		if resp.DueDate == nil {
			return fmt.Errorf("DueDate should not be nil when provided in input")
		}
	} else {
		if resp.DueDate != nil {
			return fmt.Errorf("DueDate should be nil when not provided in input")
		}
	}

	return nil
}

func validateTodoFieldsWithStrings(resp *TodoResponse, expectedTitle, expectedDesc, expectedDueDate string) error {
	if resp.Title != expectedTitle {
		return fmt.Errorf("expected title '%s', got '%s'", expectedTitle, resp.Title)
	}

	expectedDescPresent := expectedDesc != ""
	if expectedDescPresent {
		if resp.Description != expectedDesc {
			return fmt.Errorf("expected description '%s', got '%s'", expectedDesc, resp.Description)
		}
	} else {
		if resp.Description != "" {
			return fmt.Errorf("expected description to be empty, got '%s'", resp.Description)
		}
	}

	if resp.Status != "pending" {
		return fmt.Errorf("expected status 'pending', got '%s'", resp.Status)
	}

	if resp.CreatedAt.IsZero() {
		return fmt.Errorf("CreatedAt should not be zero")
	}

	if resp.UpdatedAt.IsZero() {
		return fmt.Errorf("UpdatedAt should not be zero")
	}

	expectedDueDatePresent := expectedDueDate != ""
	if expectedDueDatePresent {
		if resp.DueDate == nil {
			return fmt.Errorf("DueDate should not be nil")
		}
		parsedDueDate, err := time.Parse(time.RFC3339, expectedDueDate)
		if err != nil {
			return fmt.Errorf("invalid due_date format in test: %s", expectedDueDate)
		}
		if !resp.DueDate.Equal(parsedDueDate) {
			return fmt.Errorf("expected due_date %v, got %v", parsedDueDate, resp.DueDate)
		}
	} else {
		if resp.DueDate != nil {
			return fmt.Errorf("DueDate should be nil")
		}
	}

	return nil
}

// Partial Validators

func ShouldHaveValidTodoID(response *httptest.ResponseRecorder) error {
	var resp TodoResponse
	if err := json.Unmarshal(response.Body.Bytes(), &resp); err != nil {
		return fmt.Errorf("failed to parse todo response: %w", err)
	}

	if resp.ID == "" {
		return fmt.Errorf("todo ID should not be empty")
	}

	return nil
}

func ShouldHaveTodoTitle(response *httptest.ResponseRecorder, expectedTitle string) error {
	var resp TodoResponse
	if err := json.Unmarshal(response.Body.Bytes(), &resp); err != nil {
		return fmt.Errorf("failed to parse todo response: %w", err)
	}

	if resp.Title != expectedTitle {
		return fmt.Errorf("expected title '%s', got '%s'", expectedTitle, resp.Title)
	}

	return nil
}

func ShouldHaveTodoStatus(response *httptest.ResponseRecorder, expectedStatus string) error {
	var resp TodoResponse
	if err := json.Unmarshal(response.Body.Bytes(), &resp); err != nil {
		return fmt.Errorf("failed to parse todo response: %w", err)
	}

	if resp.Status != expectedStatus {
		return fmt.Errorf("expected status '%s', got '%s'", expectedStatus, resp.Status)
	}

	return nil
}
