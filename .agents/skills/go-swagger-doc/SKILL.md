---
name: go-swagger-doc
description: Add Swagger/OpenAPI annotations
---

## Annotations
```go
// @Summary Short description
// @Description Detailed description
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} Output
// @Failure 400 {object} ErrorResponse
// @Router /todos/{id} [get]
```

## Commands
```bash
make swagger
# or: swag init -g cmd/api/main.go -o docs/
```

## Rules
1. Add above Handle method
2. Use correct HTTP method in @Router
3. Define response types
