---
name: go-usecase-pattern
description: Create use cases following the project's established pattern
---

## What I do
Create use case implementations with consistent structure.

## When to use me
Adding new business operations.

## Pattern
```go
type (
    XxxInput struct { ... }
    XxxStore interface { ... }
    Xxx interface { Handle(context.Context, XxxInput) (XxxOutput, error) }
    xxx struct { store XxxStore; clock usecase.Clock }
)

func NewXxx(store XxxStore, clock usecase.Clock) *xxx { ... }
func (uc *xxx) Handle(ctx context.Context, input XxxInput) (XxxOutput, error) { ... }
```

## Rules
1. Unexported implementation type, exported constructor
2. Accept interfaces for dependencies
3. Return usecase.Error for errors
4. Use todo.TodoOutput or create custom Output type