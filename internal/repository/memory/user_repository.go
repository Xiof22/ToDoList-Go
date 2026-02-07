package memory

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
)

func (repo *Repository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.Users[user.ID] = &user

	return user, nil
}

func (repo *Repository) GetUserByID(ctx context.Context, userID models.UserID) (models.User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	u, ok := repo.Users[userID]
	if !ok {
		return models.User{}, errorsx.ErrUserNotFound
	}

	return *u, nil
}

func (repo *Repository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for _, u := range repo.Users {
		if u.Email == email {
			return *u, nil
		}
	}

	return models.User{}, errorsx.ErrUserNotFound
}

func (repo *Repository) DeleteUser(ctx context.Context, userID models.UserID) error {
	lists, err := repo.GetListsByUserID(ctx, userID)
	if err != nil {
		return err
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	for _, l := range lists {
		delete(repo.Lists, l.ID)
	}

	delete(repo.Users, userID)
	return nil
}
