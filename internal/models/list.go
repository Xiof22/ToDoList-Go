package models

import "github.com/google/uuid"

type ListID uuid.UUID

func (id ListID) String() string {
	return uuid.UUID(id).String()
}

type List struct {
	ID          ListID
	Title       string
	Description string
	Tasks       map[TaskID]*Task
}

func NewList(title, description string) List {
	return List{
		ID:          ListID(uuid.New()),
		Title:       title,
		Description: description,
		Tasks:       make(map[TaskID]*Task),
	}
}
