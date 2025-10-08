package test

import "github.com/Xiof22/ToDoList/internal/dto"

const (
	nilID     string = "00000000-0000-0000-0000-000000000000"
	invalidID        = "Invalid ID"
)

var (
	sampleTask dto.Task = dto.Task{
		Title:       "Sample task title",
		Description: "Sample task description",
	}

	sampleTaskMap map[string]any = map[string]any{
		"title":       sampleTask.Title,
		"description": sampleTask.Description,
	}
)
