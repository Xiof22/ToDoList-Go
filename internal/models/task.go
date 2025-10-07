package models

import (
	"fmt"
)

type Task struct {
	ID int
	Title string
	Description string
	IsCompleted bool
}

func (t Task) String() string {
	desc := t.Description
	if desc == "" {
		desc = "none"
	}

	return fmt.Sprintf("ID: %d\nTitle: %s\nDescription: %s\nCompleted: %t\n\n",
				t.ID, t.Title, desc, t.IsCompleted)
}
