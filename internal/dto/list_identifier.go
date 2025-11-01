package dto

type ListIdentifier struct {
	ID int `validate:"gt=0"`
}
