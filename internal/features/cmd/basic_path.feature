Feature: Basic local tool usage
    As user of golismero 3, i want use some formating and reporting tools
    without a server locally.

    @wip
    Scenario: tool invoke without params
        Given no params
        When I call "g3cli" command
        Then tool show me the command help