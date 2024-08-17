package cmd

import (
	"fmt"
	"os"
	"strconv"
	"tasks/parsers"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completeCmd)
}

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a task in the list",
	Long:  `Complete a task in the list by id and save it to the file`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := parsers.LoadTasks(parser)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to load tasks file")
			os.Exit(1)
		}

		id, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid task id")
			os.Exit(1)
		}

		taskFound := false

		for i, task := range tasks {
			if task.ID == id {
				taskFound = true
				tasks[i].IsComplete = time.Now()
				break
			}
		}

		if !taskFound {
			fmt.Fprintln(os.Stderr, "Task not found")
			os.Exit(1)
		}

		parsers.SaveTasks(tasks, parser)
	},
}
