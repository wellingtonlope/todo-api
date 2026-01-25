package steps

import (
	"fmt"

	"github.com/cucumber/godog"
)

type TodoGetByIDContext struct {
	BaseTestContext
	CreatedTodoID string
}

func (tc *TodoGetByIDContext) IHaveCreatedATodo(title, desc, dueDate string) error {
	id, err := tc.CreateTodoForTest(title, desc, dueDate)
	if err != nil {
		return fmt.Errorf("failed to create todo for test: %v", err)
	}
	tc.CreatedTodoID = id
	return nil
}

func (tc *TodoGetByIDContext) IRequestTheTodoWithIDFromTheCreatedTodo() error {
	return tc.IRequestTheTodoWithID(tc.CreatedTodoID)
}

func (tc *TodoGetByIDContext) IRequestTheTodoWithID(id string) error {
	client := tc.UseHTTPClient()
	rec, err := client.GetTodo(id)
	if err != nil {
		return err
	}
	tc.Response = rec
	return nil
}

func (tc *TodoGetByIDContext) TheTodoShouldBeRetrievedSuccessfullyWithTitleDescDueDate(title, desc, dueDate string) error {
	if err := validateResponseHeaders(tc.Response, StatusOK); err != nil {
		return err
	}

	return validateRetrievedTodoResponse(tc.Response, tc.CreatedTodoID, title, desc, dueDate)
}

func (tc *TodoGetByIDContext) TheRetrievalShouldFailWithNotFoundError() error {
	return validateErrorResponse(tc.Response, StatusNotFound, "not found")
}

func (tc *TodoGetByIDContext) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the database is reset$`, tc.ResetDatabase)
	ctx.Step(`^I have created a todo with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.IHaveCreatedATodo)
	ctx.Step(`^I request the todo with ID from the created todo$`, tc.IRequestTheTodoWithIDFromTheCreatedTodo)
	ctx.Step(`^I request the todo with ID "([^"]*)"$`, tc.IRequestTheTodoWithID)
	ctx.Step(`^the todo should be retrieved successfully with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.TheTodoShouldBeRetrievedSuccessfullyWithTitleDescDueDate)
	ctx.Step(`^the retrieval should fail with not found error$`, tc.TheRetrievalShouldFailWithNotFoundError)
}
