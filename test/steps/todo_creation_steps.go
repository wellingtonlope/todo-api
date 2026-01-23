package steps

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"

	"github.com/cucumber/godog"
)

type TodoCreationContext struct {
	BaseTestContext
}

func (tc *TodoCreationContext) IHaveATodoInput(title, desc, dueDate string) error {
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

func (tc *TodoCreationContext) TheTodoShouldBeCreatedSuccessfully() error {
	if err := validateResponseHeaders(tc.Response, StatusCreated); err != nil {
		return err
	}

	return validateTodoResponse(tc.Response, tc.TodoInput)
}

func (tc *TodoCreationContext) TheCreationShouldFailWithValidationError() error {
	return validateErrorResponse(tc.Response, StatusBadRequest, "invalid input")
}

func (tc *TodoCreationContext) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the database is reset$`, tc.ResetDatabase)
	ctx.Step(`^I have a todo input with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.IHaveATodoInput)
	ctx.Step(`^I create the todo$`, tc.ICreateTheTodo)
	ctx.Step(`^the todo should be created successfully$`, tc.TheTodoShouldBeCreatedSuccessfully)
	ctx.Step(`^the creation should fail with validation error$`, tc.TheCreationShouldFailWithValidationError)
}
