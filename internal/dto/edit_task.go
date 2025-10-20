package dto

type EditTaskRequest struct {
	ID          int             `validate:"gt=0"`
	Title       string          `json:"title" validate:"required"`
	Description string          `json:"description"`
	Deadline    DeadlineRequest `json:"deadline"`
}
