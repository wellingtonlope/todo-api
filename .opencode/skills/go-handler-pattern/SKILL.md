---
name: go-handler-pattern
description: Create HTTP handlers with Echo and Swagger
---

## Pattern
```go
type XxxHandler struct { xxx todo.Xxx }

func NewXxxHandler(xxx todo.Xxx) *XxxHandler { ... }

// @Summary ...
// @Description ...
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} XxxOutput
// @Router /todos/{id} [get]
func (h *XxxHandler) Handle(c echo.Context) error { ... }
func (h *XxxHandler) Path() string { return "/todos" }
func (h *XxxHandler) Method() string { return http.MethodGet }
```

## Rules
1. Input struct: snake_case JSON tags
2. Delegate to usecase, handle HTTP only
3. Return usecase.Error or JSON response
4. Add Swagger annotations
