package dto

type EditTaskRequest struct {
	Title       string          `json:"title" validate:"required"`
	Description string          `json:"description"`
	Deadline    DeadlineRequest `json:"deadline"`
}
