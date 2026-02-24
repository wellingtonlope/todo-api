---
name: go-clean-architecture
description: Create components following Clean Architecture pattern
---

## Pattern
```
internal/
  domain/<feature>.go      # entities, validation
  app/usecase/<feature>/   # business logic
  infra/handler/           # HTTP handlers
  infra/gorm/              # database
```

## Rules
1. Domain: NO external dependencies
2. Usecase: accept interfaces
3. Handler: HTTP only, delegate to usecase
4. Always use interface + constructor pattern
