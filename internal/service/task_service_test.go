package service

import (
	"context"
	"fmt"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
	"github.com/Xiof22/ToDoList/internal/repository/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCreateTask_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	info := models.UserInfo{
		ID:   models.UserID(uuid.New()),
		Role: models.Guest,
	}
	list := models.NewList(info.ID, "Test List Title", "Test List Description")
	req := dto.CreateTaskRequest{
		Title:       "Test Title",
		Description: "Test Descriprion",
		Deadline:    dto.DeadlineRequest{Value: time.Time{}},
	}

	repo.On("GetList", mock.Anything, list.ID).Return(list, nil).Once()
	repo.On("CreateTask", mock.Anything, list.ID, mock.AnythingOfType("models.Task")).Return(
		func(ctx context.Context, listID models.ListID, task models.Task) (models.Task, error) {
			return task, nil
		},
	).Once()

	gotTask, err := svc.CreateTask(t.Context(), info, list.ID, req)
	assert.NoError(t, err)
	assert.Equal(t, req.Title, gotTask.Title)
	assert.Equal(t, req.Description, gotTask.Description)
	assert.Equal(t, req.Deadline.Value, gotTask.Deadline)

	repo.AssertExpectations(t)
}

func TestCreateTask_RepoErrors(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		listID models.ListID,
	)

	tests := []struct {
		name    string
		setup   mockBehavior
		wantErr error
	}{
		{
			name: "Content forbidden",
			setup: func(
				repo *mocks.Repository,
				listID models.ListID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{}, nil).Once()
			},
			wantErr: errorsx.ErrForbidden,
		},
		{
			name: "List not found",
			setup: func(
				repo *mocks.Repository,
				listID models.ListID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{}, errorsx.ErrListNotFound).Once()
			},
			wantErr: errorsx.ErrListNotFound,
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
			listID := models.ListID(uuid.New())
			req := dto.CreateTaskRequest{
				Title:       "Test Title",
				Description: "Test Description",
				Deadline:    dto.DeadlineRequest{Value: time.Time{}},
			}

			tt.setup(repo, listID)

			task, err := svc.CreateTask(t.Context(), info, listID, req)
			assert.Equal(t, models.Task{}, task)
			assert.ErrorIs(t, err, tt.wantErr)

			repo.AssertExpectations(t)
		})
	}
}

func TestGetTasks_Success(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		list *models.List,
	)

	tests := []struct {
		name  string
		setup mockBehavior
		info  models.UserInfo
	}{
		{
			name: "For admin",
			setup: func(
				repo *mocks.Repository,
				list *models.List,
			) {
				list.OwnerID = models.UserID(uuid.New())

				repo.On("GetList", mock.Anything, list.ID).Return(*list, nil).Once()
			},
			info: models.UserInfo{
				ID:   models.UserID(uuid.New()),
				Role: models.Admin,
			},
		},
		{
			name: "For guest",
			setup: func(
				repo *mocks.Repository,
				list *models.List,
			) {
				repo.On("GetList", mock.Anything, list.ID).Return(*list, nil).Once()
			},
			info: models.UserInfo{
				ID:   models.UserID(uuid.New()),
				Role: models.Guest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			svc := New(repo)

			list := models.NewList(tt.info.ID, "Test List Title", "Test List Description")
			var wantTasks []models.Task
			for i := 0; i < 3; i++ {
				task := models.NewTask(
					fmt.Sprintf("Test Title %d", i+1),
					fmt.Sprintf("Test Description %d", i+1),
					time.Time{},
				)

				wantTasks = append(wantTasks, task)
			}

			tt.setup(repo, &list)

			repo.On("GetTasks", mock.Anything, list.ID).Return(wantTasks, nil).Once()

			gotTasks, err := svc.GetTasks(t.Context(), tt.info, list.ID)
			assert.NoError(t, err)
			assert.Equal(t, wantTasks, gotTasks)

			repo.AssertExpectations(t)
		})
	}
}

