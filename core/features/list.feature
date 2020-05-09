Feature: list
    Lists are the main feature
    They can be added and they can be removed

    Scenario: Make a new list
        Given there are 0 lists
        When a list called "Test" is added
        Then there should be 1 list
