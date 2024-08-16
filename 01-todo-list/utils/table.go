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
		fmt.Fprintln(w, "ID\tTask\tCreated\tDone")

		for _, task := range tasks {
			isComplete := "-"
			if !task.IsComplete.IsZero() {
				isComplete = timediff.TimeDiff(task.IsComplete)
			}
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", task.ID, task.Description, timediff.TimeDiff(task.CreatedAt), isComplete)
		}

		w.Flush()
		return
	}

	fmt.Fprintln(w, "ID\tTask\tCreated")

	for _, task := range tasks {
		if task.IsComplete.IsZero() {
			fmt.Fprintf(w, "%d\t%s\t%s\n", task.ID, task.Description, timediff.TimeDiff(task.CreatedAt))
		}
	}

	w.Flush()
}
