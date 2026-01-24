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

func validateResponseHeaders(response *httptest.ResponseRecorder, expectedStatus int) error {
	if response.Code != expectedStatus {
		return fmt.Errorf("expected status %d, got %d, body: %s", expectedStatus, response.Code, response.Body.String())
	}
	if expectedStatus == StatusOK && response.Header().Get("Content-Type") != ContentTypeJSON {
		return fmt.Errorf("expected Content-Type %s, got %s", ContentTypeJSON, response.Header().Get("Content-Type"))
	}
	return nil
}

func validateTodoResponse(response *httptest.ResponseRecorder, expectedInput map[string]interface{}) error {
	var resp TodoResponse
	err := json.Unmarshal(response.Body.Bytes(), &resp)
	if err != nil {
		return err
	}

	// Validate fields
	if resp.ID == "" {
		return fmt.Errorf("ID should not be empty")
	}
	if resp.Title != expectedInput["title"] {
		return fmt.Errorf("expected title %s, got %s", expectedInput["title"], resp.Title)
	}
	expectedDesc, hasDesc := expectedInput["description"]
	if hasDesc {
		if resp.Description != expectedDesc {
			return fmt.Errorf("expected description %s, got %s", expectedDesc, resp.Description)
		}
	} else {
		if resp.Description != "" {
			return fmt.Errorf("expected description to be empty, got %s", resp.Description)
		}
	}
	if resp.Status != "pending" {
		return fmt.Errorf("expected status pending, got %s", resp.Status)
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
			return fmt.Errorf("DueDate should not be nil")
		}
	} else {
		if resp.DueDate != nil {
			return fmt.Errorf("DueDate should be nil")
		}
	}

	return nil
}

func validateRetrievedTodoResponse(response *httptest.ResponseRecorder, expectedID, expectedTitle, expectedDesc, expectedDueDate string) error {
	var resp TodoResponse
	err := json.Unmarshal(response.Body.Bytes(), &resp)
	if err != nil {
		return err
	}

	// Validate fields
	if resp.ID != expectedID {
		return fmt.Errorf("expected ID %s, got %s", expectedID, resp.ID)
	}
	if resp.Title != expectedTitle {
		return fmt.Errorf("expected title %s, got %s", expectedTitle, resp.Title)
	}
	expectedDescPresent := expectedDesc != ""
	if expectedDescPresent {
		if resp.Description != expectedDesc {
			return fmt.Errorf("expected description %s, got %s", expectedDesc, resp.Description)
		}
	} else {
		if resp.Description != "" {
			return fmt.Errorf("expected description to be empty, got %s", resp.Description)
		}
	}
	if resp.Status != "pending" {
		return fmt.Errorf("expected status pending, got %s", resp.Status)
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

func validateErrorResponse(response *httptest.ResponseRecorder, expectedStatus int, expectedMessageContains string) error {
	if err := validateResponseHeaders(response, expectedStatus); err != nil {
		return err
	}

	var resp struct {
		Message string `json:"message"`
	}
	err := json.Unmarshal(response.Body.Bytes(), &resp)
	if err != nil {
		return err
	}

	if !strings.Contains(resp.Message, expectedMessageContains) {
		return fmt.Errorf("expected message to contain '%s', got %s", expectedMessageContains, resp.Message)
	}

	return nil
}
