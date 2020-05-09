Feature: items
    Into a list I want to add and remove items

    Scenario: adding an item returns 200
        Given a "POST" request is sent to "/lists" with the body
            """
            {
                "name": "Shopping"
            }
            """
        When a "POST" request is sent to "/lists/test" with the body
            """
            {
                "item": "Eggs"
            }
            """
        Then the response code should be 200

    Scenario: Lists can be retrieved
        Given a "POST" request is sent to "/lists" with the body
            """
            {
                "name": "Shopping"
            }
            """
        And a "POST" request is sent to "/lists/test" with the body
            """
            {
                "item": "Eggs"
            }
            """
        When a "GET" request is sent to "/lists/test"
        Then the response code should be 200
        And the response should contain the following json
            """
            {
                "items": [
                    "Eggs"
                ]
            }
            """
