# AGENTS.md

This document provides guidelines for agentic coding assistants working on this codebase.

## Project Overview

This is a Todo API built with Go 1.25.6, using Echo for HTTP, GORM with SQLite for persistence, and Uber FX for dependency injection. The project follows Clean Architecture with clear separation between domain, application (usecase), and infrastructure layers.

## Quick Reference

### Essential Commands
```bash
# Run tests
make test

# Build and run
make build && make server

# Format and lint
make format && make lint

# Generate Swagger docs
make swagger
```

### Prerequisites
- Go 1.25.6 installed
- Tools: `gofumpt`, `golangci-lint`, `swag`
- See [DEVELOPMENT.md](docs/DEVELOPMENT.md) for installation details

### Documentation Structure
- **[DEVELOPMENT.md](docs/DEVELOPMENT.md)** - Build commands, testing patterns, code quality
- **[STYLE_GUIDE.md](docs/STYLE_GUIDE.md)** - Naming conventions, code style, error handling
- **[ARCHITECTURE.md](docs/ARCHITECTURE.md)** - Design patterns, layer responsibilities, file structure
- **[SWAGGER.md](docs/SWAGGER.md)** - API documentation guidelines
- **[DEPENDENCIES.md](docs/DEPENDENCIES.md)** - Technology stack and key dependencies
- **[CONTRIBUTING.md](CONTRIBUTING.md)** - Commit conventions

### Key Principles
- Follow Clean Architecture: domain → application (usecase) → infrastructure
- Keep functions small and focused on a single responsibility
- Use interfaces to define contracts between layers
- Write comprehensive tests with mocking
- All code and documentation must be written in English

### File Structure
```
cmd/api/              # Application entrypoint
docs/                 # Project docs and generated Swagger outputs
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

### Docker
```bash
docker compose up
docker build -t todo-api .
```

## Getting Started

### Quick Setup
```bash
make build
make test
make server

# Swagger (optional)
make swagger
```

1. Read the relevant documentation files based on your task
2. Follow the style guidelines in [STYLE_GUIDE.md](docs/STYLE_GUIDE.md)
3. Use the development commands from [DEVELOPMENT.md](docs/DEVELOPMENT.md)
4. Understand the architecture from [ARCHITECTURE.md](docs/ARCHITECTURE.md)
5. Check dependencies in [DEPENDENCIES.md](docs/DEPENDENCIES.md)

For detailed information on any topic, refer to the specific documentation files listed above.
