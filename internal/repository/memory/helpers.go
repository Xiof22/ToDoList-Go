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
