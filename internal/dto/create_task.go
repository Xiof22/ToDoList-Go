package dto

type CreateTaskRequest struct {
	Title       string          `json:"title" validate:"required"`
	Description string          `json:"description"`
	Deadline    DeadlineRequest `json:"deadline" validate:"future_or_empty"`
}
