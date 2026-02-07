package errorsx

import "errors"

var (
	ErrQueryDB = errors.New("Failed to fetch data from DB")
	ErrExecDB  = errors.New("Failed to execute DB-operation")
)
