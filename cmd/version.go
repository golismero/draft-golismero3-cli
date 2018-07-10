package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version   = "undefined"
	BuildTime = "undefined"
	GitHash   = "undefined"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Golismero cli",
	Long:  "All software has versions. This is Golismero's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
