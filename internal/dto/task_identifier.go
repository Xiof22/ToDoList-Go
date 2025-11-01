package dto

type TaskIdentifier struct {
	ListID int `validate:"gt=0"`
	TaskID int `validate:"gt=0"`
}
