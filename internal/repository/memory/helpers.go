package memory

import "github.com/Xiof22/ToDoList/internal/models"

func sortTasks(tasks []models.Task) []models.Task {
	if len(tasks) <= 1 {
		return tasks
	}

	for {
		swapped := false

		for current := 0; current < len(tasks)-1; current++ {
			next := current + 1

			if tasks[current].ID > tasks[next].ID {
				tasks[current], tasks[next] = tasks[next], tasks[current]
				swapped = true
			}
		}

		if !swapped {
			return tasks
		}
	}
}

func sortLists(lists []models.List) []models.List {
	if len(lists) <= 1 {
		return lists
	}

	for {
		swapped := false

		for current := 1; current < len(lists)-1; current++ {
			next := current + 1

			if lists[current].ID > lists[next].ID {
				lists[current], lists[next] = lists[next], lists[current]
				swapped = true
			}
		}

		if !swapped {
			return lists
		}
	}
}
