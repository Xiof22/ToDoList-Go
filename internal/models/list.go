package models

import (
	"database/sql/driver"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/google/uuid"
)

type ListID uuid.UUID

func (id ListID) String() string {
	return uuid.UUID(id).String()
}

func (id ListID) Value() (driver.Value, error) {
	return id.String(), nil
}

func (id *ListID) Scan(value any) error {
	if value == nil {
		return nil
	}

	raw, ok := value.([]byte)
	if !ok {
		return errorsx.ErrInvalidListID
	}

	parsed, err := uuid.Parse(string(raw))
	if err != nil {
		return errorsx.ErrInvalidListID
	}

	*id = ListID(parsed)
	return nil
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
