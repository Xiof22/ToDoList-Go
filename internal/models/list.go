package models

type List struct {
	ID          int
	Title       string
	Description string
	Tasks       map[int]*Task
	NextID      int
}
