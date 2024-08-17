package cmd

import (
	"fmt"
	"os"
	"tasks/parsers"
	"tasks/utils"

	"github.com/spf13/cobra"
)

var all bool

func init() {
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "Show all tasks")
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print all tasks in the list",
	Long:  `Print all tasks in the list in table format`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := parsers.LoadTasks(parser)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to load tasks file")
			os.Exit(1)
		}

		utils.PrintTable(tasks, all)
	},
}
