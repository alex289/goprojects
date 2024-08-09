package cmd

import (
	"fmt"
	"os"
	"tasks/models"
	"tasks/utils"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the list",
	Long:  `Add a task to the list and save it to the file`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := utils.LoadTasks()

		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to load tasks file")
			os.Exit(1)
		}

		maxID := 0
		for _, item := range tasks {
			if item.ID > maxID {
				maxID = item.ID
			}
		}

		task := models.Task{
			ID:          maxID + 1,
			Description: args[0],
			CreatedAt:   time.Now(),
			IsComplete:  false,
		}

		utils.SaveTasks(append(tasks, task))
	},
}
