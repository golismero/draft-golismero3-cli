package main

import "github.com/golismero/golismero3-cli/cmd"

var (
	Version   = "undefined"
	BuildTime = "undefined"
	GitHash   = "undefined"
)

func main() {
	cmd.Version = Version
	cmd.BuildTime = BuildTime
	cmd.GitHash = GitHash
	cmd.Execute()
}
