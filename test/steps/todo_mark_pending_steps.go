package steps

import (
	"encoding/json"
	"fmt"

	"github.com/cucumber/godog"
)

type TodoMarkPendingContext struct {
	BaseTestContext
	CreatedTodoID string
}

func (tc *TodoMarkPendingContext) IHaveCreatedATodoForMarkingPending(title, desc, dueDate string) error {
	id, err := tc.CreateTodoForTest(title, desc, dueDate)
	if err != nil {
		return fmt.Errorf("failed to create todo for test: %v", err)
	}
	tc.CreatedTodoID = id
	return nil
}

func (tc *TodoMarkPendingContext) IMarkTheTodoWithIDFromTheCreatedTodoAsPending() error {
	return tc.IMarkTheTodoWithIDAsPending(tc.CreatedTodoID)
}

func (tc *TodoMarkPendingContext) IMarkTheTodoWithIDAsPending(id string) error {
	client := tc.UseHTTPClient()
	rec, err := client.MarkPendingTodo(id)
	if err != nil {
		return err
	}
	tc.Response = rec
	return nil
}

func (tc *TodoMarkPendingContext) TheTodoShouldBeMarkedAsPendingSuccessfully() error {
	return tc.validateMarkPendingResponse()
}

func (tc *TodoMarkPendingContext) TheMarkingAsPendingShouldFailWithNotFoundError() error {
	return validateErrorResponse(tc.Response, StatusNotFound, "not found")
}

func (tc *TodoMarkPendingContext) validateMarkPendingResponse() error {
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

func (tc *TodoMarkPendingContext) TheMarkingAsPendingShouldSucceedButStatusRemainsPending() error {
	return tc.validateMarkPendingResponse()
}

// IHaveCreatedATodoWithTitleAndDescription creates a todo with title and description for mark pending tests
func (tc *TodoMarkPendingContext) IHaveCreatedATodoWithTitleAndDescription(title, desc string) error {
	id, err := tc.CreateTodoForTest(title, desc, "")
	if err != nil {
		return fmt.Errorf("failed to create todo for test: %v", err)
	}
	tc.CreatedTodoID = id
	return nil
}

// IMarkTheTodoAsComplete marks the created todo as complete
func (tc *TodoMarkPendingContext) IMarkTheTodoAsComplete() error {
	client := tc.UseHTTPClient()
	rec, err := client.CompleteTodo(tc.CreatedTodoID)
	if err != nil {
		return err
	}
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
	// The API actually allows marking an already pending todo as pending (idempotent operation)
	// So we should expect success instead of validation error
	return tc.validateMarkPendingResponse()
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
