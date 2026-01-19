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

func NewTask(title, description string, deadline time.Time) Task {
	return Task{
		Title:       title,
		Description: description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		Deadline:    deadline,
	}
}
