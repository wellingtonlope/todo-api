package steps

import (
	"encoding/json"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BaseTestContext struct {
	TodoInput  map[string]interface{}
	Response   *httptest.ResponseRecorder
	EchoApp    *echo.Echo
	DB         *gorm.DB
	httpClient *HTTPClient
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
	if err := btc.DB.Exec("DELETE FROM todos").Error; err != nil {
		return err
	}
	return nil
}

func (btc *BaseTestContext) UseHTTPClient() *HTTPClient {
	if btc.httpClient == nil {
		btc.httpClient = NewHTTPClient(btc.EchoApp)
	}
	return btc.httpClient
}

func (btc *BaseTestContext) ResetHTTPClient() {
	btc.httpClient = nil
}

func (btc *BaseTestContext) CreateTodoForTest(title, desc, dueDate string) (string, error) {
	input := map[string]interface{}{
		"title": title,
	}
	if desc != "" {
		input["description"] = desc
	}
	if dueDate != "" {
		input["due_date"] = dueDate
	}
	return btc.CreateTodoWithInput(input)
}

func (btc *BaseTestContext) CreateTodoWithInput(input map[string]interface{}) (string, error) {
	client := btc.UseHTTPClient()
	rec, err := client.CreateTodo(input)
	if err != nil {
		return "", err
	}
	btc.Response = rec

	type response struct {
		ID string `json:"id"`
	}
	var resp response
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		return "", err
	}
	return resp.ID, nil
}
