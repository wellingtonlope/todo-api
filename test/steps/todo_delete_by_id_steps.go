package steps

import (
	"fmt"

	"github.com/cucumber/godog"
)

type TodoDeleteByIDContext struct {
	BaseTestContext
	CreatedTodoID string
}

func (tc *TodoDeleteByIDContext) IHaveCreatedATodoForDeletion(title, desc, dueDate string) error {
	id, err := tc.CreateTodoForTest(title, desc, dueDate)
	if err != nil {
		return fmt.Errorf("failed to create todo for test: %v", err)
	}
	tc.CreatedTodoID = id
	return nil
}

func (tc *TodoDeleteByIDContext) IDeleteTheTodoWithIDFromTheCreatedTodo() error {
	return tc.IDeleteTheTodoWithID(tc.CreatedTodoID)
}

func (tc *TodoDeleteByIDContext) IDeleteTheTodoWithID(id string) error {
	client := tc.UseHTTPClient()
	rec, err := client.DeleteTodo(id)
	if err != nil {
		return err
	}
	tc.Response = rec
	return nil
}

func (tc *TodoDeleteByIDContext) TheTodoShouldBeDeletedSuccessfully() error {
	if tc.Response.Code != 204 {
		return fmt.Errorf("expected status 204, got %d, body: %s", tc.Response.Code, tc.Response.Body.String())
	}
	return nil
}

func (tc *TodoDeleteByIDContext) TheDeletionShouldFailWithNotFoundError() error {
	return validateErrorResponse(tc.Response, StatusNotFound, "not found")
}

func (tc *TodoDeleteByIDContext) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the database is reset$`, tc.ResetDatabase)
	ctx.Step(`^I have created a todo with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.IHaveCreatedATodoForDeletion)
	ctx.Step(`^I delete the todo with ID from the created todo$`, tc.IDeleteTheTodoWithIDFromTheCreatedTodo)
	ctx.Step(`^I delete the todo with ID "([^"]*)"$`, tc.IDeleteTheTodoWithID)
	ctx.Step(`^the todo should be deleted successfully$`, tc.TheTodoShouldBeDeletedSuccessfully)
	ctx.Step(`^the deletion should fail with not found error$`, tc.TheDeletionShouldFailWithNotFoundError)
}
