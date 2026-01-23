package steps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/cucumber/godog"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TodoGetByIDContext struct {
	TodoInput     map[string]interface{}
	Response      *httptest.ResponseRecorder
	EchoApp       *echo.Echo
	DB            *gorm.DB
	CreatedTodoID string
}

func (tc *TodoGetByIDContext) SetTodoInput(title, desc, dueDate string) {
	tc.TodoInput = map[string]interface{}{
		"title": title,
	}
	if desc != "" {
		tc.TodoInput["description"] = desc
	}
	if dueDate != "" {
		tc.TodoInput["due_date"] = dueDate
	}
}

func (tc *TodoGetByIDContext) ResetDatabase() error {
	tc.DB.Exec("DELETE FROM todos")
	return nil
}

func (tc *TodoGetByIDContext) IHaveCreatedATodo(title, desc, dueDate string) error {
	tc.SetTodoInput(title, desc, dueDate)
	body, _ := json.Marshal(tc.TodoInput)
	req := httptest.NewRequest("POST", "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	tc.EchoApp.ServeHTTP(rec, req)

	if rec.Code != 201 {
		return fmt.Errorf("failed to create todo for test, status: %d, body: %s", rec.Code, rec.Body.String())
	}

	// Parse the created todo to get ID
	var resp struct {
		ID string `json:"id"`
	}
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		return err
	}
	tc.CreatedTodoID = resp.ID
	return nil
}

func (tc *TodoGetByIDContext) IRequestTheTodoWithIDFromTheCreatedTodo() error {
	return tc.IRequestTheTodoWithID(tc.CreatedTodoID)
}

func (tc *TodoGetByIDContext) IRequestTheTodoWithID(id string) error {
	req := httptest.NewRequest("GET", "/todos/"+id, nil)
	rec := httptest.NewRecorder()
	tc.EchoApp.ServeHTTP(rec, req)
	tc.Response = rec
	return nil
}

func validateGetByIDResponseHeaders(tc *TodoGetByIDContext, expectedStatus int) error {
	if tc.Response.Code != expectedStatus {
		return fmt.Errorf("expected status %d, got %d, body: %s", expectedStatus, tc.Response.Code, tc.Response.Body.String())
	}
	if expectedStatus == 200 && tc.Response.Header().Get("Content-Type") != "application/json" {
		return fmt.Errorf("expected Content-Type application/json, got %s", tc.Response.Header().Get("Content-Type"))
	}
	return nil
}

func validateRetrievedTodoResponse(tc *TodoGetByIDContext, expectedTitle, expectedDesc, expectedDueDate string) error {
	// Parse JSON
	var resp struct {
		ID          string     `json:"id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Status      string     `json:"status"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at"`
		DueDate     *time.Time `json:"due_date,omitempty"`
	}
	err := json.Unmarshal(tc.Response.Body.Bytes(), &resp)
	if err != nil {
		return err
	}

	// Validate fields
	if resp.ID != tc.CreatedTodoID {
		return fmt.Errorf("expected ID %s, got %s", tc.CreatedTodoID, resp.ID)
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

func (tc *TodoGetByIDContext) TheTodoShouldBeRetrievedSuccessfullyWithTitleDescDueDate(title, desc, dueDate string) error {
	if err := validateGetByIDResponseHeaders(tc, 200); err != nil {
		return err
	}

	return validateRetrievedTodoResponse(tc, title, desc, dueDate)
}

func (tc *TodoGetByIDContext) TheRetrievalShouldFailWithNotFoundError() error {
	if err := validateGetByIDResponseHeaders(tc, 404); err != nil {
		return err
	}

	// Parse error JSON
	var resp struct {
		Message string `json:"message"`
	}
	err := json.Unmarshal(tc.Response.Body.Bytes(), &resp)
	if err != nil {
		return err
	}

	// Validate error
	if !strings.Contains(resp.Message, "not found") {
		return fmt.Errorf("expected message to contain 'not found', got %s", resp.Message)
	}

	return nil
}

func (tc *TodoGetByIDContext) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the database is reset$`, tc.ResetDatabase)
	ctx.Step(`^I have created a todo with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.IHaveCreatedATodo)
	ctx.Step(`^I request the todo with ID from the created todo$`, tc.IRequestTheTodoWithIDFromTheCreatedTodo)
	ctx.Step(`^I request the todo with ID "([^"]*)"$`, tc.IRequestTheTodoWithID)
	ctx.Step(`^the todo should be retrieved successfully with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.TheTodoShouldBeRetrievedSuccessfullyWithTitleDescDueDate)
	ctx.Step(`^the retrieval should fail with not found error$`, tc.TheRetrievalShouldFailWithNotFoundError)
}
