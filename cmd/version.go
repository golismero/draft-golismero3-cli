package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Golismero cli",
	Long:  "All software has versions. This is Golismero's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("3.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
