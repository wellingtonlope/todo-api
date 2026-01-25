package steps

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

type HTTPClient struct {
	app *echo.Echo
}

func NewHTTPClient(app *echo.Echo) *HTTPClient {
	return &HTTPClient{app: app}
}

func (c *HTTPClient) CreateTodo(input map[string]interface{}) (*httptest.ResponseRecorder, error) {
	body, _ := json.Marshal(input)
	req := httptest.NewRequest("POST", "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c.app.ServeHTTP(rec, req)
	return rec, nil
}

func (c *HTTPClient) GetTodo(id string) (*httptest.ResponseRecorder, error) {
	req := httptest.NewRequest("GET", "/todos/"+id, nil)
	rec := httptest.NewRecorder()
	c.app.ServeHTTP(rec, req)
	return rec, nil
}

func (c *HTTPClient) UpdateTodo(id string, input map[string]interface{}) (*httptest.ResponseRecorder, error) {
	body, _ := json.Marshal(input)
	req := httptest.NewRequest("PUT", "/todos/"+id, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c.app.ServeHTTP(rec, req)
	return rec, nil
}

func (c *HTTPClient) DeleteTodo(id string) (*httptest.ResponseRecorder, error) {
	req := httptest.NewRequest("DELETE", "/todos/"+id, nil)
	rec := httptest.NewRecorder()
	c.app.ServeHTTP(rec, req)
	return rec, nil
}

func (c *HTTPClient) CompleteTodo(id string) (*httptest.ResponseRecorder, error) {
	req := httptest.NewRequest("POST", "/todos/"+id+"/complete", nil)
	rec := httptest.NewRecorder()
	c.app.ServeHTTP(rec, req)
	return rec, nil
}

func (c *HTTPClient) MarkPendingTodo(id string) (*httptest.ResponseRecorder, error) {
	req := httptest.NewRequest("POST", "/todos/"+id+"/pending", nil)
	rec := httptest.NewRecorder()
	c.app.ServeHTTP(rec, req)
	return rec, nil
}

func (c *HTTPClient) GetAllTodos() (*httptest.ResponseRecorder, error) {
	req := httptest.NewRequest("GET", "/todos", nil)
	rec := httptest.NewRecorder()
	c.app.ServeHTTP(rec, req)
	return rec, nil
}
