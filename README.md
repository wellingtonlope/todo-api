# Todo API

![CI](https://github.com/wellingtonlope/todo-api/workflows/CI/badge.svg)
![Go Version](https://img.shields.io/github/go-mod/go-version/wellingtonlope/todo-api)
![License](https://img.shields.io/badge/license-MIT-blue)
![Codecov](https://codecov.io/gh/wellingtonlope/todo-api/branch/main/graph/badge.svg)
![Tests](https://img.shields.io/badge/tests-passing-brightgreen)

A RESTful API for managing todo items built with Go, following Clean Architecture principles.

## Features

- Create, read, update, and delete todos
- Mark todos as completed or pending
- Set optional due dates
- Input validation and error handling
- Swagger/OpenAPI documentation

## Architecture

This project follows **Clean Architecture** with three layers:

- **Domain**: Business entities and rules (pure Go, no dependencies)
- **Application**: Use cases and business logic orchestration
- **Infrastructure**: HTTP handlers (Echo), database (GORM + SQLite), and DI (Uber FX)

## Tech Stack

- **Go 1.25.6**
- **Echo** - HTTP web framework
- **GORM** - ORM for database operations
- **SQLite** - Database (easily swappable)
- **Uber FX** - Dependency injection
- **Swagger** - API documentation
- **Testify + Godog** - Testing (unit and BDD)

## Quick Start

### Prerequisites

- Go 1.25.6+
- Make (optional)

### Installation

```bash
# Clone the repository
git clone <repository-url>
cd todo-api

# Install dependencies
go mod download

# Copy environment variables (optional)
cp .env.example .env
```

### Running

```bash
# Development mode
make server

# Or directly
go run ./cmd/api/
```

The API will be available at `http://localhost:1323`

### API Documentation

Access Swagger UI at: `http://localhost:1323/swagger/index.html`

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/todos` | Create a new todo |
| GET | `/todos` | List all todos |
| GET | `/todos/:id` | Get a specific todo |
| PUT | `/todos/:id` | Update a todo |
| DELETE | `/todos/:id` | Delete a todo |
| PUT | `/todos/:id/complete` | Mark todo as completed |
| PUT | `/todos/:id/pending` | Mark todo as pending |

## Development Commands

```bash
make test       # Run all tests
make format     # Format code with gofumpt
make lint       # Run golangci-lint
make swagger    # Generate Swagger documentation
make build      # Build binary
make all        # Format + lint + test
```

## Docker

```bash
# Run with Docker Compose
docker compose up

# Or build and run manually
docker build -t todo-api .
docker run -p 1323:1323 todo-api
```

## Project Structure

```
cmd/api/              # Application entrypoint
internal/
  domain/             # Business entities
  app/usecase/todo/   # Use cases (business logic)
  infra/
    handler/          # HTTP handlers
    gorm/             # GORM repositories
    memory/           # In-memory repositories (testing)
  bootstrap/          # Dependency injection setup
pkg/clock/            # Time utilities
test/                 # BDD tests (Godog)
docs/                 # Documentation and Swagger files
```

## Testing

```bash
# Run unit tests
make test

# Run BDD tests
go test ./test/...

# Run with coverage
go test ./... -cover
```

## Configuration

Environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_ENV` | Application environment | `development` |
| `PORT` | HTTP server port | `1323` |

## Documentation

- [Architecture](docs/ARCHITECTURE.md) - Design patterns and structure
- [Development](docs/DEVELOPMENT.md) - Build and testing guide
- [Dependencies](docs/DEPENDENCIES.md) - Technology stack details
- [Style Guide](docs/STYLE_GUIDE.md) - Coding conventions

## License

This project is licensed under the MIT License.

