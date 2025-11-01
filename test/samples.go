package test

import (
	"github.com/Xiof22/ToDoList/internal/dto"
	"time"
)

var sampleList dto.List = dto.List{
	Title:       "Sample list title",
	Description: "Sample list description",
}

var sampleListMap map[string]any = map[string]any{
	"title":       sampleList.Title,
	"description": sampleList.Description,
}

var sampleDeadline = time.Date(
	2030, time.December, 7,
	16, 59, 21, 0,
	time.UTC,
)

var sampleTask dto.Task = dto.Task{
	Title:       "Sample task title",
	Description: "Sample task description",
	Deadline:    &sampleDeadline,
}

var sampleTaskMap map[string]any = map[string]any{
	"title":       sampleTask.Title,
	"description": sampleTask.Description,
	"deadline":    sampleDeadline.Format(time.DateTime),
}
