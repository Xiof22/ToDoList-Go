package dto

type CreateListRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}
