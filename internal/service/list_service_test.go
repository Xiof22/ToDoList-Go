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

func TestGetLists_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	repo.On("GetLists", mock.Anything).Return([]models.List{}, nil).Once()

	lists, err := svc.GetLists(t.Context())
	assert.NoError(t, err)
	assert.Equal(t, []models.List{}, lists)

	repo.AssertExpectations(t)
}

func TestGetList_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)
	wantList := models.NewList("Test List Title", "Test List Description")

	repo.On("GetList", mock.Anything, wantList.ID).Return(wantList, nil).Once()

	gotList, err := svc.GetList(t.Context(), wantList.ID)
	assert.NoError(t, err)
	assert.Equal(t, wantList, gotList)

	repo.AssertExpectations(t)
}

func TestGetList_RepoErrors(t *testing.T) {
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

			list, err := svc.GetList(t.Context(), listID)
			assert.Equal(t, models.List{}, list)
			assert.ErrorIs(t, err, tt.wantErr)

			repo.AssertExpectations(t)
		})
	}
}

func TestEditList_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)

	oldList := models.NewList("Test List Title", "Test List Description")
	req := dto.EditListRequest{
		Title:       "Edited " + oldList.Title,
		Description: "Edited " + oldList.Description,
	}

	wantList := models.List{
		ID:          oldList.ID,
		Title:       req.Title,
		Description: req.Description,
		Tasks:       oldList.Tasks,
	}

	repo.On("GetList", mock.Anything, oldList.ID).Return(oldList, nil)
	repo.On("EditList", mock.Anything, oldList.ID, wantList).Return(wantList, nil)

	gotList, err := svc.EditList(t.Context(), oldList.ID, req)
	assert.NoError(t, err)
	assert.Equal(t, wantList, gotList)

	repo.AssertExpectations(t)
}

func TestEditList_RepoErrors(t *testing.T) {
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

			gotList, err := svc.EditList(t.Context(), listID, dto.EditListRequest{})
			assert.Equal(t, models.List{}, gotList)
			assert.ErrorIs(t, err, tt.wantErr)

			repo.AssertExpectations(t)
		})
	}
}

func TestDeleteList_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	svc := New(repo)
	list := models.NewList("Test List Title", "Test List Description")

	repo.On("GetList", mock.Anything, list.ID).Return(list, nil)
	repo.On("DeleteList", mock.Anything, list.ID).Return(nil).Once()

	err := svc.DeleteList(t.Context(), list.ID)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestDeleteList_RepoErrors(t *testing.T) {
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

			err := svc.DeleteList(t.Context(), listID)
			assert.ErrorIs(t, err, tt.wantErr)

			repo.AssertExpectations(t)
		})
	}
}
