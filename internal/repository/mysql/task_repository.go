package mysql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
)

func (repo *Repository) CreateTask(ctx context.Context, listID models.ListID, t models.Task) (models.Task, error) {
	q := `INSERT INTO tasks (id, list_id, title, description, deadline) VALUES (?, ?, ?, ?, ?)`

	if _, err := repo.db.ExecContext(ctx, q, t.ID.String(), listID, t.Title, t.Description, t.Deadline); err != nil {
		return models.Task{}, errorsx.ErrExecDB
	}

	return t, nil
}

func (repo *Repository) GetTasks(ctx context.Context, listID models.ListID) ([]models.Task, error) {
	q := `
	    SELECT id, title, description, is_completed, deadline, created_at, updated_at
	    FROM tasks
	    WHERE list_id=?
	    ORDER BY id ASC
	`

	rows, err := repo.db.QueryContext(ctx, q, listID)
	if err != nil {
		return nil, errorsx.ErrQueryDB
	}
	defer rows.Close()

	var tt []models.Task
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

		tt = append(tt, t)
	}

	if err := rows.Err(); err != nil {
		return nil, errorsx.ErrQueryDB
	}

	return tt, nil
}

func (repo *Repository) GetTask(ctx context.Context, listID models.ListID, taskID models.TaskID) (models.Task, error) {
	q := `
	    SELECT id, title, description, is_completed, deadline, created_at, updated_at
	    FROM tasks
	    WHERE list_id=?
	      AND id = ?
	`

	var t models.Task
	if err := repo.db.QueryRowContext(ctx, q, listID, taskID).Scan(
		&t.ID,
		&t.Title,
		&t.Description,
		&t.IsCompleted,
		&t.Deadline,
		&t.CreatedAt,
		&t.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Task{}, errorsx.ErrTaskNotFound
		}

		return models.Task{}, err
	}

	return t, nil
}

func (repo *Repository) EditTask(ctx context.Context, listID models.ListID, taskID models.TaskID, t models.Task) (models.Task, error) {
	q := `
	    UPDATE tasks
	    SET title=?, description=?, is_completed=?, deadline=?
	    WHERE list_id=?
	      AND id=?
	`

	result, err := repo.db.ExecContext(
		ctx, q, t.Title, t.Description, t.IsCompleted, t.Deadline, listID, taskID,
	)
	if err != nil {
		return models.Task{}, errorsx.ErrExecDB
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return models.Task{}, errorsx.ErrExecDB
	} else if affected == 0 {
		return models.Task{}, errorsx.ErrTaskNotFound
	}

	return repo.GetTask(ctx, listID, taskID)
}

func (repo *Repository) DeleteTask(ctx context.Context, listID models.ListID, taskID models.TaskID) error {
	q := `DELETE FROM tasks WHERE list_id=? AND id=?`

	result, err := repo.db.ExecContext(ctx, q, listID, taskID)
	if err != nil {
		return errorsx.ErrExecDB
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return errorsx.ErrExecDB
	} else if affected == 0 {
		return errorsx.ErrTaskNotFound
	}

	return nil
}
