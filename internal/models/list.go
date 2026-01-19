package models

type List struct {
	ID          int
	OwnerID     int
	Title       string
	Description string
	Tasks       map[int]*Task
	nextID      int
}

func NewList(ownerID int, title, description string) List {
	return List{
		OwnerID:     ownerID,
		Title:       title,
		Description: description,
		Tasks:       make(map[int]*Task),
		nextID:      1,
	}
}

func (l *List) AddTask(task Task) Task {
	task.ID = l.nextID
	l.Tasks[l.nextID] = &task
	l.nextID++

	return task
}
