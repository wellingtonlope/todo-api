package helpers

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"time"
)

var (
	StatusCreated    = 201
	StatusOK         = 200
	StatusNoContent  = 204
	StatusBadRequest = 400
	StatusNotFound   = 404
	ContentTypeJSON  = "application/json"
)

type TodoCreateInput struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

type TodoUpdateInput struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

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

func ParseTodoResponse(response *httptest.ResponseRecorder) (TodoResponse, error) {
	var resp TodoResponse
	if err := json.Unmarshal(response.Body.Bytes(), &resp); err != nil {
		return resp, fmt.Errorf("failed to parse todo response: %w", err)
	}
	return resp, nil
}

func ParseTodoListResponse(response *httptest.ResponseRecorder) ([]TodoResponse, error) {
	var todos []TodoResponse
	if err := json.Unmarshal(response.Body.Bytes(), &todos); err != nil {
		return nil, fmt.Errorf("failed to parse todo list response: %w", err)
	}
	return todos, nil
}

func ParseErrorResponse(response *httptest.ResponseRecorder) (ErrorResponse, error) {
	var resp ErrorResponse
	if err := json.Unmarshal(response.Body.Bytes(), &resp); err != nil {
		return resp, fmt.Errorf("failed to parse error response: %w", err)
	}
	return resp, nil
}

func ValidateStatus(response *httptest.ResponseRecorder, expectedStatus int) error {
	if response.Code != expectedStatus {
		return fmt.Errorf("expected status %d, got %d, body: %s", expectedStatus, response.Code, response.Body.String())
	}
	return nil
}

func ValidateJSONContentType(response *httptest.ResponseRecorder) error {
	if response.Header().Get("Content-Type") != ContentTypeJSON {
		return fmt.Errorf("expected Content-Type %s, got %s", ContentTypeJSON, response.Header().Get("Content-Type"))
	}
	return nil
}

func ValidateMessageContains(response *httptest.ResponseRecorder, expectedMessageContains string) error {
	resp, err := ParseErrorResponse(response)
	if err != nil {
		return err
	}

	if expectedMessageContains != "" && !strings.Contains(resp.Message, expectedMessageContains) {
		return fmt.Errorf("expected message to contain '%s', got '%s'", expectedMessageContains, resp.Message)
	}

	return nil
}
