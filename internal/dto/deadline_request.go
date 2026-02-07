package dto

import (
	"encoding/json"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"time"
)

type DeadlineRequest struct {
	Value time.Time
}

func (d *DeadlineRequest) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return errorsx.ErrUnmarshalDeadline
	}

	if s == "" || s == "null" {
		d.Value = time.Time{}
		return nil
	}

	parsed, err := time.Parse(time.DateTime, s)
	if err != nil {
		return errorsx.ErrInvalidDeadlineFormat
	}

	d.Value = parsed
	return nil
}
