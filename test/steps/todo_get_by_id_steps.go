package steps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"github.com/cucumber/godog"
)

type TodoGetByIDContext struct {
	BaseTestContext
	CreatedTodoID string
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
