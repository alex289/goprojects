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

		var dueDate time.Time

		if len(args) > 1 {
			var layout string = "2006-01-02"

			if len(args[1]) > 10 {
				layout = "2006-01-02 15:04:05"
			}

			dueDate, err = time.Parse(layout, args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Invalid due date")
				os.Exit(1)
			}
		} else {
			dueDate = time.Time{}
		}

		task := models.Task{
			ID:          maxID + 1,
			Description: args[0],
			CreatedAt:   time.Now(),
			DueDate:     dueDate,
			IsComplete:  time.Time{},
		}

		utils.SaveTasks(append(tasks, task))
	},
}
