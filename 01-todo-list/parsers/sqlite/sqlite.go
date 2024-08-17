package sqlite

import (
	"database/sql"
	"fmt"
	"os"
	"tasks/models"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "tasks.db"

func openDB() (*sql.DB, error) {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		file, err := os.Create(dbName)
		if err != nil {
			return nil, fmt.Errorf("failed to create database file: %w", err)
		}
		file.Close()
	}

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return db, nil
}

func ensureTableExists(db *sql.DB) error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT,
		created_at TEXT,
		is_complete TEXT,
		due_date TEXT
	);
	`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}

func LoadTasks() ([]models.Task, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := ensureTableExists(db); err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, description, created_at, is_complete, due_date FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		var createdAt, isComplete, dueDate string

		err := rows.Scan(&task.ID, &task.Description, &createdAt, &isComplete, &dueDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		task.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return nil, fmt.Errorf("failed to parse CreatedAt: %w", err)
		}

		task.IsComplete, err = time.Parse(time.RFC3339, isComplete)
		if err != nil {
			return nil, fmt.Errorf("failed to parse IsComplete: %w", err)
		}

		task.DueDate, err = time.Parse(time.RFC3339, dueDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse DueDate: %w", err)
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during row iteration: %w", err)
	}

	return tasks, nil
}

func SaveTasks(tasks []models.Task) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := ensureTableExists(db); err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM tasks")

	if err != nil {
		return fmt.Errorf("failed to clear all tasks: %w", err)
	}

	for _, task := range tasks {
		insertOrUpdateQuery := `
		INSERT INTO tasks (id, description, created_at, is_complete, due_date) 
		VALUES (?, ?, ?, ?, ?)
		`

		_, err := db.Exec(insertOrUpdateQuery,
			task.ID,
			task.Description,
			task.CreatedAt.Format(time.RFC3339),
			task.IsComplete.Format(time.RFC3339),
			task.DueDate.Format(time.RFC3339))

		if err != nil {
			return fmt.Errorf("failed to insert or update task: %w", err)
		}
	}

	return nil
}
