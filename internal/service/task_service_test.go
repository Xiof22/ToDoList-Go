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

	list := models.NewList("Test List Title", "Test List Description")
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

	gotTask, err := svc.CreateTask(t.Context(), list.ID, req)
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

			listID := models.ListID(uuid.New())
			req := dto.CreateTaskRequest{
				Title:       "Test Title",
				Description: "Test Description",
				Deadline:    dto.DeadlineRequest{Value: time.Time{}},
			}

			tt.setup(repo, listID)

			task, err := svc.CreateTask(t.Context(), listID, req)
			assert.Equal(t, models.Task{}, task)
			assert.ErrorIs(t, err, tt.wantErr)

			repo.AssertExpectations(t)
		})
	}
}

func TestGetTasks_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	list := models.NewList("Test List Title", "Test List Description")
	var wantTasks []models.Task
	for i := 0; i < 3; i++ {
		task := models.NewTask(
			fmt.Sprintf("Test Title %d", i+1),
			fmt.Sprintf("Test Description %d", i+1),
			time.Time{},
		)

		wantTasks = append(wantTasks, task)
	}

	repo.On("GetList", mock.Anything, list.ID).Return(list, nil).Once()
	repo.On("GetTasks", mock.Anything, list.ID).Return(wantTasks, nil).Once()

	gotTasks, err := svc.GetTasks(t.Context(), list.ID)
	assert.NoError(t, err)
	assert.Equal(t, wantTasks, gotTasks)

	repo.AssertExpectations(t)
}

func TestGetTasks_RepoErrors(t *testing.T) {
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

			listID := models.ListID(uuid.New())

			tt.setup(repo, listID)

			tasks, err := svc.GetTasks(t.Context(), listID)
			assert.Nil(t, tasks)
			assert.ErrorIs(t, tt.wantErr, err)

			repo.AssertExpectations(t)
		})
	}
}

func TestEditTask_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	listID := models.ListID(uuid.New())
	oldTask := models.NewTask("Test Title", "Test Description", time.Time{})
	req := dto.EditTaskRequest{
		Title:       "Edited " + oldTask.Title,
		Description: "Edited " + oldTask.Description,
	}

	wantTask := models.Task{
		ID:          oldTask.ID,
		Title:       req.Title,
		Description: req.Description,
	}

	repo.On("GetList", mock.Anything, listID).Return(models.List{}, nil)
	repo.On("GetTask", mock.Anything, listID, oldTask.ID).Return(oldTask, nil)
	repo.On("EditTask", mock.Anything, listID, oldTask.ID, wantTask).Return(wantTask, nil)

	gotTask, err := svc.EditTask(t.Context(), listID, oldTask.ID, req)
	assert.NoError(t, err)
	assert.Equal(t, wantTask.ID, gotTask.ID)
	assert.Equal(t, wantTask.Title, gotTask.Title)
	assert.Equal(t, wantTask.Description, gotTask.Description)

	repo.AssertExpectations(t)
}

func TestEditTask_RepoErrors(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
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
				listID models.ListID,
				taskID models.TaskID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{}, nil).Once()
				repo.On("GetTask", mock.Anything, listID, taskID).Return(models.Task{}, errorsx.ErrTaskNotFound).Once()
			},
			wantErr: errorsx.ErrTaskNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			svc := New(repo)

			listID := models.ListID(uuid.New())
			taskID := models.TaskID(uuid.New())
			req := dto.EditTaskRequest{}

			tt.setup(repo, listID, taskID)

			task, err := svc.EditTask(t.Context(), listID, taskID, req)

			assert.Equal(t, models.Task{}, task)
			assert.ErrorIs(t, err, tt.wantErr)
			repo.AssertExpectations(t)
		})
	}
}

func TestEditTask_DeadlineBeforeCreation(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	listID := models.ListID(uuid.New())
	task := models.NewTask("Test Title", "Test Description", time.Time{})
	task.CreatedAt = time.Now()
	req := dto.EditTaskRequest{
		Title:       task.Title,
		Description: task.Description,
		Deadline: dto.DeadlineRequest{
			Value: task.CreatedAt.Add(-1 * time.Second),
		},
	}

	repo.On("GetList", mock.Anything, listID).Return(models.List{}, nil)
	repo.On("GetTask", mock.Anything, listID, task.ID).Return(task, nil)

	_, err := svc.EditTask(t.Context(), listID, task.ID, req)
	assert.ErrorIs(t, err, errorsx.ErrDeadlineBeforeCreation)

	repo.AssertExpectations(t)
}

