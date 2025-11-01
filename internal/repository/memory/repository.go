package memory

import (
	"errors"
	"github.com/Xiof22/ToDoList/internal/models"
	"sync"
)

var (
	ErrListNotFound       = errors.New("List not found")
	ErrTaskNotFound       = errors.New("Task not found")
	ErrAlreadyCompleted   = errors.New("Task is already completed")
	ErrAlreadyUncompleted = errors.New("Task is already uncompleted")
	ErrInvalidDeadline    = errors.New("Invalid deadline")
)

type Repository struct {
	mu     sync.Mutex
	Lists  map[int]*models.List
	nextID int
}

func New() *Repository {
	return &Repository{
		Lists:  make(map[int]*models.List),
		nextID: 1,
	}
}

func sortTasks(tasks []models.Task) []models.Task {
	if len(tasks) == 1 {
		return tasks
	}

	for {
		swapped := false

		for i := 0; i < len(tasks)-1; i++ {
			current := tasks[i]
			next := tasks[i+1]

			if current.ID > next.ID {
				tasks[i], tasks[i+1] = next, current
				swapped = true
			}
		}

		if !swapped {
			return tasks
		}
	}
}

func sortLists(lists []models.List) []models.List {
	if len(lists) == 1 {
		return lists
	}

	for {
		swapped := false

		for i := 0; i < len(lists)-1; i++ {
			current := lists[i]
			next := lists[i+1]

			if current.ID > next.ID {
				lists[i], lists[i+1] = next, current
				swapped = true
			}
		}

		if !swapped {
			return lists
		}
	}
}
