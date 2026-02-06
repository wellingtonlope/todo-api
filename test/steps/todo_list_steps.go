package steps

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cucumber/godog"
)

type TodoListContext struct {
	BaseTestContext
	CreatedTodoIDs []string
	ResetStoreFunc func()
}

func (tc *TodoListContext) IRequestAllTodos() error {
	client := tc.UseHTTPClient()
	rec, err := client.ListTodos()
	if err != nil {
		return err
	}
	tc.Response = rec
	return nil
}

func (tc *TodoListContext) TheResponseShouldBeSuccessfulWithStatus(status int) error {
	if err := validateResponseHeaders(tc.Response, status); err != nil {
		return err
	}
	return nil
}

func (tc *TodoListContext) TheResponseShouldContainAnEmptyListOfTodos() error {
	var todos []TodoResponse
	err := json.Unmarshal(tc.Response.Body.Bytes(), &todos)
	if err != nil {
		return err
	}

	if len(todos) != 0 {
		return fmt.Errorf("expected 0 todos, got %d", len(todos))
	}

	return nil
}

func (tc *TodoListContext) TheResponseShouldContainAListWithTodos(count int) error {
	var todos []TodoResponse
	err := json.Unmarshal(tc.Response.Body.Bytes(), &todos)
	if err != nil {
		return err
	}

	// For debugging, print the actual count
	if len(todos) != count {
		return fmt.Errorf("expected %d todos, got %d", count, len(todos))
	}

	return nil
}

func (tc *TodoListContext) IHaveCreatedATodoWith(title, desc, dueDate string) error {
	id, err := tc.CreateTodoForTest(title, desc, dueDate)
	if err != nil {
		return fmt.Errorf("failed to create todo for test: %v", err)
	}
	tc.CreatedTodoIDs = append(tc.CreatedTodoIDs, id)
	return nil
}

func (tc *TodoListContext) IHaveCreatedACompletedTodoWith(title, desc, dueDate string) error {
	id, err := tc.CreateTodoForTest(title, desc, dueDate)
	if err != nil {
		return fmt.Errorf("failed to create todo for test: %v", err)
	}
	tc.CreatedTodoIDs = append(tc.CreatedTodoIDs, id)

	client := tc.UseHTTPClient()
	_, err = client.CompleteTodo(id)
	if err != nil {
		return fmt.Errorf("failed to complete todo: %v", err)
	}
	return nil
}

func (tc *TodoListContext) IRequestTodosWithStatus(status string) error {
	client := tc.UseHTTPClient()
	rec, err := client.ListTodosWithStatus(status)
	if err != nil {
		return err
	}
	tc.Response = rec
	return nil
}

func (tc *TodoListContext) TheFirstTodoShouldHaveTitleDescDueDate(title, desc, dueDate string) error {
	return tc.validateTodoAtIndex(0, title, desc, dueDate)
}

func (tc *TodoListContext) TheSecondTodoShouldHaveTitleDescDueDate(title, desc, dueDate string) error {
	return tc.validateTodoAtIndex(1, title, desc, dueDate)
}

func (tc *TodoListContext) validateTodoAtIndex(index int, expectedTitle, expectedDesc, expectedDueDate string) error {
	var todos []TodoResponse
	err := json.Unmarshal(tc.Response.Body.Bytes(), &todos)
	if err != nil {
		return err
	}

	if index >= len(todos) {
		return fmt.Errorf("todo index %d out of range, only %d todos available", index, len(todos))
	}

	todo := todos[index]

	if todo.Title != expectedTitle {
		return fmt.Errorf("expected title %s, got %s", expectedTitle, todo.Title)
	}

	expectedDescPresent := expectedDesc != ""
	if expectedDescPresent {
		if todo.Description != expectedDesc {
			return fmt.Errorf("expected description %s, got %s", expectedDesc, todo.Description)
		}
	} else {
		if todo.Description != "" {
			return fmt.Errorf("expected description to be empty, got %s", todo.Description)
		}
	}

	expectedDueDatePresent := expectedDueDate != ""
	if expectedDueDatePresent {
		if todo.DueDate == nil {
			return fmt.Errorf("DueDate should not be nil")
		}
		// Check if due date matches (allow for time format differences)
		expectedTime, err := time.Parse(time.RFC3339, expectedDueDate)
		if err != nil {
			return fmt.Errorf("invalid expected due date format: %s", expectedDueDate)
		}
		if !todo.DueDate.Equal(expectedTime) {
			return fmt.Errorf("expected due date %s, got %s", expectedDueDate, todo.DueDate.Format(time.RFC3339))
		}
	} else {
		if todo.DueDate != nil {
			return fmt.Errorf("DueDate should be nil")
		}
	}

	return nil
}

func (tc *TodoListContext) ResetDatabaseAndContext() error {
	tc.CreatedTodoIDs = []string{}
	// Reset both GORM database and in-memory store
	if tc.ResetStoreFunc != nil {
		tc.ResetStoreFunc()
	}
	if err := tc.ResetDatabase(); err != nil {
		return err
	}
	// Force HTTP client recreation to use new app
	tc.ResetHTTPClient()
	return nil
}

func (tc *TodoListContext) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the database is reset$`, tc.ResetDatabaseAndContext)
	ctx.Step(`^I request all todos$`, tc.IRequestAllTodos)
	ctx.Step(`^I request todos with status "([^"]*)"$`, tc.IRequestTodosWithStatus)
	ctx.Step(`^the response should be successful with status (\d+)$`, tc.TheResponseShouldBeSuccessfulWithStatus)
	ctx.Step(`^the response should contain an empty list of todos$`, tc.TheResponseShouldContainAnEmptyListOfTodos)
	ctx.Step(`^the response should contain a list with (\d+) todos$`, tc.TheResponseShouldContainAListWithTodos)
	ctx.Step(`^the response should contain a list with (\d+) todo$`, tc.TheResponseShouldContainAListWithTodos)
	ctx.Step(`^I have created a todo with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.IHaveCreatedATodoWith)
	ctx.Step(`^I have created a completed todo with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.IHaveCreatedACompletedTodoWith)
	ctx.Step(`^the first todo should have title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.TheFirstTodoShouldHaveTitleDescDueDate)
	ctx.Step(`^the second todo should have title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.TheSecondTodoShouldHaveTitleDescDueDate)
}
