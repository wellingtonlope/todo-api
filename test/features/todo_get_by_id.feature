Feature: Todo Get by ID

  Background:
    Given the database is reset

  Scenario Outline: Successfully get a todo by ID
    Given I have created a todo with title "<title>", description "<description>" and due_date "<due_date>"
    When I request the todo with ID from the created todo
    Then the todo should be retrieved successfully with title "<title>", description "<description>" and due_date "<due_date>"

    Examples:
      | title          | description       | due_date              |
      | Buy groceries  |                   |                       |
      | Buy groceries  | Milk and bread    |                       |
      | Buy groceries  | Milk and bread    | 2030-12-31T23:59:59Z  |

  Scenario: Fail to get a todo when not found
    When I request the todo with ID "nonexistent-id"
    Then the retrieval should fail with not found error