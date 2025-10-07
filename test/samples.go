package test

import "github.com/Xiof22/ToDoList/internal/dto"

var sampleTask dto.Task = dto.Task{
	Title:       "Sample task title",
	Description: "Sample task description",
}

var sampleTaskMap map[string]any = map[string]any{
	"title":       sampleTask.Title,
	"description": sampleTask.Description,
}
