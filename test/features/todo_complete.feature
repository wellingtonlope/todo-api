Feature: Todo Complete

  Background:
    Given the database is reset

  Scenario: Successfully mark a todo as complete
    Given I have a todo with title "Buy groceries" and description "Milk and bread"
    And I create the todo
    And the todo should be created successfully
    When I mark the todo as complete
    Then the todo should be marked as completed successfully

  Scenario: Fail to complete when todo not found
    When I mark todo with ID "non-existent-id" as complete
    Then the completion should fail with not found error

  Scenario: Successfully complete an already completed todo
    Given I have a todo with title "Read book" and description "Read 10 pages"
    And I create the todo
    And the todo should be created successfully
    When I mark the todo as complete
    Then the todo should be marked as completed successfully
    When I mark the todo as complete again
    Then the todo should be marked as completed successfully