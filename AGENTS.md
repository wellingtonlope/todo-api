# AGENTS.md

This document provides guidelines for agentic coding assistants working on this codebase.

## Project Overview

This is a Todo API built with Go 1.24, using Echo for HTTP, GORM with SQLite for persistence, and Uber FX for dependency injection. The project follows Clean Architecture with clear separation between domain, application (usecase), and infrastructure layers.

## Quick Reference

### Essential Commands
```bash
# Run tests
make test

# Build and run
make build && make server

# Format and lint
make format && make lint
```

### Documentation Structure
- **[DEVELOPMENT.md](docs/DEVELOPMENT.md)** - Build commands, testing patterns, code quality
- **[STYLE_GUIDE.md](docs/STYLE_GUIDE.md)** - Naming conventions, code style, error handling
- **[ARCHITECTURE.md](docs/ARCHITECTURE.md)** - Design patterns, layer responsibilities, file structure
- **[SWAGGER.md](docs/SWAGGER.md)** - API documentation guidelines
- **[DEPENDENCIES.md](docs/DEPENDENCIES.md)** - Technology stack and key dependencies

### Key Principles
- Follow Clean Architecture: domain → application (usecase) → infrastructure
- Keep functions small and focused on a single responsibility
- Use interfaces to define contracts between layers
- Write comprehensive tests with mocking
- Use conventional commits (see [CONTRIBUTING.md](CONTRIBUTING.md))

### File Structure
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

## Getting Started

1. Read the relevant documentation files based on your task
2. Follow the style guidelines in [STYLE_GUIDE.md](docs/STYLE_GUIDE.md)
3. Use the development commands from [DEVELOPMENT.md](docs/DEVELOPMENT.md)
4. Understand the architecture from [ARCHITECTURE.md](docs/ARCHITECTURE.md)
5. Check dependencies in [DEPENDENCIES.md](docs/DEPENDENCIES.md)

For detailed information on any topic, refer to the specific documentation files listed above.
