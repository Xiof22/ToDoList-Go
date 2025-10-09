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
	assert.Equal(t, wantTask.ID, gotTask.ID)
	assert.Equal(t, wantTask.Title, gotTask.Title)
	assert.Equal(t, wantTask.Description, gotTask.Description)

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
			taskID := models.TaskID(uuid.Nil)

			repo.On("GetTask", t.Context(), taskID).Return(models.Task{}, tt.mockErr)

			task, err := svc.EditTask(t.Context(), taskID, dto.EditTaskRequest{})
			assert.Equal(t, models.Task{}, task)
			assert.ErrorIs(t, err, tt.wantErr)

			repo.AssertExpectations(t)
		})
	}
}
