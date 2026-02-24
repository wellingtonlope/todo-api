---
name: go-unit-test
description: Write table-driven unit tests with testify
---

## Pattern
```go
func TestXxx_Handle(t *testing.T) {
    tests := []struct {
        name   string
        mock   *xxxStoreMock
        input  xxx.Input
        output xxx.Output
        err    error
    }{
        {name: "success", mock: mockSuccess(), input: ..., output: ..., err: nil},
        {name: "fail", mock: mockFail(), input: ..., output: nil, err: usecase.NewError("msg", err, type)},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            uc := xxx.NewXxx(tc.mock)
            got, err := uc.Handle(ctx, tc.input)
            assert.Equal(t, tc.output, got)
            assert.Equal(t, tc.err, err)
            tc.mock.AssertExpectations(t)
        })
    }
}

type xxxStoreMock struct { mock.Mock }
func (m *xxxStoreMock) Method(ctx context.Context, id string) (*domain.Xxx, error) {
    args := m.Called(ctx, id)
    return args.Get(0).(*domain.Xxx), args.Error(1)
}
```

## Commands
```bash
make test
go test ./internal/app/usecase/todo -run TestXxx -v
```

## Rules
1. Table-driven with testCases slice
2. Define mocks at bottom, embed mock.Mock
3. Name: TestUsecaseName_Method
4. Use factory functions for mocks in test cases
