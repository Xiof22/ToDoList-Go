package dto

type TaskIdentifier struct {
	ID int `validate:"gt=0"`
}
