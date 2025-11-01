package dto

type CreateTaskRequest struct {
	ListID      int             `validate:"gt=0"`
	Title       string          `json:"title" validate:"required"`
	Description string          `json:"description"`
	Deadline    DeadlineRequest `json:"deadline" validate:"future_or_empty"`
}
