# ARCHITECTURE.md

This document contains architectural patterns, design guidelines, and project structure for the Todo API project.

## Project Overview

This is a Todo API built with Go 1.24, using Echo for HTTP, GORM with SQLite for persistence, and Uber FX for dependency injection. The project follows Clean Architecture with clear separation between domain, application (usecase), and infrastructure layers.

## Design Patterns

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
    gorm/             # GORM database implementations
pkg/
  clock/              # Shared packages (clock utilities)
```

## Layer Responsibilities

### Domain Layer
- Business entities and rules
- Domain errors and validation
- Pure business logic without external dependencies

### Application Layer (Usecases)
- Application business logic
- Use case orchestration
- Input/output validation
- Error handling and mapping

### Infrastructure Layer
- HTTP handlers and routing
- Database implementations
- External service integrations
- Framework-specific code