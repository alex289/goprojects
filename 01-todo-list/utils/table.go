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
			fmt.Fprintf(w, "%d\t%s\t%s\t%t\n", task.ID, task.Description, timediff.TimeDiff(task.CreatedAt), task.IsComplete)
		}

		w.Flush()
		return
	}

	fmt.Fprintln(w, "ID\tTask\tCreated")

	for _, task := range tasks {
		if !task.IsComplete {
			fmt.Fprintf(w, "%d\t%s\t%s\n", task.ID, task.Description, timediff.TimeDiff(task.CreatedAt))
		}
	}

	w.Flush()
}
