Feature: Basic tool usage
    As user of golismero 3, i want lunch a security scan against a target.
    For that i need a golismero server and golismero cli.
    Using only those tools i must be able to get some security info.

    Background: Golismero server
        Given a golismero server
        And the list of default plugins
    
    @wip
    Scenario: Basic scan against url
        As a basic user i want.

        Given a golismero cli
        And access thru cli to the server configured
        And the correct permisions
        Given a target "http://example.com"
        When i execute the command "g3cli scan"
        Then i want to see th e scan progess
        And the tool send the scan info to stdout

    @wip
    Scenario: Basic plugin discovery
        As a basic user i want to know how many plugins are and what they do.

        Given a golismero cli
        And access thru cli to the server configured
        And the correct permisions
        When i execute the command "g3cli plugins"
        Then i recieve the list of available plugins with their short description

    @wip
    Scenario: Basic plugin details
        As a basic user i want to know how to use a given plugin.

        Given a golismero cli
        And access thru cli to the server configured
        And the correct permisions
        When i execute the command "g3cli plugins dns-recon"
        Then i recieve the complete plugin description
        And i recieve the list of parameters
        And i recieve the info of the basic command
        And i recieve the url of plugin repository