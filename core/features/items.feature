Feature: items
    Items should be able to be added to lists

    Scenario: adding an item to an empty list
        Given a list called "Shopping" is added
        When "Eggs" is added to the "Shopping" list
        And "Bread" is added to the "Shopping" list
        Then the "Shopping" list contains
            | Eggs  |
            | Bread |
        And the "Shopping" list has 2 items
