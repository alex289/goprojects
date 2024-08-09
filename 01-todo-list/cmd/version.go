package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Tasks",
	Long:  `All software has versions. This is Tasks's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tasks CLI v0.1")
	},
}
