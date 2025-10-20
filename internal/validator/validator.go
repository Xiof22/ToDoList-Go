package validator

import (
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/go-playground/validator/v10"
	"time"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()

	Validate.RegisterValidation("future_or_empty", func(fl validator.FieldLevel) bool {
		d := fl.Field().Interface().(dto.DeadlineRequest)
		return d.Value.After(time.Now()) || d.Value.IsZero()
	})
}
