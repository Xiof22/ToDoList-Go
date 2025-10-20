package test

import (
	"github.com/Xiof22/ToDoList/internal/dto"
	"time"
)

var sampleDeadline = time.Date(
	2030, time.December, 7,
	16, 59, 21, 0,
	time.UTC,
)

const (
	nilID                 string = "00000000-0000-0000-0000-000000000000"
	invalidID                    = "Invalid ID"
	pastDeadline                 = "2004-07-12 16:59:27"
	invalidFormatDeadline        = "The 7-th of December 2030 year"
)

var (
	sampleTask dto.Task = dto.Task{
		Title:       "Sample task title",
		Description: "Sample task description",
		Deadline:    &sampleDeadline,
	}

	sampleTaskMap map[string]any = map[string]any{
		"title":       sampleTask.Title,
		"description": sampleTask.Description,
		"deadline":    sampleDeadline.Format(time.DateTime),
	}
)
