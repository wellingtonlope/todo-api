package steps

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

func validateResponseHeaders(response *httptest.ResponseRecorder, expectedStatus int) error {
	if response.Code != expectedStatus {
		return fmt.Errorf("expected status %d, got %d, body: %s", expectedStatus, response.Code, response.Body.String())
	}
	return nil
}

func validateTodoResponse(response *httptest.ResponseRecorder, input map[string]interface{}) error {
	var resp TodoResponse
	if err := json.Unmarshal(response.Body.Bytes(), &resp); err != nil {
		return fmt.Errorf("failed to parse todo response: %w", err)
	}

	if resp.ID == "" {
		return fmt.Errorf("todo ID should not be empty")
	}

	if resp.Title != input["title"] {
		return fmt.Errorf("expected title '%v', got '%s'", input["title"], resp.Title)
	}

	if input["description"] != nil {
		if resp.Description != input["description"] {
			return fmt.Errorf("expected description '%v', got '%s'", input["description"], resp.Description)
		}
	}

	if resp.Status != "pending" {
		return fmt.Errorf("expected status 'pending', got '%s'", resp.Status)
	}

	return nil
}

func validateRetrievedTodoResponse(response *httptest.ResponseRecorder, expectedID, title, desc, dueDate string) error {
	var resp TodoResponse
	if err := json.Unmarshal(response.Body.Bytes(), &resp); err != nil {
		return fmt.Errorf("failed to parse todo response: %w", err)
	}

	if resp.ID != expectedID {
		return fmt.Errorf("expected ID %s, got %s", expectedID, resp.ID)
	}

	if resp.Title != title {
		return fmt.Errorf("expected title '%s', got '%s'", title, resp.Title)
	}

	if desc != "" {
		if resp.Description != desc {
			return fmt.Errorf("expected description '%s', got '%s'", desc, resp.Description)
		}
	} else {
		if resp.Description != "" {
			return fmt.Errorf("expected description to be empty, got '%s'", resp.Description)
		}
	}

	return nil
}

func validateErrorResponse(response *httptest.ResponseRecorder, expectedStatus int, expectedMessageContains string) error {
	if err := validateResponseHeaders(response, expectedStatus); err != nil {
		return err
	}

	var errorResp map[string]interface{}
	if err := json.Unmarshal(response.Body.Bytes(), &errorResp); err != nil {
		return fmt.Errorf("failed to parse error response: %w", err)
	}

	message, ok := errorResp["message"].(string)
	if !ok {
		return fmt.Errorf("error response should contain a message field")
	}

	if expectedMessageContains != "" && !strings.Contains(message, expectedMessageContains) {
		return fmt.Errorf("expected message to contain '%s', got '%s'", expectedMessageContains, message)
	}

	return nil
}
