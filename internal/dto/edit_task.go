package dto

type EditTaskRequest struct {
	ListID      int             `validate:"gt=0"`
	TaskID      int             `validate:"gt=0"`
	Title       string          `json:"title" validate:"required"`
	Description string          `json:"description"`
	Deadline    DeadlineRequest `json:"deadline"`
}
