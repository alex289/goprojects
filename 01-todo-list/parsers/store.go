package parsers

import (
	"fmt"
	"tasks/models"
	"tasks/parsers/csv"
	"tasks/parsers/json"
	"tasks/parsers/sqlite"
)

func LoadTasks(parser string) ([]models.Task, error) {
	switch parser {
	case "csv":
		return csv.LoadTasks()
	case "sqlite":
		return sqlite.LoadTasks()
	case "json":
		return json.LoadTasks()
	default:
		return nil, fmt.Errorf("invalid parser: %s", parser)
	}
}

func SaveTasks(tasks []models.Task, parser string) error {
	switch parser {
	case "csv":
		return csv.SaveTasks(tasks)
	case "sqlite":
		return sqlite.SaveTasks(tasks)
	case "json":
		return json.SaveTasks(tasks)
	default:
		return fmt.Errorf("invalid parser: %s", parser)
	}
}
