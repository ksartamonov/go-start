Feature: GetParameterValue Endpoint
  Scenario: Getting parameter values from the repository should return
  all parameter values, if they exist
    Given a parameter name "p1"
    When the GetParameterValue endpoint is called
    Then the response should contain all parameter values

  Scenario: Getting parameter values from the repository should return
  nil, if they not exist
    Given a parameter name "no_param"
    When the GetParameterValue endpoint is called
    Then the response should contain no values