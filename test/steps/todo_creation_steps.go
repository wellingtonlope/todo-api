package steps

import (
	"github.com/cucumber/godog"

	"github.com/wellingtonlope/todo-api/test/helpers"
)

type TodoCreationContext struct {
	BaseTestContext
}

func (tc *TodoCreationContext) IHaveATodoInput(title, desc, dueDate string) error {
	tc.SetTodoInput(title, desc, dueDate)
	return nil
}

func (tc *TodoCreationContext) ICreateTheTodo() error {
	client := tc.UseHTTPClient()
	rec, err := client.CreateTodo(tc.TodoInput)
	if err != nil {
		return err
	}
	tc.Response = rec
	return nil
}

func (tc *TodoCreationContext) TheTodoShouldBeCreatedSuccessfully() error {
	if err := validateResponseHeaders(tc.Response, helpers.StatusCreated); err != nil {
		return err
	}

	return validateTodoResponse(tc.Response, tc.TodoInput)
}

func (tc *TodoCreationContext) TheCreationShouldFailWithValidationError() error {
	return validateErrorResponse(tc.Response, helpers.StatusBadRequest, "invalid input")
}

func (tc *TodoCreationContext) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the database is reset$`, tc.ResetDatabase)
	ctx.Step(`^I have a todo input with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.IHaveATodoInput)
	ctx.Step(`^I create the todo$`, tc.ICreateTheTodo)
	ctx.Step(`^the todo should be created successfully$`, tc.TheTodoShouldBeCreatedSuccessfully)
	ctx.Step(`^the creation should fail with validation error$`, tc.TheCreationShouldFailWithValidationError)
}