func TestGetTasks_RepoErrors(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		info models.UserInfo,
		listID models.ListID,
	)

	tests := []struct {
		name    string
		setup   mockBehavior
		wantErr error
	}{
		{
			name: "Content forbidden",
			setup: func(
				repo *mocks.Repository,
				info models.UserInfo,
				listID models.ListID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{}, nil).Once()
			},
			wantErr: errorsx.ErrForbidden,
		},
		{
			name: "List not found",
			setup: func(
				repo *mocks.Repository,
				info models.UserInfo,
				listID models.ListID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{}, errorsx.ErrListNotFound).Once()
			},
			wantErr: errorsx.ErrListNotFound,
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
			listID := models.ListID(uuid.New())

			tt.setup(repo, info, listID)

			tasks, err := svc.GetTasks(t.Context(), info, listID)
			assert.Nil(t, tasks)
			assert.ErrorIs(t, tt.wantErr, err)

			repo.AssertExpectations(t)
		})
	}
}

func TestGetTask_Success(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		list *models.List,
	)

	tests := []struct {
		name  string
		setup mockBehavior
		info  models.UserInfo
	}{
		{
			name: "For admin",
			setup: func(
				repo *mocks.Repository,
				list *models.List,
			) {
				list.OwnerID = models.UserID(uuid.New())

				repo.On("GetList", mock.Anything, list.ID).Return(*list, nil).Once()
			},
			info: models.UserInfo{
				ID:   models.UserID(uuid.New()),
				Role: models.Admin,
			},
		},
		{
			name: "For guest",
			setup: func(
				repo *mocks.Repository,
				list *models.List,
			) {
				repo.On("GetList", mock.Anything, list.ID).Return(*list, nil).Once()
			},
			info: models.UserInfo{
				ID:   models.UserID(uuid.New()),
				Role: models.Guest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			svc := New(repo)

			list := models.NewList(tt.info.ID, "Test List Title", "Test List Description")
			wantTask := models.NewTask("Test Title", "Test Description", time.Time{})

			tt.setup(repo, &list)

			repo.On("GetTask", mock.Anything, list.ID, wantTask.ID).Return(wantTask, nil).Once()

			gotTask, err := svc.GetTask(t.Context(), tt.info, list.ID, wantTask.ID)
			assert.NoError(t, err)
			assert.Equal(t, wantTask, gotTask)

			repo.AssertExpectations(t)
		})
	}
}

func TestGetTask_RepoErrors(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		info models.UserInfo,
		listID models.ListID,
		taskID models.TaskID,
	)

	tests := []struct {
		name    string
		setup   mockBehavior
		wantErr error
	}{
		{
			name: "Content forbidden",
			setup: func(
				repo *mocks.Repository,
				info models.UserInfo,
				listID models.ListID,
				taskID models.TaskID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{}, nil).Once()
			},
			wantErr: errorsx.ErrForbidden,
		},
		{
			name: "List not found",
			setup: func(
				repo *mocks.Repository,
				info models.UserInfo,
				listID models.ListID,
				taskID models.TaskID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{}, errorsx.ErrListNotFound).Once()
			},
			wantErr: errorsx.ErrListNotFound,
		},
		{
			name: "Task not found",
			setup: func(
				repo *mocks.Repository,
				info models.UserInfo,
				listID models.ListID,
				taskID models.TaskID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{
					ID:          listID,
					OwnerID:     info.ID,
					Title:       "Test List Title",
					Description: "Test List Description",
				}, nil).Once()
				repo.On("GetTask", mock.Anything, listID, taskID).Return(models.Task{}, errorsx.ErrTaskNotFound).Once()
			},
			wantErr: errorsx.ErrTaskNotFound,
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
			listID := models.ListID(uuid.New())
			taskID := models.TaskID(uuid.New())

			tt.setup(repo, info, listID, taskID)

			task, err := svc.GetTask(t.Context(), info, listID, taskID)
			assert.Equal(t, models.Task{}, task)
			assert.ErrorIs(t, tt.wantErr, err)

			repo.AssertExpectations(t)
		})
	}
}