func TestCompleteTask_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	listID := models.ListID(uuid.New())
	task := models.NewTask("Test Title", "Test Description", time.Time{})
	completedTask := task
	completedTask.IsCompleted = true

	repo.On("GetList", mock.Anything, listID).Return(models.List{}, nil)
	repo.On("GetTask", mock.Anything, listID, task.ID).Return(task, nil)
	repo.On("EditTask", mock.Anything, listID, task.ID, completedTask).Return(completedTask, nil)

	err := svc.CompleteTask(t.Context(), listID, task.ID)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestCompleteTask_RepoErrors(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
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
				listID models.ListID,
				taskID models.TaskID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{}, nil).Once()
				repo.On("GetTask", mock.Anything, listID, taskID).Return(models.Task{}, errorsx.ErrTaskNotFound).Once()
			},
			wantErr: errorsx.ErrTaskNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			svc := New(repo)

			listID := models.ListID(uuid.New())
			taskID := models.TaskID(uuid.New())

			tt.setup(repo, listID, taskID)

			err := svc.CompleteTask(t.Context(), listID, taskID)

			assert.ErrorIs(t, err, tt.wantErr)
			repo.AssertExpectations(t)
		})
	}
}

func TestCompleteTask_AlreadyCompleted(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	listID := models.ListID(uuid.New())
	task := models.NewTask("Test Title", "Test Description", time.Time{})
	task.IsCompleted = true

	repo.On("GetList", mock.Anything, listID).Return(models.List{}, nil)
	repo.On("GetTask", mock.Anything, listID, task.ID).Return(task, nil)

	err := svc.CompleteTask(t.Context(), listID, task.ID)
	assert.ErrorIs(t, err, errorsx.ErrAlreadyCompleted)

	repo.AssertExpectations(t)
}

func TestUncompleteTask_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	listID := models.ListID(uuid.New())
	task := models.NewTask("Test Title", "Test Description", time.Time{})
	completedTask := task
	completedTask.IsCompleted = true

	repo.On("GetList", mock.Anything, listID).Return(models.List{}, nil)
	repo.On("GetTask", mock.Anything, listID, completedTask.ID).Return(completedTask, nil)
	repo.On("EditTask", mock.Anything, listID, completedTask.ID, task).Return(task, nil)

	err := svc.UncompleteTask(t.Context(), listID, completedTask.ID)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestUncompleteTask_RepoErrors(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
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
				listID models.ListID,
				taskID models.TaskID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{}, nil).Once()
				repo.On("GetTask", mock.Anything, listID, taskID).Return(models.Task{}, errorsx.ErrTaskNotFound).Once()
			},
			wantErr: errorsx.ErrTaskNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			svc := New(repo)

			listID := models.ListID(uuid.New())
			taskID := models.TaskID(uuid.New())

			tt.setup(repo, listID, taskID)

			err := svc.UncompleteTask(t.Context(), listID, taskID)

			assert.ErrorIs(t, err, tt.wantErr)
			repo.AssertExpectations(t)
		})
	}
}

func TestUncompleteTask_AlreadyUncompleted(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	listID := models.ListID(uuid.New())
	task := models.NewTask("Test Title", "Test Description", time.Time{})

	repo.On("GetList", mock.Anything, listID).Return(models.List{}, nil)
	repo.On("GetTask", mock.Anything, listID, task.ID).Return(task, nil)

	err := svc.UncompleteTask(t.Context(), listID, task.ID)
	assert.ErrorIs(t, err, errorsx.ErrAlreadyUncompleted)

	repo.AssertExpectations(t)
}

func TestDeleteTask_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	listID := models.ListID(uuid.New())
	task := models.NewTask("Test Title", "Test Description", time.Time{})

	repo.On("GetList", mock.Anything, listID).Return(models.List{}, nil)
	repo.On("GetTask", mock.Anything, listID, task.ID).Return(task, nil)
	repo.On("DeleteTask", mock.Anything, listID, task.ID).Return(nil)

	err := svc.DeleteTask(t.Context(), listID, task.ID)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestDeleteTask_RepoErrors(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
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
				listID models.ListID,
				taskID models.TaskID,
			) {
				repo.On("GetList", mock.Anything, listID).Return(models.List{}, nil).Once()
				repo.On("GetTask", mock.Anything, listID, taskID).Return(models.Task{}, errorsx.ErrTaskNotFound).Once()
			},
			wantErr: errorsx.ErrTaskNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			svc := New(repo)

			listID := models.ListID(uuid.New())
			taskID := models.TaskID(uuid.New())

			tt.setup(repo, listID, taskID)

			err := svc.DeleteTask(t.Context(), listID, taskID)

			assert.ErrorIs(t, err, tt.wantErr)
			repo.AssertExpectations(t)
		})
	}
}
