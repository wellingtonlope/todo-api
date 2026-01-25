package steps

import (
	"encoding/json"
	"fmt"

	"github.com/cucumber/godog"
)

type TodoCompleteContext struct {
	BaseTestContext
	createdTodoID string
}

func (tc *TodoCompleteContext) IHaveATodoWithTitleAndDescription(title, desc string) error {
	tc.SetTodoInput(title, desc, "")
	return nil
}

func (tc *TodoCompleteContext) ICreateTheTodo() error {
	client := tc.UseHTTPClient()
	rec, err := client.CreateTodo(tc.TodoInput)
	if err != nil {
		return err
	}
	tc.Response = rec

	// Extract the created todo ID
	if rec.Code == 201 {
		var resp TodoResponse
		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		if err == nil {
			tc.createdTodoID = resp.ID
		}
	}
	return nil
}

func (tc *TodoCompleteContext) TheTodoShouldBeCreatedSuccessfully() error {
	if err := validateResponseHeaders(tc.Response, StatusCreated); err != nil {
		return err
	}
	return validateTodoResponse(tc.Response, tc.TodoInput)
}

func (tc *TodoCompleteContext) IMarkTheTodoAsComplete() error {
	client := tc.UseHTTPClient()
	rec, err := client.CompleteTodo(tc.createdTodoID)
	if err != nil {
		return err
	}
	tc.Response = rec
	return nil
}

func (tc *TodoCompleteContext) IMarkTheTodoAsCompleteAgain() error {
	return tc.IMarkTheTodoAsComplete()
}

func (tc *TodoCompleteContext) IMarkTodoWithIDAsComplete(id string) error {
	client := tc.UseHTTPClient()
	rec, err := client.CompleteTodo(id)
	if err != nil {
		return err
	}
	tc.Response = rec
	return nil
}

func (tc *TodoCompleteContext) TheTodoShouldBeMarkedAsCompletedSuccessfully() error {
	if err := validateResponseHeaders(tc.Response, StatusOK); err != nil {
		return err
	}

	var resp TodoResponse
	err := json.Unmarshal(tc.Response.Body.Bytes(), &resp)
	if err != nil {
		return err
	}

	if resp.Status != "completed" {
		return fmt.Errorf("expected todo status to be 'completed', got '%s'", resp.Status)
	}

	return nil
}

func (tc *TodoCompleteContext) TheCompletionShouldFailWithNotFoundError() error {
	return validateErrorResponse(tc.Response, StatusNotFound, "todo not found")
}

func (tc *TodoCompleteContext) TheCompletionShouldFailWithValidationError() error {
	return validateErrorResponse(tc.Response, StatusBadRequest, "")
}

func (tc *TodoCompleteContext) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the database is reset$`, tc.ResetDatabase)
	ctx.Step(`^I have a todo with title "([^"]*)" and description "([^"]*)"$`, tc.IHaveATodoWithTitleAndDescription)
	ctx.Step(`^I create the todo$`, tc.ICreateTheTodo)
	ctx.Step(`^the todo should be created successfully$`, tc.TheTodoShouldBeCreatedSuccessfully)
	ctx.Step(`^I mark the todo as complete$`, tc.IMarkTheTodoAsComplete)
	ctx.Step(`^I mark the todo as complete again$`, tc.IMarkTheTodoAsCompleteAgain)
	ctx.Step(`^I mark todo with ID "([^"]*)" as complete$`, tc.IMarkTodoWithIDAsComplete)
	ctx.Step(`^the todo should be marked as completed successfully$`, tc.TheTodoShouldBeMarkedAsCompletedSuccessfully)
	ctx.Step(`^the completion should fail with not found error$`, tc.TheCompletionShouldFailWithNotFoundError)
}
