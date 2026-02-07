package models

import (
	"database/sql/driver"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserID uuid.UUID

func (id UserID) String() string {
	return uuid.UUID(id).String()
}

func (id UserID) Value() (driver.Value, error) {
	return id.String(), nil
}

func (id *UserID) Scan(value any) error {
	if value == nil {
		return nil
	}

	raw, ok := value.([]byte)
	if !ok {
		return errorsx.ErrInvalidUserID
	}

	parsed, err := uuid.Parse(string(raw))
	if err != nil {
		return errorsx.ErrInvalidUserID
	}

	*id = UserID(parsed)
	return nil
}

type Role int

const (
	Admin Role = iota
	Guest
)

type User struct {
	ID           UserID
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
		ID:           UserID(uuid.New()),
		Email:        email,
		PasswordHash: passwordHash,
		Role:         Guest,
	}, nil
}

type UserInfo struct {
	ID   UserID
	Role Role
}

func (u User) Info() UserInfo {
	return UserInfo{
		ID:   u.ID,
		Role: u.Role,
	}
}
