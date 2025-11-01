package dto

type EditListRequest struct {
	ListID      int    `validate:"gt=0"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}
