package test

import (
	"github.com/Xiof22/ToDoList/internal/dto"
	"time"
)

func newUserMap(email, password string) map[string]any {
	return map[string]any{
		"email":    email,
		"password": password,
	}
}

const (
	nilID                 string = "00000000-0000-0000-0000-000000000000"
	pastDeadline                 = "2004-07-12 16:59:27"
	invalidID                    = "Invalid ID"
	invalidFormatDeadline        = "The 7-th of December 2030 year"
	invalidEmail                 = "Invalid email"
	tooShortPassword             = "0"
	tooLongPassword              = "Very very very very very very long password"
)

var (
	sampleList dto.List = dto.List{
		Title:       "Sample list title",
		Description: "Sample list description",
	}

	sampleListMap map[string]any = map[string]any{
		"title":       sampleList.Title,
		"description": sampleList.Description,
	}

	sampleDeadline = time.Date(
		2030, time.December, 7,
		16, 59, 21, 0,
		time.UTC,
	)

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
