package models

import "github.com/google/uuid"

type ListID uuid.UUID

func (id ListID) String() string {
	return uuid.UUID(id).String()
}

type List struct {
	ID          ListID
	OwnerID     UserID
	Title       string
	Description string
	Tasks       map[TaskID]*Task
}

func NewList(ownerID UserID, title, description string) List {
	return List{
		ID:          ListID(uuid.New()),
		OwnerID:     ownerID,
		Title:       title,
		Description: description,
		Tasks:       make(map[TaskID]*Task),
	}
}
