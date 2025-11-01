package models

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound        = errors.New("Task not found")
	ErrInvalidDeadline = errors.New("Invalid deadline")
)

type List struct {
	ID          int
	Title       string
	Description string
	Tasks       map[int]*Task
	NextID      int
}

func (l *List) String() string {
	desc := l.Description
	if desc == "" {
		desc = "none"
	}

	return fmt.Sprintf("ID: %d\nTitle: %s\nDescription: %s\nTasks Count: %d\n\n",
		l.ID, l.Title, desc, len(l.Tasks))
}
