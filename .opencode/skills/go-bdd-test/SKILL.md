---
name: go-bdd-test
description: Add BDD tests using godog/cucumber framework
---

## What I do
Create BDD tests with Gherkin syntax and godog steps.

## When to use me
Adding integration/acceptance tests.

## Pattern
1. Create step file: `test/steps/<feature>_steps.go`
2. Define context struct embedding BaseTestContext
3. Implement step methods
4. Register in InitializeScenario

## Structure
```go
type XxxContext struct { BaseTestContext }

func (c *XxxContext) StepName() error { ... }
func (c *XxxContext) InitializeScenario(ctx *godog.ScenarioContext) {
    ctx.Step(`^pattern$`, c.StepName)
}
```

## Commands (use Makefile)
```bash
# Run BDD tests
make test

# Run specific test
go test ./test -v
```

## Rules
1. Use validation_helpers.go for common validations
2. Use HTTP client from http_client.go
3. Register context in suite_test.go
4. Use Given/When/Then pattern