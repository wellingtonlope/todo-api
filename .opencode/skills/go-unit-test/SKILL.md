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
    testCases := []struct {
        name   string
        store  *xxxStoreMock
        input  xxx.Input
        result xxx.Output
        err    error
    }{
        {
            name: "should fail when store fails",
            store: func() *xxxStoreMock {
                m := new(xxxStoreMock)
                m.On("Method", context.TODO(), mock.Anything).
                    Return(nil, assert.AnError).Once()
                return m
            }(),
            input:  xxx.Input{},
            result: nil,
            err:    usecase.NewError("fail to xxx", assert.AnError, usecase.ErrorTypeInternalError),
        },
        {
            name: "should succeed",
            store: func() *xxxStoreMock {
                m := new(xxxStoreMock)
                m.On("Method", context.TODO(), mock.Anything).
                    Return(&domain.Xxx{ID: "123"}, nil).Once()
                return m
            }(),
            input:  xxx.Input{ID: "123"},
            result: &xxx.Output{ID: "123"},
            err:    nil,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            uc := xxx.NewXxx(tc.store)
            result, err := uc.Handle(context.TODO(), tc.input)
            assert.Equal(t, tc.result, result)
            assert.Equal(t, tc.err, err)
            tc.store.AssertExpectations(t)
        })
    }
}

type xxxStoreMock struct {
    mock.Mock
}

func (m *xxxStoreMock) Method(ctx context.Context, id string) (*domain.Xxx, error) {
    args := m.Called(ctx, id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*domain.Xxx), args.Error(1)
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
1. Use table-driven tests with `testCases` slice and `t.Run()`
2. Define mocks at the bottom of the test file, embedding `mock.Mock`
3. Name tests: `TestUsecaseName_Handle` (e.g., `TestList_Handle`)
4. Create mock instances using factory functions inside test cases: `func() *xxxMock { ... }()`
5. Use `assert.Equal(t, tc.result, result)` and `assert.Equal(t, tc.err, err)` for assertions
6. Call `tc.mock.AssertExpectations(t)` to verify mock expectations
7. Use `usecase.NewError()` for error assertions with proper error types
8. Store mocks directly in the test case struct (e.g., `store *xxxStoreMock`)