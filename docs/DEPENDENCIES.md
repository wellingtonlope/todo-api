# DEPENDENCIES.md

This document contains information about the key dependencies and technology stack used in the Todo API project.

## Technology Stack

### Core Framework
- **Go 1.25** - Programming language and runtime
- **Echo v4** - HTTP web framework for building REST APIs

### Database & Persistence
- **GORM** - ORM (Object-Relational Mapping) for database operations
- **SQLite** - Database engine (configurable, can be swapped for PostgreSQL, MySQL, etc.)

### Dependency Injection
- **Uber FX** - Dependency injection framework for Go applications

### Testing
- **Testify** - Testing utilities including assertions and mocking
  - `github.com/stretchr/testify/assert` - Assertion functions
  - `github.com/stretchr/testify/mock` - Mocking framework

### Utilities
- **Google UUID** - UUID generation and parsing
- **Clock utilities** - Time abstraction for testing (located in `pkg/clock/`)

### API Documentation
- **Swagger/OpenAPI** - API documentation specification
- **Echo-Swagger** - Swagger UI integration for Echo
- **Swag** - Swagger documentation generator

## Key Dependencies by Category

### HTTP Layer
```
github.com/labstack/echo/v4          # HTTP framework
github.com/swaggo/echo-swagger       # Swagger UI for Echo
```

### Business Logic
```
go.uber.org/fx                       # Dependency injection
```

### Data Layer
```
gorm.io/gorm                         # ORM
gorm.io/driver/sqlite                # SQLite driver
```

### Testing
```
github.com/stretchr/testify          # Testing utilities
github.com/stretchr/testify/mock     # Mocking framework
```

### Utilities
```
github.com/google/uuid               # UUID generation
github.com/swaggo/swag               # Swagger generator
```

## Dependency Management

### Adding New Dependencies
1. Add the dependency to `go.mod` using `go get`
2. Update this document with the new dependency and its purpose
3. Follow the existing import organization patterns

### Version Management
- Keep dependencies updated to stable versions
- Check for security updates regularly
- Test thoroughly after dependency updates

### External Service Integration
When adding external service integrations:
- Create interfaces in the usecase layer
- Implement concrete adapters in the infrastructure layer
- Use dependency injection to provide implementations
- Include mock implementations for testing
