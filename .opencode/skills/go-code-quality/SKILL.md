---
name: go-code-quality
description: Run formatting and linting before commits
---

## What I do
Ensure code quality with formatting and linting.

## When to use me
Before committing changes.

## Commands (use Makefile)
```bash
# Format code
make format

# Lint code
make lint

# Run both + tests
make all

# Direct commands (avoid if possible)
gofumpt -w .
golangci-lint run
```

## Rules
1. Always run format before commit
2. Fix all lint errors
3. Run tests after: `make test`
4. Ensure no unused imports