func TestEditTask_Success(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		list *models.List,
	)

	tests := []struct {
		name  string
		setup mockBehavior
		info  models.UserInfo
	}{
		{
			name: "For admin",
			setup: func(
				repo *mocks.Repository,
				list *models.List,
			) {
				list.OwnerID = models.UserID(uuid.New())

				repo.On("GetList", mock.Anything, list.ID).Return(*list, nil).Once()
			},
			info: models.UserInfo{
				ID:   models.UserID(uuid.New()),
				Role: models.Admin,
			},
		},
		{
			name: "For guest",
			setup: func(
				repo *mocks.Repository,
				list *models.List,
			) {
				repo.On("GetList", mock.Anything, list.ID).Return(*list, nil).Once()
			},
			info: models.UserInfo{
				ID:   models.UserID(uuid.New()),
				Role: models.Guest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			svc := New(repo)

			list := models.NewList(tt.info.ID, "Test List Title", "Test List Description")
			oldTask := models.NewTask("Test Title", "Test Description", time.Time{})
			req := dto.EditTaskRequest{
				Title:       "Edited " + oldTask.Title,
				Description: "Edited " + oldTask.Description,
				Deadline:    dto.DeadlineRequest{Value: time.Now().Add(1 * time.Hour)},
			}

			wantTask := models.Task{
				ID:          oldTask.ID,
				Title:       req.Title,
				Description: req.Description,
				Deadline:    req.Deadline.Value,
			}

			tt.setup(repo, &list)

			repo.On("GetTask", mock.Anything, list.ID, oldTask.ID).Return(oldTask, nil).Once()
			repo.On("EditTask", mock.Anything, list.ID, oldTask.ID, wantTask).Return(wantTask, nil).Once()

			gotTask, err := svc.EditTask(t.Context(), tt.info, list.ID, oldTask.ID, req)
			assert.NoError(t, err)
			assert.Equal(t, wantTask, gotTask)

			repo.AssertExpectations(t)
		})
	}
}

func TestEditTask_RepoErrors(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		info models.UserInfo,
		listID models.ListID,
		taskID models.TaskID,
	)

	tests := []struct {
		name    string
		setup   mockBehavior
		wantErr error
	}{
		{
			name: "Content forbidden",
			setup: func(
				repo *mocks.Repository,
				info models.UserInfo,
				listID models.ListID,
				taskID models.TaskID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{}, nil).Once()
			},
			wantErr: errorsx.ErrForbidden,
		},
		{
			name: "List not found",
			setup: func(
				repo *mocks.Repository,
				info models.UserInfo,
				listID models.ListID,
				taskID models.TaskID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{}, errorsx.ErrListNotFound).Once()
			},
			wantErr: errorsx.ErrListNotFound,
		},
		{
			name: "Task not found",
			setup: func(
				repo *mocks.Repository,
				info models.UserInfo,
				listID models.ListID,
				taskID models.TaskID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{
					ID:          listID,
					OwnerID:     info.ID,
					Title:       "Test List Title",
					Description: "Test List Description",
				}, nil).Once()
				repo.On("GetTask", mock.Anything, listID, taskID).Return(models.Task{}, errorsx.ErrTaskNotFound).Once()
			},
			wantErr: errorsx.ErrTaskNotFound,
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
			listID := models.ListID(uuid.New())
			taskID := models.TaskID(uuid.New())
			req := dto.EditTaskRequest{}

			tt.setup(repo, info, listID, taskID)

			task, err := svc.EditTask(t.Context(), info, listID, taskID, req)

			assert.Equal(t, models.Task{}, task)
			assert.ErrorIs(t, err, tt.wantErr)
			repo.AssertExpectations(t)
		})
	}
}

func TestEditTask_DeadlineBeforeCreation(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	info := models.UserInfo{
		ID:   models.UserID(uuid.New()),
		Role: models.Guest,
	}
	list := models.NewList(info.ID, "Test List Title", "Test List Description")
	task := models.NewTask("Test Title", "Test Description", time.Time{})
	task.CreatedAt = time.Now()
	req := dto.EditTaskRequest{
		Title:       task.Title,
		Description: task.Description,
		Deadline: dto.DeadlineRequest{
			Value: task.CreatedAt.Add(-1 * time.Second),
		},
	}

	repo.On("GetList", mock.Anything, list.ID).Return(list, nil).Once()
	repo.On("GetTask", mock.Anything, list.ID, task.ID).Return(task, nil).Once()

	_, err := svc.EditTask(t.Context(), info, list.ID, task.ID, req)
	assert.ErrorIs(t, err, errorsx.ErrDeadlineBeforeCreation)

	repo.AssertExpectations(t)
}

