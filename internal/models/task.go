package models

import (
	"database/sql/driver"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/google/uuid"
	"time"
)

type TaskID uuid.UUID

func (id TaskID) String() string {
	return uuid.UUID(id).String()
}

func (id TaskID) Value() (driver.Value, error) {
	return id.String(), nil
}

func (id *TaskID) Scan(value any) error {
	if value == nil {
		return nil
	}

	raw, ok := value.([]byte)
	if !ok {
		return errorsx.ErrInvalidTaskID
	}

	parsed, err := uuid.Parse(string(raw))
	if err != nil {
		return errorsx.ErrInvalidTaskID
	}

	*id = TaskID(parsed)
	return nil
}

type Task struct {
	ID          TaskID
	Title       string
	Description string
	IsCompleted bool
	Deadline    time.Time
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func NewTask(title, description string, deadline time.Time) Task {
	return Task{
		ID:          TaskID(uuid.New()),
		Title:       title,
		Description: description,
		IsCompleted: false,
		Deadline:    deadline,
	}
}
