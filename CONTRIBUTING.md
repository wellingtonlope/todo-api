# Contributing to Todo API

## Commit Message Convention

This project follows the [Conventional Commits](https://conventionalcommits.org/) specification to ensure clear and consistent commit messages. This helps with automated versioning, changelog generation, and easier code reviews.

### Format
```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

- `<type>`: Required. Describes the type of change.
- `[optional scope]`: Optional. Provides additional context (e.g., `api`, `ui`).
- `<description>`: Required. Brief description in imperative mood (e.g., "add feature" not "added feature").
- Body: Optional. Detailed explanation.
- Footer: Optional. For breaking changes or references (e.g., issues).

### Types
- `feat`: A new feature.
- `fix`: A bug fix.
- `docs`: Documentation changes.
- `style`: Code style changes (formatting, etc.).
- `refactor`: Code refactoring without functional changes.
- `test`: Adding or fixing tests.
- `chore`: Maintenance tasks (e.g., build scripts).
- `ci`: CI/CD changes.
- `build`: Build system changes.
- `perf`: Performance improvements.

### Examples
- `feat: add user authentication`
- `fix(api): handle null response errors`
- `docs: update README with setup instructions`
- `refactor: simplify error handling logic`

### Breaking Changes
- Use `!` after type/scope: `feat!: change API signature`
- Or add footer: `BREAKING CHANGE: description`

### Why Conventional Commits?
- Enables automatic semantic versioning.
- Improves changelog readability.
- Facilitates better collaboration and code history tracking.