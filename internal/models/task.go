package models

import "github.com/google/uuid"

type TaskID uuid.UUID

func (id TaskID) String() string {
	return uuid.UUID(id).String()
}

type Task struct {
	ID          TaskID
	Title       string
	Description string
	IsCompleted bool
}

func NewTask(title, description string) Task {
	return Task{
		ID:          TaskID(uuid.New()),
		Title:       title,
		Description: description,
	}
}
