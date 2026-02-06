---
name: go-code-quality
description: Run formatting and linting before commits
---

## What I do
Ensure code quality with formatting and linting.

## When to use me
Before committing changes.

## Commands
```bash
# Format code
gofumpt -w .

# Or via Makefile
make format

# Lint code
golangci-lint run

# Or via Makefile
make lint

# Run both
make format && make lint
```

## Rules
1. Always run format before commit
2. Fix all lint errors
3. Run tests after: `make test`
4. Ensure no unused imports