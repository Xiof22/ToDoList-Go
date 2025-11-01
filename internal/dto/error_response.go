package dto

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

func ErrorToResponse(err error) (resp ErrorResponse) {
	ve, ok := err.(validator.ValidationErrors)
	if ok {
		resp.Errors = formatValidationErrors(ve)
		return
	}

	if err.Error() == "EOF" {
		resp.Errors = []string{"JSON decoding error"}
		return
	}

	resp.Errors = []string{err.Error()}
	return
}

func formatValidationErrors(ve validator.ValidationErrors) (messages []string) {
	for _, e := range ve {
		msg := fmt.Sprintf("Field '%s' doesn't match the rule '%s'", e.Field(), e.Tag())
		messages = append(messages, msg)
	}

	return
}
