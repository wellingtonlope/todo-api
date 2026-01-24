Feature: Todo Update

  Background:
    Given the database is reset

  Scenario Outline: Successfully update a todo with all fields
    Given I have created a todo with title "Original Title", description "Original Description" and due_date ""
    And I have a todo update input with title "<title>", description "<description>" and due_date "<due_date>"
    When I update the todo with ID from the created todo
    Then the todo should be updated successfully with title "<title>", description "<description>" and due_date "<due_date>"

    Examples:
      | title          | description       | due_date              |
      | Updated Title  | Updated Desc      |                       |
      | Buy groceries  | Milk and bread    | 2030-12-31T23:59:59Z  |
      | Walk dog       |                   | 2030-06-30T10:00:00Z  |

  Scenario: Successfully update a todo with only title
    Given I have created a todo with title "Original Title", description "Original Description" and due_date "2030-12-31T23:59:59Z"
    And I have a todo update input with title "New Title Only", description "" and due_date ""
    When I update the todo with ID from the created todo
    Then the todo should be updated successfully with title "New Title Only", description "" and due_date ""

  Scenario: Successfully update a todo with title and description
    Given I have created a todo with title "Original Title", description "Original Description" and due_date "2030-12-31T23:59:59Z"
    And I have a todo update input with title "Updated Title", description "Updated description only" and due_date ""
    When I update the todo with ID from the created todo
    Then the todo should be updated successfully with title "Updated Title", description "Updated description only" and due_date ""

  Scenario: Fail to update a todo when not found
    When I update the todo with ID "nonexistent-id" with title "Updated Title", description "Updated Description" and due_date ""
    Then the update should fail with not found error

  Scenario Outline: Fail to update a todo with invalid input
    Given I have created a todo with title "Original Title", description "Original Description" and due_date ""
    When I update the todo with ID from the created todo with title "<title>", description "<description>" and due_date "<due_date>"
    Then the update should fail with validation error

    Examples:
      | title | description | due_date              |
      |       |             |                       |
      | Test  | Desc        | 2020-01-01T00:00:00Z  |