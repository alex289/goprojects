package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Tasks is a cli tool for managing tasks",
	Long:  `Tasks is a cli tool for managing tasks. It is inspired by the todo.txt format. Tasks is a simple command line tool that allows you to add, list, complete and delete tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tasks CLI v0.1")
		fmt.Println("Use 'tasks help' to see all available commands")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
