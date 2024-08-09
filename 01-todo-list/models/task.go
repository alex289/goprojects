package models

import "time"

type Task struct {
	ID          int
	Description string
	CreatedAt   time.Time
	IsComplete  bool
}
