package models

import "time"

type Task struct {
	ID          int
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   time.Time
	Deadline    time.Time
	UpdatedAt   time.Time
}
