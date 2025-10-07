package dto

type TasksResponse struct {
	Count int    `json:"tasks_count"`
	Tasks []Task `json:"tasks"`
}
