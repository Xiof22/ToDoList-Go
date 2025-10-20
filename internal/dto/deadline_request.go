package dto

import (
	"encoding/json"
	"errors"
	"time"
)

type DeadlineRequest struct {
	Value time.Time
}

func (d *DeadlineRequest) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return errors.New("Deadline unmarshalling error")
	}

	if s == "" || s == "null" {
		d.Value = time.Time{}
		return nil
	}

	parsed, err := time.Parse(time.DateTime, s)
	if err != nil {
		return errors.New("Unexpected deadline format")
	}

	d.Value = parsed
	return nil
}
