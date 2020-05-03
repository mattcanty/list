Feature: list
    In order to add items to a list
    As a user
    I need to be able to add lists

    Scenario: Make a new list
        Given there are 0 lists
        When a list called "Test" is added
        Then there should be 1 list
