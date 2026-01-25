package steps

import (
	"encoding/json"
	"fmt"

	"github.com/cucumber/godog"
)

type TodoUpdateContext struct {
	BaseTestContext
	CreatedTodoID string
	UpdateInput   map[string]interface{}
}

func (tc *TodoUpdateContext) IHaveCreatedATodoForUpdate(title, desc, dueDate string) error {
	id, err := tc.CreateTodoForTest(title, desc, dueDate)
	if err != nil {
		return fmt.Errorf("failed to create todo for test: %v", err)
	}
	tc.CreatedTodoID = id
	return nil
}

func (tc *TodoUpdateContext) IHaveATodoUpdateInput(title, desc, dueDate string) error {
	tc.UpdateInput = map[string]interface{}{
		"title": title,
	}
	if desc != "" {
		tc.UpdateInput["description"] = desc
	}
	if dueDate != "" {
		tc.UpdateInput["due_date"] = dueDate
	}
	return nil
}

func (tc *TodoUpdateContext) IUpdateTheTodoWithIDFromTheCreatedTodo() error {
	return tc.IUpdateTheTodoWithID(tc.CreatedTodoID)
}

func (tc *TodoUpdateContext) IUpdateTheTodoWithID(id string) error {
	client := tc.UseHTTPClient()
	rec, err := client.UpdateTodo(id, tc.UpdateInput)
	if err != nil {
		return err
	}
	tc.Response = rec
	return nil
}

func (tc *TodoUpdateContext) IUpdateTheTodoWithIDWithTitleDescAndDueDate(id, title, desc, dueDate string) error {
	tc.UpdateInput = map[string]interface{}{
		"title": title,
	}
	if desc != "" {
		tc.UpdateInput["description"] = desc
	}
	if dueDate != "" {
		tc.UpdateInput["due_date"] = dueDate
	}
	return tc.IUpdateTheTodoWithID(id)
}

func (tc *TodoUpdateContext) IUpdateTheTodoWithIDFromCreatedTodoWithTitleDescAndDueDate(title, desc, dueDate string) error {
	tc.UpdateInput = map[string]interface{}{
		"title": title,
	}
	if desc != "" {
		tc.UpdateInput["description"] = desc
	}
	if dueDate != "" {
		tc.UpdateInput["due_date"] = dueDate
	}
	return tc.IUpdateTheTodoWithIDFromTheCreatedTodo()
}

func (tc *TodoUpdateContext) TheTodoShouldBeUpdatedSuccessfullyWithTitleDescAndDueDate(title, desc, dueDate string) error {
	if err := validateResponseHeaders(tc.Response, StatusOK); err != nil {
		return err
	}

	// Parse the response to verify the updated todo
	var resp TodoResponse
	err := json.Unmarshal(tc.Response.Body.Bytes(), &resp)
	if err != nil {
		return err
	}

	// Validate updated fields
	if resp.Title != title {
		return fmt.Errorf("expected title %s, got %s", title, resp.Title)
	}

	if desc != "" {
		if resp.Description != desc {
			return fmt.Errorf("expected description %s, got %s", desc, resp.Description)
		}
	} else {
		// When description is empty in input, it should become empty in response
		if resp.Description != "" {
			return fmt.Errorf("expected description to be empty, got %s", resp.Description)
		}
	}

	if dueDate != "" {
		if resp.DueDate == nil {
			return fmt.Errorf("DueDate should not be nil")
		}
	} else {
		// When due_date is empty in input, it should become nil in response
		if resp.DueDate != nil {
			return fmt.Errorf("DueDate should be nil")
		}
	}

	return nil
}

func (tc *TodoUpdateContext) TheUpdateShouldFailWithNotFoundError() error {
	return validateErrorResponse(tc.Response, StatusNotFound, "not found")
}

func (tc *TodoUpdateContext) TheUpdateShouldFailWithValidationError() error {
	return validateErrorResponse(tc.Response, StatusBadRequest, "invalid input")
}

func (tc *TodoUpdateContext) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the database is reset$`, tc.ResetDatabase)
	ctx.Step(`^I have created a todo with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.IHaveCreatedATodoForUpdate)
	ctx.Step(`^I have a todo update input with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.IHaveATodoUpdateInput)
	ctx.Step(`^I update the todo with ID from the created todo$`, tc.IUpdateTheTodoWithIDFromTheCreatedTodo)
	ctx.Step(`^I update the todo with ID "([^"]*)"$`, tc.IUpdateTheTodoWithID)
	ctx.Step(`^I update the todo with ID "([^"]*)" with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.IUpdateTheTodoWithIDWithTitleDescAndDueDate)
	ctx.Step(`^I update the todo with ID from the created todo with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.IUpdateTheTodoWithIDFromCreatedTodoWithTitleDescAndDueDate)
	ctx.Step(`^the todo should be updated successfully with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.TheTodoShouldBeUpdatedSuccessfullyWithTitleDescAndDueDate)
	ctx.Step(`^the update should fail with not found error$`, tc.TheUpdateShouldFailWithNotFoundError)
	ctx.Step(`^the update should fail with validation error$`, tc.TheUpdateShouldFailWithValidationError)
}
