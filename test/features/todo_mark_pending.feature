Feature: Todo Mark Pending

  Background:
    Given the database is reset

  Scenario: Successfully mark a todo as pending
    Given I have created a todo with title "Buy groceries", description "Milk and bread" and due_date ""
    And I mark the todo with ID from the created todo as pending
    Then the todo should be marked as pending successfully

  Scenario: Fail to mark a todo as pending when not found
    When I mark the todo with ID "nonexistent-id" as pending
    Then the marking as pending should fail with not found error

  Scenario: Fail to mark an already pending todo as pending
    Given I have created a todo with title "Buy groceries", description "Milk and bread" and due_date ""
    And I mark the todo with ID from the created todo as pending
    When I mark the todo with ID from the created todo as pending
    Then the marking as pending should succeed but status remains pending