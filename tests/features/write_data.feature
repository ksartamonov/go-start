Feature: WriteData Endpoint
  Scenario: Writing data to the repository
    Given a request to save data
    When the WriteData endpoint is called
    Then the data should be written successfully