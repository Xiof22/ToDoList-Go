package models

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   time.Time
	Deadline    time.Time
	UpdatedAt   time.Time
}

func (t Task) String() string {
	desc := t.Description
	if desc == "" {
		desc = "none"
	}

	deadline := "whenever"
	updatedAt := "none"
	createdAt := t.CreatedAt.Format(time.DateTime)

	if !t.Deadline.IsZero() {
		deadline = t.Deadline.Format(time.DateTime)
	}

	if !t.UpdatedAt.IsZero() {
		updatedAt = t.UpdatedAt.Format(time.DateTime)
	}

	return fmt.Sprintf("ID: %d\nTitle: %s\nDescription: %s\nCompleted: %t\nCreatedAt: %s\nDeadline: %s\nUpdatedAt: %s\n\n",
		t.ID, t.Title, desc, t.IsCompleted, createdAt, deadline, updatedAt)
}
