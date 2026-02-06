# SWAGGER.md

This document contains guidelines for generating and maintaining Swagger API documentation for the Todo API project.

## Swagger Annotations

Use the following annotations above handler methods to document API endpoints:

### Basic Annotations
- `// @Summary` - Brief description of the endpoint
- `// @Description` - Detailed explanation of what the endpoint does
- `// @Tags` - Group related endpoints together
- `// @Param` - Document request parameters
- `// @Success` - Document successful responses
- `// @Failure` - Document error responses
- `// @Router` - Define the route path and HTTP method

### Example Usage

```go
// @Summary Create a new todo
// @Description Creates a new todo item with the provided title and description
// @Tags todos
// @Param todo body todo.CreateInput true "Todo data"
// @Success 201 {object} todo.Output
// @Failure 400 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /todos [post]
func (h *Create) Handle(c echo.Context) error {
    // handler implementation
}
```

## Global API Configuration

Define global API info in `cmd/api/main.go` with these annotations:

```go
// @title Todo API
// @version 1.0
// @description This is a simple todo API built with Go and Echo
// @host localhost:8080
// @BasePath /api/v1
```

## Error Response Structure

Add `ErrorResponse` struct in `internal/infra/handler/handler.go` for consistent error responses:

```go
type ErrorResponse struct {
    Message string `json:"message"`
    Type    string `json:"type"`
}
```

## Generating Documentation

### Generate Swagger Docs

```bash
make swagger
```

### Access Documentation

When the server is running, access the Swagger UI at:
- `/swagger/index.html` - Interactive API documentation
- `/swagger/doc.json` - Raw OpenAPI specification

### Regenerating After Changes

Always regenerate documentation after making changes to handlers or API endpoints:

1. Update handler annotations
2. Run `make swagger`
3. Restart the server to see changes

## Best Practices

- Keep summaries concise and descriptive
- Use clear, detailed descriptions
- Group related endpoints with the same tag
- Document all possible response codes
- Include request/response body examples
- Use appropriate HTTP methods (GET, POST, PUT, DELETE)
