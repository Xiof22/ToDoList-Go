package memory

import (
	"cmp"
	"github.com/Xiof22/ToDoList/internal/models"
	"slices"
)

func sortTasksByID(tasks []models.Task) []models.Task {
	slices.SortFunc(tasks, func(a, b models.Task) int {
		return cmp.Compare(a.ID, b.ID)
	})

	return tasks
}

func sortListsByID(lists []models.List) []models.List {
	slices.SortFunc(lists, func(a, b models.List) int {
		return cmp.Compare(a.ID, b.ID)
	})

	return lists
}

func sortListsByOwnerID(lists []models.List) []models.List {
	slices.SortFunc(lists, func(a, b models.List) int {
		return cmp.Compare(a.OwnerID, b.OwnerID)
	})

	return lists
}
