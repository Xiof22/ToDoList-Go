package dto

type EditListRequest struct {
	ListID      int    `json:"-" validate:"gt=0"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}
