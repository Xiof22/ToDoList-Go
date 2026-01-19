package errorsx

import "errors"

var (
	ErrInvalidUserID = errors.New("Invalid user ID")
	ErrUserNotFound  = errors.New("User not found")
)
