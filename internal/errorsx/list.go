package errorsx

import "errors"

var (
	ErrInvalidListID = errors.New("Invalid list ID")
	ErrListNotFound  = errors.New("List not found")
)
