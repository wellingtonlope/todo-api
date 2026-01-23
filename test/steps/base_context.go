package steps

import (
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BaseTestContext struct {
	TodoInput map[string]interface{}
	Response  *httptest.ResponseRecorder
	EchoApp   *echo.Echo
	DB        *gorm.DB
}

func (btc *BaseTestContext) SetTodoInput(title, desc, dueDate string) {
	btc.TodoInput = map[string]interface{}{
		"title": title,
	}
	if desc != "" {
		btc.TodoInput["description"] = desc
	}
	if dueDate != "" {
		btc.TodoInput["due_date"] = dueDate
	}
}

func (btc *BaseTestContext) ResetDatabase() error {
	btc.DB.Exec("DELETE FROM todos")
	return nil
}
