---
name: go-swagger-doc
description: Add or update Swagger/OpenAPI documentation for handlers
---

## What I do
Add Swagger annotations to HTTP handlers.

## When to use me
Creating or modifying API endpoints.

## Annotations
```go
// @Summary Short description
// @Description Detailed description
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param todo body createInput true "Todo data"
// @Success 200 {object} todoOutput
// @Success 201 {object} todoOutput
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /todos [get|post|put|delete]
```

## Rules
1. Add above the Handle method
2. Use correct HTTP method in @Router
3. Define response types (use existing todoOutput or create new)
4. Run `swag init -g cmd/api/main.go -o docs/` after changes