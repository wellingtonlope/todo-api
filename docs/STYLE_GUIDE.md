# STYLE_GUIDE.md

This document contains code style guidelines, naming conventions, and formatting standards for the Todo API project.

## General Principles

See [ARCHITECTURE.md](ARCHITECTURE.md) for architectural principles and [AGENTS.md](../AGENTS.md) for overall project guidelines.

## Naming Conventions

### Packages
- Use singular, lowercase names (e.g., `domain`, `handler`, `usecase`)

### Variables
- Use camelCase for local variables
- Use PascalCase for exported identifiers

### Interfaces
- Name interfaces after the action they enable (e.g., `Create`, `List`, `Clock`)

### Types
- Public types use PascalCase
- Internal types are unexported
- Suffix output types with `Output`
- Suffix input types with `Input`

### Test Files
- Name `*_test.go` in the same directory as the code being tested

### Test Functions
- Name `TestPackageName_MethodName` (e.g., `TestCreate_Handle`)

## Import Organization

Organize imports in three groups separated by blank lines: stdlib, third-party, then internal packages.

```go
import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
	"github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
)
```

## Error Handling

Use the custom `usecase.Error` type for consistent error handling:

```go
type ErrorType string
type Error struct {
	Message string
	Cause   error
	Type    ErrorType
}

var (
	ErrorTypeBadRequest    = ErrorType("bad_request")
	ErrorTypeInternalError = ErrorType("internal_error")
	ErrorTypeNotFound      = ErrorType("not_found")
)

usecase.NewError("message", cause, ErrorTypeBadRequest)
```

Wrap errors using `%w` with sentinel errors from the domain layer (e.g., `ErrTodoInvalidInput`).

## JSON Conventions

- Use snake_case for JSON field tags (e.g., `json:"created_at"`)
- Use PascalCase for struct field names

## Language Requirements

All code and documentation must be written in English.
