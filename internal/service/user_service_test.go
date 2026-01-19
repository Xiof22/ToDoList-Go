package service

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
	"github.com/Xiof22/ToDoList/internal/repository/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestRegister_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	req := dto.AuthRequest{
		Email:    "testUser@gmail.com",
		Password: "0000",
	}
	wantUser, err := models.NewUser(req.Email, req.Password)
	require.NoError(t, err)

	repo.On("GetUserByEmail", mock.Anything, req.Email).Return(models.User{}, errorsx.ErrUserNotFound).Once()
	repo.On("CreateUser", mock.Anything, mock.AnythingOfType("models.User")).Return(func(ctx context.Context, user models.User) (models.User, error) {
		return user, nil
	}).Once()

	gotUser, err := svc.Register(t.Context(), req)
	assert.NoError(t, err)
	assert.Equal(t, wantUser.Email, gotUser.Email)
	assert.Equal(t, wantUser.Role, gotUser.Role)

	err = bcrypt.CompareHashAndPassword(gotUser.PasswordHash, []byte(req.Password))
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestRegister_RepoErrors(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		email string,
	)

	tests := []struct {
		name    string
		setup   mockBehavior
		wantErr error
	}{
		{
			name: "Already registered",
			setup: func(
				repo *mocks.Repository,
				email string,
			) {
				repo.On("GetUserByEmail", mock.Anything, email).Return(models.User{}, nil).Once()
			},
			wantErr: errorsx.ErrEmailRegistered,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			svc := New(repo)

			req := dto.AuthRequest{
				Email:    "testUser@gmail.com",
				Password: "0000",
			}

			tt.setup(repo, req.Email)

			user, err := svc.Register(t.Context(), req)
			assert.Equal(t, models.User{}, user)
			assert.ErrorIs(t, err, tt.wantErr)

			repo.AssertExpectations(t)
		})
	}
}

func TestLogin_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	req := dto.AuthRequest{
		Email:    "testUser@gmail.com",
		Password: "0000",
	}
	wantUser, err := models.NewUser(req.Email, req.Password)
	require.NoError(t, err)

	repo.On("GetUserByEmail", mock.Anything, req.Email).Return(wantUser, nil).Once()

	gotUser, err := svc.Login(t.Context(), req)
	assert.NoError(t, err)
	assert.Equal(t, wantUser, gotUser)

	repo.AssertExpectations(t)
}

func TestLogin_RepoErrors(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		email string,
	)

	tests := []struct {
		name    string
		setup   mockBehavior
		wantErr error
	}{
		{
			name: "User not found",
			setup: func(
				repo *mocks.Repository,
				email string,
			) {
				repo.On("GetUserByEmail", mock.Anything, email).Return(models.User{}, errorsx.ErrUserNotFound).Once()
			},
			wantErr: errorsx.ErrInvalidCredentials,
		},
		{
			name: "Wrong password",
			setup: func(
				repo *mocks.Repository,
				email string,
			) {
				wrongHash, err := bcrypt.GenerateFromPassword([]byte("oops"), bcrypt.DefaultCost)
				require.NoError(t, err)

				repo.On("GetUserByEmail", mock.Anything, email).Return(models.User{
					PasswordHash: wrongHash,
				}, nil).Once()
			},
			wantErr: errorsx.ErrInvalidCredentials,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			svc := New(repo)

			req := dto.AuthRequest{
				Email:    "testUser@gmail.com",
				Password: "0000",
			}

			tt.setup(repo, req.Email)

			user, err := svc.Login(t.Context(), req)
			assert.Equal(t, models.User{}, user)
			assert.ErrorIs(t, err, tt.wantErr)

			repo.AssertExpectations(t)
		})
	}
}

func TestDeleteUser_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	info := models.UserInfo{
		ID:   models.UserID(uuid.New()),
		Role: models.Guest,
	}

	repo.On("GetUserByID", mock.Anything, info.ID).Return(models.User{}, nil).Once()
	repo.On("DeleteUser", mock.Anything, info.ID).Return(nil).Once()

	err := svc.DeleteUser(t.Context(), info)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestDeleteUser_RepoErrors(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		userID models.UserID,
	)

	tests := []struct {
		name    string
		setup   mockBehavior
		wantErr error
	}{
		{
			name: "User not found",
			setup: func(
				repo *mocks.Repository,
				userID models.UserID,
			) {
				repo.On("GetUserByID", mock.Anything, userID).Return(models.User{}, errorsx.ErrUserNotFound).Once()
			},
			wantErr: errorsx.ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			svc := New(repo)

			info := models.UserInfo{
				ID:   models.UserID(uuid.New()),
				Role: models.Guest,
			}

			tt.setup(repo, info.ID)

			err := svc.DeleteUser(t.Context(), info)
			assert.ErrorIs(t, err, tt.wantErr)

			repo.AssertExpectations(t)
		})
	}
}
