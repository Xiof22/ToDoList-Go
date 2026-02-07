package mysql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
)

func (repo *Repository) CreateList(ctx context.Context, l models.List) (models.List, error) {
	q := `INSERT INTO lists (id, owner_id, title, description) VALUES (?, ?, ?, ?)`

	if _, err := repo.db.ExecContext(ctx, q, l.ID, l.OwnerID, l.Title, l.Description); err != nil {
		return models.List{}, errorsx.ErrExecDB
	}

	return l, nil
}

func (repo *Repository) GetLists(ctx context.Context) ([]models.List, error) {
	q := `SELECT id, owner_id, title, description FROM lists ORDER BY owner_id`

	rows, err := repo.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []models.List
	for rows.Next() {
		l := models.List{
			Tasks: make(map[models.TaskID]*models.Task),
		}

		if err := rows.Scan(
			&l.ID,
			&l.OwnerID,
			&l.Title,
			&l.Description,
		); err != nil {
			return nil, errorsx.ErrQueryDB
		}

		q := `
		    SELECT id, title, description, is_completed, deadline, created_at, updated_at
		    FROM tasks
		    WHERE list_id=?
		    ORDER BY id
		`

		rows, err := repo.db.QueryContext(ctx, q, l.ID)
		if err != nil {
			return nil, errorsx.ErrQueryDB
		}
		defer rows.Close()

		for rows.Next() {
			var t models.Task
			if err := rows.Scan(
				&t.ID,
				&t.Title,
				&t.Description,
				&t.IsCompleted,
				&t.Deadline,
				&t.CreatedAt,
				&t.UpdatedAt,
			); err != nil {
				return nil, errorsx.ErrQueryDB
			}

			l.Tasks[t.ID] = &t
		}
		if err := rows.Err(); err != nil {
			return nil, errorsx.ErrQueryDB
		}

		lists = append(lists, l)
	}
	if err := rows.Err(); err != nil {
		return nil, errorsx.ErrQueryDB
	}

	return lists, nil
}

func (repo *Repository) GetListsByUserID(ctx context.Context, userID models.UserID) ([]models.List, error) {
	q := `SELECT id, owner_id, title, description FROM lists WHERE owner_id=? ORDER BY id`

	rows, err := repo.db.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, errorsx.ErrQueryDB
	}
	defer rows.Close()

	var lists []models.List
	for rows.Next() {
		l := models.List{
			Tasks: make(map[models.TaskID]*models.Task),
		}

		if err := rows.Scan(
			&l.ID,
			&l.OwnerID,
			&l.Title,
			&l.Description,
		); err != nil {
			return nil, errorsx.ErrQueryDB
		}

		q := `
		    SELECT id, title, description, is_completed, deadline, created_at, updated_at
		    FROM tasks
		    WHERE list_id=?
		    ORDER BY id
		`

		rows, err := repo.db.QueryContext(ctx, q, l.ID)
		if err != nil {
			return nil, errorsx.ErrQueryDB
		}
		defer rows.Close()

		for rows.Next() {
			var t models.Task
			if err := rows.Scan(
				&t.ID,
				&t.Title,
				&t.Description,
				&t.IsCompleted,
				&t.Deadline,
				&t.CreatedAt,
				&t.UpdatedAt,
			); err != nil {
				return nil, errorsx.ErrQueryDB
			}

			l.Tasks[t.ID] = &t
		}
		if err := rows.Err(); err != nil {
			return nil, errorsx.ErrQueryDB
		}

		lists = append(lists, l)
	}
	if err := rows.Err(); err != nil {
		return nil, errorsx.ErrQueryDB
	}

	return lists, nil
}

func (repo *Repository) GetList(ctx context.Context, listID models.ListID) (models.List, error) {
	q := `SELECT id, owner_id, title, description FROM lists WHERE id=?`

	l := models.List{
		Tasks: make(map[models.TaskID]*models.Task),
	}

	if err := repo.db.QueryRowContext(ctx, q, listID).Scan(
		&l.ID,
		&l.OwnerID,
		&l.Title,
		&l.Description,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.List{}, errorsx.ErrListNotFound
		}

		return models.List{}, errorsx.ErrQueryDB
	}

	q = `
	    SELECT id, title, description, is_completed, deadline, created_at, updated_at
	    FROM tasks
	    WHERE list_id=?
	    ORDER BY id
	`

	rows, err := repo.db.QueryContext(ctx, q, l.ID)
	if err != nil {
		return models.List{}, errorsx.ErrQueryDB
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Task
		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.IsCompleted,
			&t.Deadline,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			return models.List{}, errorsx.ErrQueryDB
		}

		l.Tasks[t.ID] = &t
	}
	if err := rows.Err(); err != nil {
		return models.List{}, errorsx.ErrQueryDB
	}

	return l, nil
}

func (repo *Repository) EditList(ctx context.Context, listID models.ListID, list models.List) (models.List, error) {
	q := `UPDATE lists SET title=?, description=? WHERE id=?`

	result, err := repo.db.ExecContext(ctx, q, list.Title, list.Description, listID)
	if err != nil {
		return models.List{}, errorsx.ErrExecDB
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return models.List{}, errorsx.ErrExecDB
	} else if affected == 0 {
		return models.List{}, errorsx.ErrListNotFound
	}

	return repo.GetList(ctx, listID)
}

func (repo *Repository) DeleteList(ctx context.Context, listID models.ListID) error {
	q := `DELETE FROM lists WHERE id=?`

	result, err := repo.db.ExecContext(ctx, q, listID)
	if err != nil {
		return errorsx.ErrExecDB
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return errorsx.ErrExecDB
	} else if affected == 0 {
		return errorsx.ErrListNotFound
	}

	return nil
}
