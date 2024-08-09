package cmd

import (
	"fmt"
	"os"
	"strconv"
	"tasks/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from the list",
	Long:  `Delete a task from the list by id by removing it from the file`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := utils.LoadTasks()

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
				tasks = append(tasks[:i], tasks[i+1:]...)
				utils.SaveTasks(tasks)
				break
			}
		}

		if !taskFound {
			fmt.Fprintln(os.Stderr, "Task not found")
			os.Exit(1)
		}
	},
}
