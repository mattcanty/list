Feature: api
    The API should be able to manage lists
    as an API should
    with proper HTTP codes

    Scenario: Make a new list
        When a "POST" request is sent to to "/lists" with the body
            """
            {
                "name": "Test"
            }
            """
        Then the response code should be 200
