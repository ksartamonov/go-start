Feature: GetByPair Endpoint
  Scenario: Getting data by pair should return entities, if they exist
    Given a property with parameter "p1" and value "v1"
    When the GetByPair endpoint is called
    Then the response should contain all entities with the given pair

  Scenario: Getting data by pair should return no entities, if they not exist
    Given a property with parameter "no_param" and value "no_value"
    When the GetByPair endpoint is called
    Then the response should have no entities