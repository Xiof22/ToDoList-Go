package mysql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
)

func (repo *Repository) CreateUser(ctx context.Context, u models.User) (models.User, error) {
	q := `INSERT INTO users (id, email, password_hash, role) VALUES (?, ?, ?, ?)`

	if _, err := repo.db.ExecContext(ctx, q, u.ID, u.Email, u.PasswordHash, u.Role); err != nil {
		return models.User{}, errorsx.ErrExecDB
	}

	return u, nil
}

func (repo *Repository) GetUserByID(ctx context.Context, userID models.UserID) (models.User, error) {
	q := `
	    SELECT id, email, password_hash, role
	    FROM users
	    WHERE id=?
	`

	return repo.getUserByQuery(ctx, q, userID)
}

func (repo *Repository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	q := `
	    SELECT id, email, password_hash, role
	    FROM users
	    WHERE email=?
	`

	return repo.getUserByQuery(ctx, q, email)
}

func (repo *Repository) getUserByQuery(ctx context.Context, query string, value any) (models.User, error) {
	var u models.User
	if err := repo.db.QueryRowContext(ctx, query, value).Scan(
		&u.ID,
		&u.Email,
		&u.PasswordHash,
		&u.Role,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, errorsx.ErrUserNotFound
		}

		return models.User{}, errorsx.ErrQueryDB
	}

	return u, nil
}

func (repo *Repository) DeleteUser(ctx context.Context, userID models.UserID) error {
	q := `DELETE FROM users WHERE id=?`

	result, err := repo.db.ExecContext(ctx, q, userID)
	if err != nil {
		return errorsx.ErrExecDB
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return errorsx.ErrExecDB
	} else if affected == 0 {
		return errorsx.ErrUserNotFound
	}

	return nil
}
