package models

import (
	"github.com/google/uuid"
	"time"
)

type TaskID uuid.UUID

func (id TaskID) String() string {
	return uuid.UUID(id).String()
}

type Task struct {
	ID          TaskID
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   time.Time
	Deadline    time.Time
	UpdatedAt   *time.Time
}

func NewTask(title, description string, deadline time.Time) Task {
	return Task{
		ID:          TaskID(uuid.New()),
		Title:       title,
		Description: description,
		Deadline:    deadline,
	}
}
