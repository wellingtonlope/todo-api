# BDD Improvement Roadmap

Este documento descreve as melhorias planejadas para os testes BDD do projeto Todo API, organizadas em estágios para implementação gradual.

## Estágio 1: Redução de Duplicação e Simplificação

### Objetivo
Eliminar código repetido e simplificar a configuração dos testes BDD.

### Tarefas

- [ ] **Criar TestFactory**
  - Implementar fábrica para reduzir duplicação em `test/suite_test.go:149-262`
  - Criar método `SetupBDDTest(featureType)` que retorna dependências configuradas
  - Reduzir código repetido de ~100 linhas para ~20 linhas

- [ ] **Simplificar configuração de Echo App**
  - Refatorar `setupEchoApp()` para aceitar enum de features em vez de múltiplos booleanos
  - Criar `FeatureFlags` struct com configurações predefinidas

- [ ] **Unificar setup de database**
  - Combinar `setupDatabase()` e `setupDependencies()` em único método
  - Adicionar método `Reset()` para limpar estado entre testes

### Arquivos a modificar
- `test/suite_test.go`
- `test/steps/base_context.go`

---

## Estágio 2: Biblioteca de Asserts BDD

### Objetivo
Centralizar validações e criar asserts mais expressivos.

### Tarefas

- [ ] **Criar BDD Asserts Library**
  - Extrair validações de `test/steps/validation.go` para biblioteca dedicada
  - Criar asserts com nomes de negócio (ex: `ShouldHaveCreatedTodo`, `ShouldReturnValidationError`)
  - Adicionar mensagens de erro mais claras

- [ ] **Implementar Response Validators**
  - Criar validadores específicos para diferentes tipos de resposta
  - Separar validação de headers, body e status code
  - Adicionar suporte a validações parciais

- [ ] **Criar Test Data Builders**
  - Implementar builders para criar dados de teste de forma fluente
  - Ex: `TodoBuilder().WithTitle("Buy groceries").WithDueDate("2025-01-01").Build()`

### Arquivos a criar/modificar
- `test/asserts/bdd_asserts.go`
- `test/builders/todo_builder.go`
- `test/steps/validation.go` (refactor)

---

## Estágio 3: Melhoria nos Feature Files

### Objetivo
Tornar os cenários mais expressivos e reutilizáveis.

### Tarefas

- [ ] **Revisar Linguagem de Negócio**
  - Substituir linguagem técnica por termos de negócio
  - Ex: "response should be successful with status 200" → "operation should complete successfully"
  - Adicionar contexto de usuário quando aplicável

- [ ] **Implementar Steps Reutilizáveis**
  - Criar steps genéricos que possam ser reutilizados entre features
  - Ex: `Given I have valid todo data`, `When I perform the operation`
  - Reduzir duplicação de steps similares

- [ ] **Adicionar Scenario Outlines**
  - Converter cenários repetidos para Scenario Outlines
  - Melhorar cobertura de testes com exemplos variados
  - Facilitar manutenção de casos de teste

### Arquivos a modificar
- `test/features/*.feature`
- `test/steps/common_steps.go` (novo)

---

## Estágio 4: Gerenciamento de Estado e Isolamento

### Objetivo
Melhorar isolamento entre testes e gerenciar estado de forma mais eficiente.

### Tarefas

- [ ] **Implementar Test Context Manager**
  - Criar gerenciador centralizado para estado de testes
  - Garantir isolamento completo entre cenários
  - Adicionar suporte a teardown automático

- [ ] **Otimizar Database Reset**
  - Implementar transações para rollback mais rápido
  - Evitar deletes manuais como em `test/steps/base_context.go:29`
  - Adicionar seeds para dados de teste consistentes

- [ ] **Adicionar Test Fixtures**
  - Criar fixtures para dados comuns de teste
  - Implementar sistema de tags para categorizar fixtures
  - Facilitar criação de cenários complexos

### Arquivos a criar/modificar
- `test/context/manager.go`
- `test/fixtures/todo_fixtures.go`
- `test/steps/base_context.go` (refactor)

---

## Estágio 5: Documentação e Ferramentas

### Objetivo
Documentar padrões e criar utilitários para facilitar o desenvolvimento.

### Tarefas

- [ ] **Criar BDD Guidelines**
  - Documentar padrões para writing features e steps
  - Criar guia de boas práticas específico do projeto
  - Adicionar exemplos e templates

- [ ] **Implementar BDD Commands**
  - Adicionar comandos Make específicos para BDD
  - Ex: `make test-bdd`, `make test-bdd-verbose`
  - Integrar com fluxo de CI/CD

- [ ] **Criar Step Generation Tools**
  - Implementar utilitário para scaffolding de novos steps
  - Adicionar validação automática de step definitions
  - Criar linter para boas práticas de BDD

### Arquivos a criar/modificar
- `docs/BDD_GUIDELINES.md`
- `Makefile`
- `tools/step_generator.go`

---

## Estágio 6: Integração e Performance

### Objetivo
Integrar BDD com ecossistema e otimizar performance.

### Tarefas

- [ ] **Integrar com Test Reports**
  - Configurar relatórios HTML/Cucumber para BDD
  - Adicionar métricas de cobertura específicas
  - Integrar com ferramentas de CI

- [ ] **Otimizar Performance**
  - Implementar paralelização segura de testes
  - Reduzir tempo de setup/teardown
  - Adicionar caching para operações repetidas

- [ ] **Adicionar Test Data Management**
  - Implementar gerador de dados de teste faker
  - Criar estratégias para edge cases
  - Adicionar suporte a localization/internationalization

### Arquivos a criar/modificar
- `test/performance/optimizer.go`
- `test/data/faker.go`
- Configurações de CI

---

## Como Usar Este Roadmap

### Para cada estágio:
1. **Marque as tarefas** com `[x]` conforme completar
2. **Execute os testes** após cada estágio: `make test`
3. **Verifique a cobertura** para garantir regressões
4. **Documente aprendizados** no final de cada estágio

### Dependências:
- Estágio 1 é prerequisite para todos os outros
- Estágios 2-4 podem ser feitos em paralelo após o 1
- Estágios 5-6 dependem dos anteriores

### Métricas de Sucesso:
- Redução de linhas de código em ~40%
- Tempo de execução dos testes BDD reduzido em ~30%
- Novos testes BDD 50% mais rápidos de escrever
- Cobertura mantida ou aumentada

---

## Checklist Final

Ao completar todos os estágios:
- [ ] Todos os testes BDD existentes continuam passando
- [ ] Novo teste BDD pode ser criado em <5 minutos
- [ ] Documentação está completa e acessível
- [ ] Performance dos testes é aceitável
- [ ] Equipe está treinada nos novos padrões

**Importante**: Commit após cada estágio para facilitar rollback se necessário.