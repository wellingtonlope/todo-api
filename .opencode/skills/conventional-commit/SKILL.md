---
name: conventional-commit
description: Format commit messages following Conventional Commits specification
---

## What I do
Format commit messages following the project's CONTRIBUTING.md.

## When to use me
Before creating any git commit.

## Format
```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

## Types
- feat: new feature
- fix: bug fix
- docs: documentation
- style: code style changes
- refactor: refactoring
- test: adding/fixing tests
- chore: maintenance
- ci: CI/CD changes
- build: build system
- perf: performance improvements

## Rules
1. Use imperative mood ("add" not "added")
2. Description is required
3. Breaking change: use `!` after type or BREAKING CHANGE footer
4. Keep first line under 72 characters