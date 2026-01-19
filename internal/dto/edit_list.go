package dto

type EditListRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}
