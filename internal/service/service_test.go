package service

import (
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
	"github.com/Xiof22/ToDoList/internal/repository/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestEditTask_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	oldTask := models.NewTask("Test Title", "Test Description")
	req := dto.EditTaskRequest{
		Title:       "Edited " + oldTask.Title,
		Description: "Edited " + oldTask.Description,
	}

	wantTask := models.Task{
		ID:          oldTask.ID,
		Title:       req.Title,
		Description: req.Description,
	}

	repo.On("GetTask", mock.Anything, oldTask.ID).Return(oldTask, nil)
	repo.On("EditTask", mock.Anything, oldTask.ID, wantTask).Return(wantTask, nil)

	gotTask, err := svc.EditTask(t.Context(), oldTask.ID, req)
	assert.NoError(t, err)
	assert.Equal(t, wantTask, gotTask)

	repo.AssertExpectations(t)
}

func TestEditTask_RepoErrors(t *testing.T) {
	tests := []struct {
		name    string
		mockErr error
		wantErr error
	}{
		{
			name:    "Task not found",
			mockErr: errorsx.ErrTaskNotFound,
			wantErr: errorsx.ErrTaskNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			svc := New(repo)
			taskID := models.TaskID(uuid.New())

			repo.On("GetTask", mock.Anything, taskID).Return(models.Task{}, tt.mockErr)

			task, err := svc.EditTask(t.Context(), taskID, dto.EditTaskRequest{})
			assert.Equal(t, models.Task{}, task)
			assert.ErrorIs(t, err, tt.wantErr)

			repo.AssertExpectations(t)
		})
	}
}

func TestCompleteTask_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)
	task := models.NewTask("Test Title", "Test Description")
	completedTask := task
	completedTask.IsCompleted = true

	repo.On("GetTask", mock.Anything, task.ID).Return(task, nil)
	repo.On("EditTask", mock.Anything, task.ID, completedTask).Return(completedTask, nil)

	err := svc.CompleteTask(t.Context(), task.ID)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestCompleteTask_RepoErrors(t *testing.T) {
	tests := []struct {
		name    string
		mockErr error
		wantErr error
	}{
		{
			name:    "Task not found",
			mockErr: errorsx.ErrTaskNotFound,
			wantErr: errorsx.ErrTaskNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			svc := New(repo)
			taskID := models.TaskID(uuid.New())

			repo.On("GetTask", mock.Anything, taskID).Return(models.Task{}, tt.mockErr)

			err := svc.CompleteTask(t.Context(), taskID)
			assert.ErrorIs(t, err, tt.wantErr)

			repo.AssertExpectations(t)
		})
	}
}

func TestCompleteTask_AlreadyCompleted(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)
	task := models.NewTask("Test Title", "Test Description")
	task.IsCompleted = true

	repo.On("GetTask", mock.Anything, task.ID).Return(task, nil)

	err := svc.CompleteTask(t.Context(), task.ID)
	assert.ErrorIs(t, err, errorsx.ErrAlreadyCompleted)

	repo.AssertExpectations(t)
}

func TestUncompleteTask_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)
	task := models.NewTask("Test Title", "Test Description")
	completedTask := task
	completedTask.IsCompleted = true

	repo.On("GetTask", mock.Anything, completedTask.ID).Return(completedTask, nil)
	repo.On("EditTask", mock.Anything, completedTask.ID, task).Return(task, nil)

	err := svc.UncompleteTask(t.Context(), completedTask.ID)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestUncompleteTask_RepoErrors(t *testing.T) {
	tests := []struct {
		name    string
		mockErr error
		wantErr error
	}{
		{
			name:    "Task not found",
			mockErr: errorsx.ErrTaskNotFound,
			wantErr: errorsx.ErrTaskNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			svc := New(repo)
			taskID := models.TaskID(uuid.New())

			repo.On("GetTask", mock.Anything, taskID).Return(models.Task{}, tt.mockErr)

			err := svc.CompleteTask(t.Context(), taskID)
			assert.ErrorIs(t, err, tt.wantErr)

			repo.AssertExpectations(t)
		})
	}
}

func TestUncompleteTask_AlreadyUncompleted(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)
	task := models.NewTask("Test Title", "Test Description")

	repo.On("GetTask", mock.Anything, task.ID).Return(task, nil)

	err := svc.UncompleteTask(t.Context(), task.ID)
	assert.ErrorIs(t, err, errorsx.ErrAlreadyUncompleted)
}

func TestDeleteTask_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)
	task := models.NewTask("Test Title", "Test Description")

	repo.On("GetTask", mock.Anything, task.ID).Return(task, nil)
	repo.On("DeleteTask", mock.Anything, task.ID).Return(nil)

	err := svc.DeleteTask(t.Context(), task.ID)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestDeleteTask_RepoErrors(t *testing.T) {
	tests := []struct {
		name    string
		mockErr error
		wantErr error
	}{
		{
			name:    "Task not found",
			mockErr: errorsx.ErrTaskNotFound,
			wantErr: errorsx.ErrTaskNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			svc := New(repo)
			taskID := models.TaskID(uuid.New())

			repo.On("GetTask", mock.Anything, taskID).Return(models.Task{}, tt.mockErr)

			err := svc.DeleteTask(t.Context(), taskID)
			assert.ErrorIs(t, err, tt.wantErr)

			repo.AssertExpectations(t)
		})
	}
}
