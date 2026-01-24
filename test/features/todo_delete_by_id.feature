Feature: Todo Delete by ID

  Background:
    Given the database is reset

  Scenario: Successfully delete a todo by ID
    Given I have created a todo with title "Buy groceries", description "Milk and bread" and due_date ""
    When I delete the todo with ID from the created todo
    Then the todo should be deleted successfully

  Scenario: Fail to delete a todo when not found
    When I delete the todo with ID "nonexistent-id"
    Then the deletion should fail with not found error