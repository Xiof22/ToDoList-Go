package models

import (
	"fmt"
	"github.com/dustin/go-humanize"
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

	createdAt := humanize.Time(t.CreatedAt)
	deadline := ""
	updatedAt := ""

	if !t.Deadline.IsZero() {
		deadline = fmt.Sprintf("Deadline: %s\n", humanize.Time(t.Deadline))
	}

	if !t.UpdatedAt.IsZero() {
		updatedAt = fmt.Sprintf("Updated: %s\n", humanize.Time(t.UpdatedAt))
	}

	return fmt.Sprintf("ID: %d\nTitle: %s\nDescription: %s\nCompleted: %t\nCreated: %s\n%s%s\n",
		t.ID, t.Title, desc, t.IsCompleted, createdAt, deadline, updatedAt)
}
