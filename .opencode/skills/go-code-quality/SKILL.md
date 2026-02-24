---
name: go-code-quality
description: Format and lint code
---

## Commands
```bash
make format   # gofumpt
make lint     # golangci-lint
make test     # run tests
```

## Rules
1. Run format before commit
2. Fix all lint errors
3. Ensure no unused imports
