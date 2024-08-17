package utils

import (
	"fmt"
	"os"
	"tasks/models"
	"text/tabwriter"

	"github.com/mergestat/timediff"
)

func PrintTable(tasks []models.Task, all bool) {
	w := tabwriter.NewWriter(os.Stdout, 4, 0, 4, ' ', 0)

	if all {
		fmt.Fprintln(w, "ID\tTask\tCreated\tDone\tDue")

		for _, task := range tasks {
			isComplete := "-"
			if !task.IsComplete.IsZero() {
				isComplete = timediff.TimeDiff(task.IsComplete)
			}

			dueDate := "-"
			if !task.DueDate.IsZero() {
				dueDate = timediff.TimeDiff(task.DueDate)
			}
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", task.ID, task.Description, timediff.TimeDiff(task.CreatedAt), isComplete, dueDate)
		}

		w.Flush()
		return
	}

	fmt.Fprintln(w, "ID\tTask\tCreated\tDue")

	for _, task := range tasks {
		if task.IsComplete.IsZero() {
			dueDate := "-"
			if !task.DueDate.IsZero() {
				dueDate = timediff.TimeDiff(task.DueDate)
			}
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", task.ID, task.Description, timediff.TimeDiff(task.CreatedAt), dueDate)
		}
	}

	w.Flush()
}