func TestCompleteTask_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	info := models.UserInfo{
		ID:   models.UserID(uuid.New()),
		Role: models.Guest,
	}
	list := models.NewList(info.ID, "Test List Title", "Test List Description")
	task := models.NewTask("Test Title", "Test Description", time.Time{})
	completedTask := task
	completedTask.IsCompleted = true

	repo.On("GetList", mock.Anything, list.ID).Return(list, nil).Once()
	repo.On("GetTask", mock.Anything, list.ID, task.ID).Return(task, nil).Once()
	repo.On("EditTask", mock.Anything, list.ID, task.ID, completedTask).Return(completedTask, nil).Once()

	err := svc.CompleteTask(t.Context(), info, list.ID, task.ID)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestCompleteTask_RepoErrors(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		info models.UserInfo,
		listID models.ListID,
		taskID models.TaskID,
	)

	tests := []struct {
		name    string
		setup   mockBehavior
		wantErr error
	}{
		{
			name: "List not found",
			setup: func(
				repo *mocks.Repository,
				info models.UserInfo,
				listID models.ListID,
				taskID models.TaskID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{}, errorsx.ErrListNotFound).Once()
			},
			wantErr: errorsx.ErrListNotFound,
		},
		{
			name: "Task not found",
			setup: func(
				repo *mocks.Repository,
				info models.UserInfo,
				listID models.ListID,
				taskID models.TaskID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{
					ID:          listID,
					OwnerID:     info.ID,
					Title:       "Test List Title",
					Description: "Test List Description",
				}, nil).Once()
				repo.On("GetTask", mock.Anything, listID, taskID).Return(models.Task{}, errorsx.ErrTaskNotFound).Once()
			},
			wantErr: errorsx.ErrTaskNotFound,
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
			listID := models.ListID(uuid.New())
			taskID := models.TaskID(uuid.New())

			tt.setup(repo, info, listID, taskID)

			err := svc.CompleteTask(t.Context(), info, listID, taskID)

			assert.ErrorIs(t, err, tt.wantErr)
			repo.AssertExpectations(t)
		})
	}
}

func TestCompleteTask_AlreadyCompleted(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	info := models.UserInfo{
		ID:   models.UserID(uuid.New()),
		Role: models.Guest,
	}
	list := models.NewList(info.ID, "Test List Title", "Test List Description")
	task := models.NewTask("Test Title", "Test Description", time.Time{})
	task.IsCompleted = true

	repo.On("GetList", mock.Anything, list.ID).Return(list, nil).Once()
	repo.On("GetTask", mock.Anything, list.ID, task.ID).Return(task, nil).Once()

	err := svc.CompleteTask(t.Context(), info, list.ID, task.ID)
	assert.ErrorIs(t, err, errorsx.ErrAlreadyCompleted)

	repo.AssertExpectations(t)
}

func TestUncompleteTask_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	info := models.UserInfo{
		ID:   models.UserID(uuid.New()),
		Role: models.Guest,
	}
	list := models.NewList(info.ID, "Test List Title", "Test List Description")
	task := models.NewTask("Test Title", "Test Description", time.Time{})
	completedTask := task
	completedTask.IsCompleted = true

	repo.On("GetList", mock.Anything, list.ID).Return(list, nil).Once()
	repo.On("GetTask", mock.Anything, list.ID, completedTask.ID).Return(completedTask, nil).Once()
	repo.On("EditTask", mock.Anything, list.ID, completedTask.ID, task).Return(task, nil).Once()

	err := svc.UncompleteTask(t.Context(), info, list.ID, completedTask.ID)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestUncompleteTask_RepoErrors(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		info models.UserInfo,
		listID models.ListID,
		taskID models.TaskID,
	)

	tests := []struct {
		name    string
		setup   mockBehavior
		wantErr error
	}{
		{
			name: "List not found",
			setup: func(
				repo *mocks.Repository,
				info models.UserInfo,
				listID models.ListID,
				taskID models.TaskID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{}, errorsx.ErrListNotFound).Once()
			},
			wantErr: errorsx.ErrListNotFound,
		},
		{
			name: "Task not found",
			setup: func(
				repo *mocks.Repository,
				info models.UserInfo,
				listID models.ListID,
				taskID models.TaskID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{
					ID:          listID,
					OwnerID:     info.ID,
					Title:       "Test List Title",
					Description: "Test List Description",
				}, nil).Once()
				repo.On("GetTask", mock.Anything, listID, taskID).Return(models.Task{}, errorsx.ErrTaskNotFound).Once()
			},
			wantErr: errorsx.ErrTaskNotFound,
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
			listID := models.ListID(uuid.New())
			taskID := models.TaskID(uuid.New())

			tt.setup(repo, info, listID, taskID)

			err := svc.UncompleteTask(t.Context(), info, listID, taskID)

			assert.ErrorIs(t, err, tt.wantErr)
			repo.AssertExpectations(t)
		})
	}
}

