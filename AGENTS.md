# AGENTS.md

Go Todo API - Clean Architecture with Echo, GORM, Uber FX.

## Commands

```bash
make test    # Run tests
make build   # Build binary
make server  # Run server
make format  # Format code
make lint    # Lint code
make swagger # Generate Swagger
```

## Constraints

1. Write all code and docs in English
2. Run `make lint && make test` before finishing
3. Keep functions small and focused

## Available Skills

Use `$skill-name` or let the agent auto-trigger:

| Skill | Trigger | Purpose |
|-------|---------|---------|
| `go-clean-architecture` | "clean arch", "layer" | Domain/usecase/infra structure |
| `go-usecase-pattern` | "usecase", "use case" | Create new use cases |
| `go-handler-pattern` | "handler", "endpoint" | Create HTTP handlers |
| `go-unit-test` | "test", "unit test" | Write table-driven tests |
| `go-bdd-test` | "bdd", "cucumber", "godog" | BDD tests with godog |
| `go-code-quality` | "lint", "format", "quality" | Format & lint rules |
| `go-swagger-doc` | "swagger", "openapi" | API documentation |
| `conventional-commit` | "commit" | Commit message format |

## Documentation

- [DEVELOPMENT.md](docs/DEVELOPMENT.md) - Build, testing, tooling
- [STYLE_GUIDE.md](docs/STYLE_GUIDE.md) - Naming, code style, errors
- [ARCHITECTURE.md](docs/ARCHITECTURE.md) - Patterns and layer responsibilities

## Structure

```
cmd/api/     # Entrypoint
internal/
  domain/    # Entities, errors
  app/usecase/   # Business logic
  infra/     # Handlers, gorm, memory
pkg/         # Shared utilities
```
