package errorsx

import "fmt"

func ErrParseContext(key string) error {
	return fmt.Errorf("Failed to parse '%s' from context", key)
}
