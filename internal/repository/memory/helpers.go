package memory

import (
	"github.com/Xiof22/ToDoList/internal/models"
	"slices"
)

func sortTasksByCreationTime(tasks []models.Task) []models.Task {
	slices.SortFunc(tasks, func(a, b models.Task) int {
		if a.CreatedAt.Before(b.CreatedAt) {
			return -1
		}

		return 1
	})

	return tasks
}
