# DEVELOPMENT.md

This document contains development commands, testing patterns, and code quality guidelines for the Todo API project.

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

Prefer using the `make` targets for consistency; the direct commands are equivalent and provided for reference.

## Testing Patterns

- Use table-driven tests with `t.Run()` for multiple scenarios
- Mock external dependencies using `testify/mock`
- Create mock types in the same test file or in a dedicated `*_mock.go` file
- Assert both expected results and error conditions
- Use `mock.Mock.AssertExpectations(t)` to verify all expectations were met

### Testing Guidelines

- Use `github.com/stretchr/testify/assert` and `github.com/stretchr/testify/mock`
- Create mock types that embed `mock.Mock`
- Use table-driven tests for multiple scenarios
- Name test cases clearly (e.g., "should fail when input is invalid")
- Always call `AssertExpectations()` on mocks

## Code Quality

### Code Formatting

Run `make format` (gofumpt) before committing. Gofumpt is stricter than gofmt: no unused imports, no redundant blank lines.

### Linting

Run `make lint` (golangci-lint) before committing.

### Commit Guidelines

Follow the commit message convention defined in [CONTRIBUTING.md](../CONTRIBUTING.md). Use Conventional Commits format (e.g., `feat: add new feature`, `fix: resolve bug`). Always read CONTRIBUTING.md before creating commits to ensure proper formatting.

### Language

All code and documentation must be written in English.
