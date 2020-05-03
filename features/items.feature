Feature: items
    Into a list I want to add and remove items

    Scenario: adding an item to an empty list
        Given a list called "Shopping" is added
        When "Eggs" is added to the "Shopping" list
        And "Bread" is added to the "Shopping" list
        Then the "Shopping" list contains
            | Eggs  |
            | Bread |
        And the "Shopping" list has 2 items
