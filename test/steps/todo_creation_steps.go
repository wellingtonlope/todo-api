package steps

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/cucumber/godog"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TodoCreationContext struct {
	TodoInput map[string]interface{}
	Response  *httptest.ResponseRecorder
	EchoApp   *echo.Echo
	DB        *gorm.DB
}

func (tc *TodoCreationContext) SetTodoInput(title, desc, dueDate string) {
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

func (tc *TodoCreationContext) IHaveATodoInputWithTitle(title string) error {
	tc.SetTodoInput(title, "", "")
	return nil
}

func (tc *TodoCreationContext) IHaveATodoInputWithTitleAndDescription(title, desc string) error {
	tc.SetTodoInput(title, desc, "")
	return nil
}

func (tc *TodoCreationContext) IHaveATodoInputWithTitleDescriptionAndDueDate(title, desc, dueDate string) error {
	tc.SetTodoInput(title, desc, dueDate)
	return nil
}

func (tc *TodoCreationContext) ICreateTheTodo() error {
	body, _ := json.Marshal(tc.TodoInput)
	req := httptest.NewRequest("POST", "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	tc.EchoApp.ServeHTTP(rec, req)
	tc.Response = rec
	return nil
}

func validateResponseHeaders(tc *TodoCreationContext, expectedStatus int) error {
	if tc.Response.Code != expectedStatus {
		return fmt.Errorf("expected status %d, got %d, body: %s", expectedStatus, tc.Response.Code, tc.Response.Body.String())
	}
	if tc.Response.Header().Get("Content-Type") != "application/json" {
		return fmt.Errorf("expected Content-Type application/json, got %s", tc.Response.Header().Get("Content-Type"))
	}
	return nil
}

func (tc *TodoCreationContext) TheTodoShouldBeCreatedSuccessfully() error {
	if err := validateResponseHeaders(tc, 201); err != nil {
		return err
	}

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
	if resp.ID == "" {
		return fmt.Errorf("ID should not be empty")
	}
	if resp.Title != tc.TodoInput["title"] {
		return fmt.Errorf("expected title %s, got %s", tc.TodoInput["title"], resp.Title)
	}
	expectedDesc, hasDesc := tc.TodoInput["description"]
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
	_, hasDueDate := tc.TodoInput["due_date"]
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

func (tc *TodoCreationContext) TheCreationShouldFailWithValidationError() error {
	if err := validateResponseHeaders(tc, 400); err != nil {
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
	if !strings.Contains(resp.Message, "invalid input") {
		return fmt.Errorf("expected message to contain 'invalid input', got %s", resp.Message)
	}

	return nil
}

func (tc *TodoCreationContext) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, _ *godog.Scenario) (context.Context, error) {
		// Reset DB
		tc.DB.Exec("DELETE FROM todos")
		return ctx, nil
	})

	ctx.Step(`^I have a todo input with title "([^"]*)"$`, tc.IHaveATodoInputWithTitle)
	ctx.Step(`^I have a todo input with title "([^"]*)" and description "([^"]*)"$`, tc.IHaveATodoInputWithTitleAndDescription)
	ctx.Step(`^I have a todo input with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.IHaveATodoInputWithTitleDescriptionAndDueDate)
	ctx.Step(`^I create the todo$`, tc.ICreateTheTodo)
	ctx.Step(`^the todo should be created successfully$`, tc.TheTodoShouldBeCreatedSuccessfully)
	ctx.Step(`^the creation should fail with validation error$`, tc.TheCreationShouldFailWithValidationError)
}
