Feature: Todo Creation

  Scenario Outline: Successfully create a todo
    Given I have a todo input with title "<title>", description "<description>" and due_date "<due_date>"
    When I create the todo
    Then the todo should be created successfully

    Examples:
      | title          | description       | due_date              |
      | Buy groceries  |                   |                       |
      | Buy groceries  | Milk and bread    |                       |
      | Buy groceries  | Milk and bread    | 2030-12-31T23:59:59Z  |

  Scenario: Fail to create a todo with invalid input
    Given I have a todo input with title ""
    When I create the todo
    Then the creation should fail with validation error

  Scenario: Fail to create a todo with due_date in the past
    Given I have a todo input with title "Test", description "Desc" and due_date "2020-01-01T00:00:00Z"
    When I create the todo
    Then the creation should fail with validation error