Feature: Todo Mark Pending

  Background:
    Given the database is reset

  Scenario: Successfully mark a todo as pending
    Given I have created a todo with title "Completed Task" and description "Details"
    And I mark the todo as complete
    When I mark the todo as pending
    Then the todo should be marked as pending successfully

  Scenario: Fail to mark pending a todo when not found
    When I mark the todo with ID "nonexistent-id" as pending
    Then the mark pending should fail with not found error

  Scenario: Fail to mark pending an already pending todo
    Given I have created a todo with title "Pending Task" and description "Details"
    When I mark the todo as pending
    Then the mark pending should fail with validation error