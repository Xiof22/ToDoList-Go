package errorsx

import "errors"

var (
	ErrListNotFound  = errors.New("List not found")
	ErrInvalidListID = errors.New("Invalid list ID")
)
