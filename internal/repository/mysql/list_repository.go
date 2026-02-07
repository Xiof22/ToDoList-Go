package mysql

import (
	"context"
	"database/sql"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
	"github.com/google/uuid"
)

func (repo *Repository) CreateList(ctx context.Context, l models.List) (models.List, error) {
	q := `INSERT INTO lists (id, owner_id, title, description) VALUES (?, ?, ?, ?)`

	if _, err := repo.db.ExecContext(ctx, q, l.ID, l.OwnerID, l.Title, l.Description); err != nil {
		return models.List{}, errorsx.ErrExecDB
	}

	return l, nil
}

func (repo *Repository) GetLists(ctx context.Context) ([]models.List, error) {
	q := `
		SELECT
		    l.id, l.owner_id, l.title, l.description,
		    t.id, t.title, t.description, t.is_completed, t.deadline, t.created_at, t.updated_at
		FROM lists l
		LEFT JOIN tasks t ON l.id = t.list_id
		ORDER BY l.id
	`

	rows, err := repo.db.QueryContext(ctx, q)
	if err != nil {
		return nil, errorsx.ErrQueryDB
	}
	defer rows.Close()

	listsMap := make(map[models.ListID]*models.List)
	var listIDs []models.ListID
	for rows.Next() {
		var (
			listID          models.ListID
			listOwnerID     models.UserID
			listTitle       string
			listDescription string

			taskID          uuid.NullUUID
			taskTitle       sql.NullString
			taskDescription sql.NullString
			taskCompleted   sql.NullBool
			taskDeadline    sql.NullTime
			taskCreatedAt   sql.NullTime
			taskUpdatedAt   sql.NullTime
		)

		if err := rows.Scan(
			&listID,
			&listOwnerID,
			&listTitle,
			&listDescription,

			&taskID,
			&taskTitle,
			&taskDescription,
			&taskCompleted,
			&taskDeadline,
			&taskCreatedAt,
			&taskUpdatedAt,
		); err != nil {
			return nil, errorsx.ErrQueryDB
		}

		list, exists := listsMap[listID]
		if !exists {
			list = &models.List{
				ID:          listID,
				OwnerID:     listOwnerID,
				Title:       listTitle,
				Description: listDescription,
				Tasks:       make(map[models.TaskID]*models.Task),
			}

			listsMap[listID] = list
			listIDs = append(listIDs, listID)
		}

		if taskID.Valid {
			t := models.Task{
				ID:          models.TaskID(taskID.UUID),
				Title:       taskTitle.String,
				Description: taskDescription.String,
				IsCompleted: taskCompleted.Bool,
				Deadline:    taskDeadline.Time,
				CreatedAt:   taskCreatedAt.Time,
			}

			if taskUpdatedAt.Valid {
				t.UpdatedAt = &taskUpdatedAt.Time
			}

			list.Tasks[t.ID] = &t
		}
	}
	if err := rows.Err(); err != nil {
		return nil, errorsx.ErrQueryDB
	}

	lists := make([]models.List, 0, len(listIDs))

	for _, id := range listIDs {
		lists = append(lists, *listsMap[id])
	}

	return lists, nil
}

