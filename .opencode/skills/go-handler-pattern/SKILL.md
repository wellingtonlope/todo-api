---
name: go-handler-pattern
description: Create HTTP handlers following Echo framework pattern
---

## What I do
Create HTTP handlers with Echo framework and Swagger documentation.

## When to use me
Adding new API endpoints.

## Pattern
```go
type (
    xxxInput struct {
        Field string `json:"field"`
    }
    XxxHandler struct { xxx todo.Xxx }
)

func NewXxxHandler(xxx todo.Xxx) *XxxHandler { ... }

// @Summary ...
// @Description ...
// @Tags todos
// @Accept json
// @Produce json
// @Param ...
// @Success ...
// @Failure ...
// @Router /todos [method]
func (h *XxxHandler) Handle(c echo.Context) error { ... }
func (h *XxxHandler) Path() string { return "/todos" }
func (h *XxxHandler) Method() string { return http.MethodXxx }
```

## Rules
1. Input struct uses snake_case JSON tags
2. Delegate to usecase, handle only HTTP
3. Return usecase.Error or JSON response
4. Add Swagger annotations