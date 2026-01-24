package steps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"github.com/cucumber/godog"
)

type TodoMarkPendingContext struct {
	BaseTestContext
	CreatedTodoID string
}

func (tc *TodoMarkPendingContext) IHaveCreatedATodoForMarkingPending(title, desc, dueDate string) error {
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

func (tc *TodoMarkPendingContext) IMarkTheTodoWithIDFromTheCreatedTodoAsPending() error {
	return tc.IMarkTheTodoWithIDAsPending(tc.CreatedTodoID)
}

func (tc *TodoMarkPendingContext) IMarkTheTodoWithIDAsPending(id string) error {
	req := httptest.NewRequest("POST", "/todos/"+id+"/pending", nil)
	rec := httptest.NewRecorder()
	tc.EchoApp.ServeHTTP(rec, req)
	tc.Response = rec
	return nil
}

func (tc *TodoMarkPendingContext) TheTodoShouldBeMarkedAsPendingSuccessfully() error {
	if err := validateResponseHeaders(tc.Response, StatusOK); err != nil {
		return err
	}

	var resp TodoResponse
	err := json.Unmarshal(tc.Response.Body.Bytes(), &resp)
	if err != nil {
		return err
	}

	if resp.ID != tc.CreatedTodoID {
		return fmt.Errorf("expected ID %s, got %s", tc.CreatedTodoID, resp.ID)
	}

	if resp.Status != "pending" {
		return fmt.Errorf("expected status pending, got %s", resp.Status)
	}

	return nil
}

func (tc *TodoMarkPendingContext) TheMarkingAsPendingShouldFailWithNotFoundError() error {
	return validateErrorResponse(tc.Response, StatusNotFound, "not found")
}

func (tc *TodoMarkPendingContext) TheMarkingAsPendingShouldSucceedButStatusRemainsPending() error {
	if err := validateResponseHeaders(tc.Response, StatusOK); err != nil {
		return err
	}

	var resp TodoResponse
	err := json.Unmarshal(tc.Response.Body.Bytes(), &resp)
	if err != nil {
		return err
	}

	if resp.ID != tc.CreatedTodoID {
		return fmt.Errorf("expected ID %s, got %s", tc.CreatedTodoID, resp.ID)
	}

	if resp.Status != "pending" {
		return fmt.Errorf("expected status pending, got %s", resp.Status)
	}

	return nil
}

// IHaveCreatedATodoWithTitleAndDescription creates a todo with title and description for mark pending tests
func (tc *TodoMarkPendingContext) IHaveCreatedATodoWithTitleAndDescription(title, desc string) error {
	tc.SetTodoInput(title, desc, "")
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

// IMarkTheTodoAsComplete marks the created todo as complete
func (tc *TodoMarkPendingContext) IMarkTheTodoAsComplete() error {
	req := httptest.NewRequest("POST", "/todos/"+tc.CreatedTodoID+"/complete", nil)
	rec := httptest.NewRecorder()
	tc.EchoApp.ServeHTTP(rec, req)
	tc.Response = rec
	return nil
}

// IMarkTheTodoAsPending marks the created todo as pending
func (tc *TodoMarkPendingContext) IMarkTheTodoAsPending() error {
	return tc.IMarkTheTodoWithIDFromTheCreatedTodoAsPending()
}

// TheMarkPendingShouldFailWithNotFoundError validates 404 error for mark pending
func (tc *TodoMarkPendingContext) TheMarkPendingShouldFailWithNotFoundError() error {
	return validateErrorResponse(tc.Response, StatusNotFound, "not found")
}

// TheMarkPendingShouldFailWithValidationError validates 400 error for mark pending
func (tc *TodoMarkPendingContext) TheMarkPendingShouldFailWithValidationError() error {
	return validateErrorResponse(tc.Response, StatusBadRequest, "")
}

func (tc *TodoMarkPendingContext) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the database is reset$`, tc.ResetDatabase)
	ctx.Step(`^I have created a todo with title "([^"]*)" and description "([^"]*)"$`, tc.IHaveCreatedATodoWithTitleAndDescription)
	ctx.Step(`^I have created a todo with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.IHaveCreatedATodoForMarkingPending)
	ctx.Step(`^I mark the todo as complete$`, tc.IMarkTheTodoAsComplete)
	ctx.Step(`^I mark the todo as pending$`, tc.IMarkTheTodoAsPending)
	ctx.Step(`^I mark the todo with ID from the created todo as pending$`, tc.IMarkTheTodoWithIDFromTheCreatedTodoAsPending)
	ctx.Step(`^I mark the todo with ID "([^"]*)" as pending$`, tc.IMarkTheTodoWithIDAsPending)
	ctx.Step(`^the todo should be marked as pending successfully$`, tc.TheTodoShouldBeMarkedAsPendingSuccessfully)
	ctx.Step(`^the mark pending should fail with not found error$`, tc.TheMarkPendingShouldFailWithNotFoundError)
	ctx.Step(`^the mark pending should fail with validation error$`, tc.TheMarkPendingShouldFailWithValidationError)
	ctx.Step(`^the marking as pending should fail with not found error$`, tc.TheMarkingAsPendingShouldFailWithNotFoundError)
	ctx.Step(`^the marking as pending should succeed but status remains pending$`, tc.TheMarkingAsPendingShouldSucceedButStatusRemainsPending)
}
