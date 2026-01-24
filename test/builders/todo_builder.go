package builders

import (
	"time"

	"github.com/wellingtonlope/todo-api/internal/domain"
)

type TodoBuilder struct {
	title       string
	description string
	dueDate     *time.Time
	createdAt   time.Time
}

func NewTodoBuilder() *TodoBuilder {
	now := time.Now()
	return &TodoBuilder{
		title:       "Default Todo",
		description: "",
		dueDate:     nil,
		createdAt:   now,
	}
}

func (tb *TodoBuilder) WithTitle(title string) *TodoBuilder {
	tb.title = title
	return tb
}

func (tb *TodoBuilder) WithDescription(description string) *TodoBuilder {
	tb.description = description
	return tb
}

func (tb *TodoBuilder) WithDueDate(dueDate time.Time) *TodoBuilder {
	tb.dueDate = &dueDate
	return tb
}

func (tb *TodoBuilder) WithDueDateString(dueDateStr string) (*TodoBuilder, error) {
	if dueDateStr == "" {
		tb.dueDate = nil
		return tb, nil
	}

	parsedDueDate, err := time.Parse(time.RFC3339, dueDateStr)
	if err != nil {
		return nil, err
	}

	tb.dueDate = &parsedDueDate
	return tb, nil
}

func (tb *TodoBuilder) WithCreatedAt(createdAt time.Time) *TodoBuilder {
	tb.createdAt = createdAt
	return tb
}

func (tb *TodoBuilder) Build() (domain.Todo, error) {
	return domain.NewTodo(tb.title, tb.description, tb.createdAt, tb.dueDate)
}

func (tb *TodoBuilder) BuildDomain() domain.Todo {
	todo, _ := tb.Build() // We assume valid data in tests
	return todo
}

func (tb *TodoBuilder) BuildRequestMap() map[string]interface{} {
	request := map[string]interface{}{
		"title": tb.title,
	}

	if tb.description != "" {
		request["description"] = tb.description
	}

	if tb.dueDate != nil {
		request["due_date"] = tb.dueDate.Format(time.RFC3339)
	}

	return request
}

// Convenience methods for common scenarios

func (tb *TodoBuilder) Minimal() *TodoBuilder {
	return tb.WithTitle("Minimal Todo")
}

func (tb *TodoBuilder) Complete() *TodoBuilder {
	tomorrow := time.Now().Add(24 * time.Hour)
	return tb.
		WithTitle("Complete Todo").
		WithDescription("This is a complete todo with description and due date").
		WithDueDate(tomorrow)
}

func (tb *TodoBuilder) WithFutureDueDate(daysFromNow int) *TodoBuilder {
	futureDate := time.Now().AddDate(0, 0, daysFromNow)
	return tb.WithDueDate(futureDate)
}

func (tb *TodoBuilder) WithPastDueDate(daysAgo int) *TodoBuilder {
	pastDate := time.Now().AddDate(0, 0, -daysAgo)
	return tb.WithDueDate(pastDate)
}

// Static convenience methods

func NewTodo() *TodoBuilder {
	return NewTodoBuilder()
}

func MinimalTodo() domain.Todo {
	return NewTodoBuilder().Minimal().BuildDomain()
}

func CompleteTodo() domain.Todo {
	return NewTodoBuilder().Complete().BuildDomain()
}

func TodoWithTitle(title string) domain.Todo {
	return NewTodoBuilder().WithTitle(title).BuildDomain()
}

func TodoWithTitleAndDescription(title, description string) domain.Todo {
	return NewTodoBuilder().
		WithTitle(title).
		WithDescription(description).
		BuildDomain()
}
