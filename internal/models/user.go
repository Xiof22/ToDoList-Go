package models

import (
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"golang.org/x/crypto/bcrypt"
)

type Role int

const (
	Admin Role = iota
	Guest
)

type User struct {
	ID           int
	Email        string
	PasswordHash []byte
	Role         Role
}

func NewUser(email string, password string) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, errorsx.ErrHashPassword
	}

	return User{
		Email:        email,
		PasswordHash: passwordHash,
		Role:         Guest,
	}, nil
}

type UserInfo struct {
	ID   int
	Role Role
}

func (u User) Info() UserInfo {
	return UserInfo{
		ID:   u.ID,
		Role: u.Role,
	}
}
