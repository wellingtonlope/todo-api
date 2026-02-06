---
name: go-clean-architecture
description: Create components following Clean Architecture pattern (domain -> usecase -> infra)
---

## What I do
Create new features with proper Clean Architecture separation:
- Domain: entities, validation, pure business logic
- Usecase: application logic, orchestration  
- Infra: HTTP handlers, database implementations

## When to use me
Adding new domains, features, or resources to the API.

## Pattern
File structure per feature:
```
internal/
  domain/<feature>.go
  app/usecase/<feature>/
    create.go, update.go, delete.go, get.go
  infra/handler/<feature>_create.go
  infra/gorm/<feature>.go
  infra/memory/<feature>.go
```

## Rules
1. Domain has NO external dependencies
2. Usecases accept interfaces for dependencies
3. Handlers only handle HTTP concerns
4. Always create interface + constructor pattern