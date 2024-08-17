package json

import (
	"encoding/json"
	"fmt"
	"os"
	"tasks/models"
	"tasks/utils"
)

const fileName = "tasks.json"

func LoadTasks() ([]models.Task, error) {
	file, err := utils.LoadFile(fileName)

	if err != nil {
		return nil, err
	}

	defer utils.CloseFile(file)

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if fileInfo.Size() == 0 {
		return []models.Task{}, nil
	}

	reader := json.NewDecoder(file)

	var tasks []models.Task
	err = reader.Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(tasks []models.Task) error {
	file, err := os.Create(fileName)

	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer file.Close()

	writer := json.NewEncoder(file)
	writer.SetIndent("", "  ")

	err = writer.Encode(tasks)
	if err != nil {
		return fmt.Errorf("failed to encode tasks to JSON: %w", err)
	}

	return nil
}