func (repo *Repository) GetListsByUserID(ctx context.Context, userID models.UserID) ([]models.List, error) {
	q := `
		SELECT
		    l.id, l.owner_id, l.title, l.description,
		    t.id, t.title, t.description, t.is_completed, t.deadline, t.created_at, t.updated_at
		FROM lists l
		LEFT JOIN tasks t ON l.id = t.list_id
		WHERE l.owner_id=?
		ORDER BY l.id
	`

	rows, err := repo.db.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, errorsx.ErrQueryDB
	}
	defer rows.Close()

	listsMap := make(map[models.ListID]*models.List)
	var listIDs []models.ListID
	for rows.Next() {
		var (
			listID          models.ListID
			listOwnerID     models.UserID
			listTitle       string
			listDescription string

			taskID          uuid.NullUUID
			taskTitle       sql.NullString
			taskDescription sql.NullString
			taskCompleted   sql.NullBool
			taskDeadline    sql.NullTime
			taskCreatedAt   sql.NullTime
			taskUpdatedAt   sql.NullTime
		)

		if err := rows.Scan(
			&listID,
			&listOwnerID,
			&listTitle,
			&listDescription,

			&taskID,
			&taskTitle,
			&taskDescription,
			&taskCompleted,
			&taskDeadline,
			&taskCreatedAt,
			&taskUpdatedAt,
		); err != nil {
			return nil, errorsx.ErrQueryDB
		}

		list, exists := listsMap[listID]
		if !exists {
			list = &models.List{
				ID:          listID,
				OwnerID:     listOwnerID,
				Title:       listTitle,
				Description: listDescription,
				Tasks:       make(map[models.TaskID]*models.Task),
			}

			listsMap[listID] = list
			listIDs = append(listIDs, listID)
		}

		if taskID.Valid {
			t := models.Task{
				ID:          models.TaskID(taskID.UUID),
				Title:       taskTitle.String,
				Description: taskDescription.String,
				IsCompleted: taskCompleted.Bool,
				Deadline:    taskDeadline.Time,
				CreatedAt:   taskCreatedAt.Time,
			}

			if taskUpdatedAt.Valid {
				t.UpdatedAt = &taskUpdatedAt.Time
			}

			list.Tasks[t.ID] = &t
		}
	}
	if err := rows.Err(); err != nil {
		return nil, errorsx.ErrQueryDB
	}

	lists := make([]models.List, 0, len(listIDs))

	for _, id := range listIDs {
		lists = append(lists, *listsMap[id])
	}

	return lists, nil
}

func (repo *Repository) GetList(ctx context.Context, listID models.ListID) (models.List, error) {
	q := `
		SELECT
		    l.id, l.owner_id, l.title, l.description,
		    t.id, t.title, t.description, t.is_completed, t.deadline, t.created_at, t.updated_at
		FROM lists l
		LEFT JOIN tasks t ON l.id = t.list_id
		WHERE l.id=?
	`

	l := models.List{
		Tasks: make(map[models.TaskID]*models.Task),
	}

	rows, err := repo.db.QueryContext(ctx, q, listID)
	if err != nil {
		return models.List{}, errorsx.ErrQueryDB
	}
	defer rows.Close()

	listFound := false

	for rows.Next() {
		listFound = true

		var (
			taskID          uuid.NullUUID
			taskTitle       sql.NullString
			taskDescription sql.NullString
			taskCompleted   sql.NullBool
			taskDeadline    sql.NullTime
			taskCreatedAt   sql.NullTime
			taskUpdatedAt   sql.NullTime
		)

		if err := rows.Scan(
			&l.ID,
			&l.OwnerID,
			&l.Title,
			&l.Description,

			&taskID,
			&taskTitle,
			&taskDescription,
			&taskCompleted,
			&taskDeadline,
			&taskCreatedAt,
			&taskUpdatedAt,
		); err != nil {
			return models.List{}, errorsx.ErrQueryDB
		}

		if taskID.Valid {
			t := models.Task{
				ID:          models.TaskID(taskID.UUID),
				Title:       taskTitle.String,
				Description: taskDescription.String,
				IsCompleted: taskCompleted.Bool,
				Deadline:    taskDeadline.Time,
				CreatedAt:   taskCreatedAt.Time,
			}

			if taskUpdatedAt.Valid {
				t.UpdatedAt = &taskUpdatedAt.Time
			}

			l.Tasks[t.ID] = &t
		}
	}
	if err := rows.Err(); err != nil {
		return models.List{}, errorsx.ErrQueryDB
	}

	if !listFound {
		return models.List{}, errorsx.ErrListNotFound
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
