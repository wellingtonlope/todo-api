# AGENTS.md

This document provides guidelines for agentic coding assistants working on this codebase.

## Project Overview

This is a Todo API built with Go 1.24, using Echo for HTTP, GORM with SQLite for persistence, and Uber FX for dependency injection. The project follows Clean Architecture with clear separation between domain, application (usecase), and infrastructure layers.

## Build, Lint, and Test Commands

```bash
# Run all tests
make test
go test ./... -count=1

# Run a single test
go test ./internal/app/usecase/todo -run TestCreate_Handle -v

# Build the binary
make build
go build ./cmd/api/

# Run the server
make server
go run ./cmd/api/

# Format code
make format
gofumpt -w .

# Run linter
make lint
golangci-lint run

# Generate Swagger docs
swag init -g cmd/api/main.go -o docs/
```

### Testing Patterns

- Use table-driven tests with `t.Run()` for multiple scenarios
- Mock external dependencies using `testify/mock`
- Create mock types in the same test file or in a dedicated `*_mock.go` file
- Assert both expected results and error conditions
- Use `mock.Mock.AssertExpectations(t)` to verify all expectations were met

## Code Style Guidelines

### General Principles

- Follow Clean Architecture: domain → application (usecase) → infrastructure
- Keep functions small and focused on a single responsibility
- Use interfaces to define contracts between layers
- Prefer composition over inheritance

### Naming Conventions

- **Packages**: Use singular, lowercase names (e.g., `domain`, `handler`, `usecase`)
- **Variables**: Use camelCase for local variables, PascalCase for exported identifiers
- **Interfaces**: Name interfaces after the action they enable (e.g., `Create`, `GetAll`, `Clock`)
- **Types**: Public types use PascalCase, internal types are unexported, suffix output types with `Output`, input types with `Input`
- **Test files**: Name `*_test.go` in the same directory as the code being tested
- **Test functions**: Name `TestPackageName_MethodName` (e.g., `TestCreate_Handle`)

### Imports

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

### Error Handling

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

### Handler Design

- Handlers should only handle HTTP concerns (binding, response formatting)
- Delegate business logic to use cases
- Handler types are unexported with exported constructors
- Use snake_case for JSON field tags (e.g., `json:"created_at"`)
- Use PascalCase for struct field names

### Usecase Design

- Use case inputs/outputs should be simple structs with camelCase fields
- Private implementation types with exported constructor functions
- Accept interfaces for dependencies to enable testing with mocks

### Domain Layer

- Keep domain models pure with no external dependencies
- Validate input in factory functions
- Use sentinel errors for domain validation failures
- Return both the model and an error from constructors

### Dependency Injection

Use Uber FX for dependency injection with `fx.Provide()` and `fx.Invoke()`.

### Code Formatting

Run `make format` (gofumpt) before committing. Gofumpt is stricter than gofmt: no unused imports, no redundant blank lines.

### Linting

Run `make lint` (golangci-lint) before committing.

### Testing Guidelines

- Use `github.com/stretchr/testify/assert` and `github.com/stretchr/testify/mock`
- Create mock types that embed `mock.Mock`
- Use table-driven tests for multiple scenarios
- Name test cases clearly (e.g., "should fail when input is invalid")
- Always call `AssertExpectations()` on mocks

## Swagger Documentation Guidelines

- Use `// @Summary`, `// @Description`, `// @Tags`, `// @Param`, `// @Success`, `// @Failure`, `// @Router` annotations above handler methods
- Define global API info in `cmd/api/main.go` with `@title`, `@version`, `@description`, `@host`, `@BasePath`
- Add `ErrorResponse` struct in `internal/infra/handler/handler.go` for error responses
- Access documentation at `/swagger/index.html` when server is running
- Regenerate docs after changes: `swag init -g cmd/api/main.go -o docs/`

## File Structure

```
cmd/api/              # Application entrypoint
docs/                 # Generated Swagger documentation
internal/
  domain/             # Business entities and domain errors
  app/
    usecase/          # Application business logic (use cases)
      todo/           # Todo-related use cases
  infra/
    handler/          # HTTP handlers
    memory/           # In-memory implementations
pkg/
  clock/              # Shared packages (clock utilities)
```

## Key Dependencies

- `github.com/labstack/echo/v4` - HTTP framework
- `go.uber.org/fx` - Dependency injection
- `gorm.io/gorm` - ORM
- `github.com/stretchr/testify` - Testing utilities
- `github.com/google/uuid` - UUID generation
- `github.com/swaggo/echo-swagger` - Swagger UI for Echo
- `github.com/swaggo/swag` - Swagger documentation generator
