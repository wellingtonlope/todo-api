---
name: go-usecase-pattern
description: Create use cases following project pattern
---

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
1. Unexported type, exported constructor
2. Accept interfaces for dependencies
3. Return usecase.Error for errors
4. Use todo.TodoOutput or create custom Output
