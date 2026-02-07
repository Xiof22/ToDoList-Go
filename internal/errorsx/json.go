package errorsx

import "errors"

var (
	ErrWriteJSON   = errors.New("Failed to write JSON")
	ErrMissingJSON = errors.New("Missing JSON")
)
