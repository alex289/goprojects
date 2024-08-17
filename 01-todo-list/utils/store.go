package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"tasks/models"
	"time"
)

const fileName = "tasks.csv"

func LoadTasks() ([]models.Task, error) {
	file, err := loadFile(fileName)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 5
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	closeFile(file)

	if len(data) == 0 {
		return []models.Task{}, nil
	}

	tasks := []models.Task{}

	for _, row := range data[1:] {
		id, err := strconv.Atoi(row[0])

		if err != nil {
			return nil, err
		}

		createdAt, err := time.Parse(time.RFC3339, row[2])

		if err != nil {
			return nil, err
		}

		isComplete, err := time.Parse(time.RFC3339, row[3])

		if err != nil {
			return nil, err
		}

		dueDate, err := time.Parse(time.RFC3339, row[4])

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, models.Task{
			ID:          id,
			Description: row[1],
			CreatedAt:   createdAt,
			IsComplete:  isComplete,
			DueDate:     dueDate,
		})
	}

	return tasks, nil
}

func SaveTasks(tasks []models.Task) error {
	file, err := os.Create(fileName)

	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"ID", "Description", "CreatedAt", "IsComplete", "DueDate"}
	data := [][]string{}

	for _, row := range tasks {
		data = append(data, []string{
			strconv.Itoa(row.ID),
			row.Description,
			row.CreatedAt.Format(time.RFC3339),
			row.IsComplete.Format(time.RFC3339),
			row.DueDate.Format(time.RFC3339)})
	}

	writer.Write(headers)
	for _, row := range data {
		writer.Write(row)
	}

	return nil
}
