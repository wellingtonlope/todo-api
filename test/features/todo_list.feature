Feature: Todo List

  Background:
    Given the database is reset

  Scenario: Get empty list when no todos exist
    When I request all todos
    Then the response should be successful with status 200
    And the response should contain an empty list of todos

  Scenario Outline: Successfully get list with multiple todos
    Given I have created a todo with title "<title1>", description "<description1>" and due_date "<due_date1>"
    And I have created a todo with title "<title2>", description "<description2>" and due_date "<due_date2>"
    When I request all todos
    Then the response should be successful with status 200
    And the response should contain a list with 2 todos
    And the first todo should have title "<title1>", description "<description1>" and due_date "<due_date1>"
    And the second todo should have title "<title2>", description "<description2>" and due_date "<due_date2>"

    Examples:
      | title1         | description1    | due_date1              | title2         | description2    | due_date2              |
      | Buy groceries  | Milk and bread  |                        | Walk the dog   |                 |                        |
      | Buy groceries  | Milk and bread  | 2030-12-31T23:59:59Z  | Walk the dog   | Park exercise   | 2030-12-30T23:59:59Z  |

  Scenario: Successfully get list with single todo
    Given I have created a todo with title "Read book", description "Technical book" and due_date ""
    When I request all todos
    Then the response should be successful with status 200
    And the response should contain a list with 1 todo
    And the first todo should have title "Read book", description "Technical book" and due_date ""

  Scenario: Get empty list when filtering by status and no todos match
    Given I have created a todo with title "Pending task", description "" and due_date ""
    When I request todos with status "completed"
    Then the response should be successful with status 200
    And the response should contain an empty list of todos

  Scenario: Get only pending todos when filtering by pending status
    Given I have created a todo with title "Pending task 1", description "" and due_date ""
    And I have created a todo with title "Pending task 2", description "" and due_date ""
    And I have created a completed todo with title "Completed task", description "" and due_date ""
    When I request todos with status "pending"
    Then the response should be successful with status 200
    And the response should contain a list with 2 todos

  Scenario: Get only completed todos when filtering by completed status
    Given I have created a todo with title "Pending task", description "" and due_date ""
    And I have created a completed todo with title "Completed task 1", description "" and due_date ""
    And I have created a completed todo with title "Completed task 2", description "" and due_date ""
    When I request todos with status "completed"
    Then the response should be successful with status 200
    And the response should contain a list with 2 todos
