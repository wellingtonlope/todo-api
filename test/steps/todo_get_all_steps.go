package steps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"github.com/cucumber/godog"
)

type TodoGetAllContext struct {
	BaseTestContext
	CreatedTodoIDs []string
	ResetStoreFunc func()
}

func (tc *TodoGetAllContext) IRequestAllTodos() error {
	req := httptest.NewRequest("GET", "/todos", nil)
	rec := httptest.NewRecorder()
	tc.EchoApp.ServeHTTP(rec, req)
	tc.Response = rec
	return nil
}

func (tc *TodoGetAllContext) TheResponseShouldBeSuccessfulWithStatus(status int) error {
	if err := validateResponseHeaders(tc.Response, status); err != nil {
		return err
	}
	return nil
}

func (tc *TodoGetAllContext) TheResponseShouldContainAnEmptyListOfTodos() error {
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

func (tc *TodoGetAllContext) TheResponseShouldContainAListWithTodos(count int) error {
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

func (tc *TodoGetAllContext) IHaveCreatedATodoWith(title, desc, dueDate string) error {
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
	tc.CreatedTodoIDs = append(tc.CreatedTodoIDs, resp.ID)
	return nil
}

func (tc *TodoGetAllContext) TheFirstTodoShouldHaveTitleDescDueDate(title, desc, dueDate string) error {
	return tc.validateTodoAtIndex(0, title, desc, dueDate)
}

func (tc *TodoGetAllContext) TheSecondTodoShouldHaveTitleDescDueDate(title, desc, dueDate string) error {
	return tc.validateTodoAtIndex(1, title, desc, dueDate)
}

func (tc *TodoGetAllContext) validateTodoAtIndex(index int, expectedTitle, expectedDesc, expectedDueDate string) error {
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
	} else {
		if todo.DueDate != nil {
			return fmt.Errorf("DueDate should be nil")
		}
	}

	return nil
}

func (tc *TodoGetAllContext) ResetDatabaseAndContext() error {
	tc.CreatedTodoIDs = []string{}
	// Reset both GORM database and in-memory store
	if err := tc.ResetDatabase(); err != nil {
		return err
	}
	if tc.ResetStoreFunc != nil {
		tc.ResetStoreFunc()
	}
	return nil
}

func (tc *TodoGetAllContext) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the database is reset$`, tc.ResetDatabaseAndContext)
	ctx.Step(`^I request all todos$`, tc.IRequestAllTodos)
	ctx.Step(`^the response should be successful with status (\d+)$`, tc.TheResponseShouldBeSuccessfulWithStatus)
	ctx.Step(`^the response should contain an empty list of todos$`, tc.TheResponseShouldContainAnEmptyListOfTodos)
	ctx.Step(`^the response should contain a list with (\d+) todos$`, tc.TheResponseShouldContainAListWithTodos)
	ctx.Step(`^the response should contain a list with (\d+) todo$`, tc.TheResponseShouldContainAListWithTodos)
	ctx.Step(`^I have created a todo with title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.IHaveCreatedATodoWith)
	ctx.Step(`^the first todo should have title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.TheFirstTodoShouldHaveTitleDescDueDate)
	ctx.Step(`^the second todo should have title "([^"]*)", description "([^"]*)" and due_date "([^"]*)"$`, tc.TheSecondTodoShouldHaveTitleDescDueDate)
}
