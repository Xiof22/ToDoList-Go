package service

import (
	"context"
	"errors"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (svc *Service) GetUserByID(ctx context.Context, userID models.UserID) (models.User, error) {
	return svc.repo.GetUserByID(ctx, userID)
}

func (svc *Service) Register(ctx context.Context, req dto.AuthRequest) (models.User, error) {
	if _, err := svc.repo.GetUserByEmail(ctx, req.Email); err == nil {
		return models.User{}, errorsx.ErrEmailRegistered
	} else if !errors.Is(err, errorsx.ErrUserNotFound) {
		return models.User{}, err
	}

	user, err := models.NewUser(req.Email, req.Password)
	if err != nil {
		return models.User{}, err
	}

	return svc.repo.CreateUser(ctx, user)
}

func (svc *Service) Login(ctx context.Context, req dto.AuthRequest) (models.User, error) {
	user, err := svc.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, errorsx.ErrUserNotFound) {
			err = errorsx.ErrInvalidCredentials
		}

		return models.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(req.Password)); err != nil {
		return models.User{}, errorsx.ErrInvalidCredentials
	}

	return user, nil
}

func (svc *Service) DeleteUser(ctx context.Context, info models.UserInfo) error {
	if _, err := svc.repo.GetUserByID(ctx, info.ID); err != nil {
		return errorsx.ErrUserNotFound
	}

	return svc.repo.DeleteUser(ctx, info.ID)
}