func TestUncompleteTask_AlreadyUncompleted(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	info := models.UserInfo{
		ID:   models.UserID(uuid.New()),
		Role: models.Guest,
	}
	list := models.NewList(info.ID, "Test List Title", "Test List Description")
	task := models.NewTask("Test Title", "Test Description", time.Time{})

	repo.On("GetList", mock.Anything, list.ID).Return(list, nil).Once()
	repo.On("GetTask", mock.Anything, list.ID, task.ID).Return(task, nil).Once()

	err := svc.UncompleteTask(t.Context(), info, list.ID, task.ID)
	assert.ErrorIs(t, err, errorsx.ErrAlreadyUncompleted)

	repo.AssertExpectations(t)
}

func TestDeleteTask_Success(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		list *models.List,
	)

	tests := []struct {
		name  string
		setup mockBehavior
		info  models.UserInfo
	}{
		{
			name: "For admin",
			setup: func(
				repo *mocks.Repository,
				list *models.List,
			) {
				list.OwnerID = models.UserID(uuid.New())

				repo.On("GetList", mock.Anything, list.ID).Return(*list, nil).Once()
			},
			info: models.UserInfo{
				ID:   models.UserID(uuid.New()),
				Role: models.Admin,
			},
		},
		{
			name: "For guest",
			setup: func(
				repo *mocks.Repository,
				list *models.List,
			) {
				repo.On("GetList", mock.Anything, list.ID).Return(*list, nil).Once()
			},
			info: models.UserInfo{
				ID:   models.UserID(uuid.New()),
				Role: models.Guest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			svc := New(repo)

			list := models.NewList(tt.info.ID, "Test List Title", "Test List Description")
			task := models.NewTask("Test Title", "Test Description", time.Time{})

			tt.setup(repo, &list)

			repo.On("GetTask", mock.Anything, list.ID, task.ID).Return(task, nil).Once()
			repo.On("DeleteTask", mock.Anything, list.ID, task.ID).Return(nil).Once()

			err := svc.DeleteTask(t.Context(), tt.info, list.ID, task.ID)
			assert.NoError(t, err)

			repo.AssertExpectations(t)
		})
	}
}

func TestDeleteTask_RepoErrors(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		info models.UserInfo,
		listID models.ListID,
		taskID models.TaskID,
	)

	tests := []struct {
		name    string
		setup   mockBehavior
		wantErr error
	}{
		{
			name: "List not found",
			setup: func(
				repo *mocks.Repository,
				info models.UserInfo,
				listID models.ListID,
				taskID models.TaskID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{}, errorsx.ErrListNotFound).Once()
			},
			wantErr: errorsx.ErrListNotFound,
		},
		{
			name: "Task not found",
			setup: func(
				repo *mocks.Repository,
				info models.UserInfo,
				listID models.ListID,
				taskID models.TaskID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{
					ID:          listID,
					OwnerID:     info.ID,
					Title:       "Test List Title",
					Description: "Test List Description",
				}, nil).Once()
				repo.On("GetTask", mock.Anything, listID, taskID).Return(models.Task{}, errorsx.ErrTaskNotFound).Once()
			},
			wantErr: errorsx.ErrTaskNotFound,
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
			listID := models.ListID(uuid.New())
			taskID := models.TaskID(uuid.New())

			tt.setup(repo, info, listID, taskID)

			err := svc.DeleteTask(t.Context(), info, listID, taskID)

			assert.ErrorIs(t, err, tt.wantErr)
			repo.AssertExpectations(t)
		})
	}
}
