package steps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"github.com/cucumber/godog"
)

type TodoDeleteByIDContext struct {
	BaseTestContext
	CreatedTodoID string
}

func (tc *TodoDeleteByIDContext) IHaveCreatedATodoForDeletion(title, desc, dueDate string) error {
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

func (tc *TodoDeleteByIDContext) IDeleteTheTodoWithIDFromTheCreatedTodo() error {
	return tc.IDeleteTheTodoWithID(tc.CreatedTodoID)
}

func (tc *TodoDeleteByIDContext) IDeleteTheTodoWithID(id string) error {
	req := httptest.NewRequest("DELETE", "/todos/"+id, nil)
	rec := httptest.NewRecorder()
	tc.EchoApp.ServeHTTP(rec, req)
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
