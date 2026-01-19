package dto

import "github.com/Xiof22/ToDoList/internal/models"

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func ToUserDTO(user models.User) User {
	return User{
		ID:    user.ID,
		Email: user.Email,
	}
}
