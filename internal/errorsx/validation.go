package errorsx

import "fmt"

func ErrValidation(field, rule string) error {
	return fmt.Errorf("Field '%s' doesn't match the rule '%s'", field, rule)
}
