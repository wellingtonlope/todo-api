---
name: go-unit-test
description: Write table-driven unit tests with testify/mock
---

## What I do
Create comprehensive unit tests with mocking.

## When to use me
Testing usecases, handlers, or domain logic.

## Pattern
```go
func TestXxx_Handle(t *testing.T) {
    tests := []struct {
        name    string
        mock    func(*mockStore, *mockClock)
        input   xxxInput
        want    xxxOutput
        wantErr error
    }{
        {name: "should succeed", ...},
        {name: "should fail when ...", ...},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            store := &mockStore{}
            clock := &mockClock{}
            tt.mock(store, clock)
            
            uc := NewXxx(store, clock)
            got, err := uc.Handle(context.Background(), tt.input)
            
            assert.Equal(t, tt.want, got)
            assert.Equal(t, tt.wantErr, err)
            store.AssertExpectations(t)
        })
    }
}
```

## Commands (use Makefile)
```bash
# Run all tests
make test

# Run specific test
go test ./internal/app/usecase/todo -run TestCreate_Handle -v
```

## Rules
1. Use table-driven tests with t.Run()
2. Create mocks embedding mock.Mock
3. Name tests: `TestPackage_Method`
4. Call AssertExpectations(t) on mocks