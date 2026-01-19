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
	type mockBehavior func(
		repo *mocks.Repository,
		info models.UserInfo,
	)

	tests := []struct {
		name  string
		info  models.UserInfo
		setup mockBehavior
	}{
		{
			name: "For Admin",
			info: models.UserInfo{
				ID:   models.UserID(uuid.New()),
				Role: models.Admin,
			},
			setup: func(
				repo *mocks.Repository,
				info models.UserInfo,
			) {
				repo.On("GetLists", mock.Anything).Return([]models.List{}, nil).Once()
			},
		},
		{
			name: "For Guest",
			info: models.UserInfo{
				ID:   models.UserID(uuid.New()),
				Role: models.Guest,
			},
			setup: func(
				repo *mocks.Repository,
				info models.UserInfo,
			) {
				repo.On("GetListsByUserID", mock.Anything, info.ID).Return([]models.List{}, nil).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			svc := New(repo)

			tt.setup(repo, tt.info)

			lists, err := svc.GetLists(t.Context(), tt.info)
			assert.NoError(t, err)
			assert.Equal(t, []models.List{}, lists)

			repo.AssertExpectations(t)
		})
	}
}

func TestGetList_Success(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		list models.List,
	)

	tests := []struct {
		name  string
		setup mockBehavior
		info  models.UserInfo
	}{
		{
			name: "For Admin",
			setup: func(
				repo *mocks.Repository,
				list models.List,
			) {
				list.OwnerID = models.UserID(uuid.New())

				repo.On("GetList", mock.Anything, list.ID).Return(list, nil).Once()
			},
			info: models.UserInfo{
				ID:   models.UserID(uuid.New()),
				Role: models.Admin,
			},
		},
		{
			name: "For Guest",
			setup: func(
				repo *mocks.Repository,
				list models.List,
			) {
				repo.On("GetList", mock.Anything, list.ID).Return(list, nil).Once()
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
			wantList := models.NewList(tt.info.ID, "Test List Title", "Test List Description")

			tt.setup(repo, wantList)

			gotList, err := svc.GetList(t.Context(), tt.info, wantList.ID)
			assert.NoError(t, err)
			assert.Equal(t, wantList.ID, gotList.ID)
			assert.Equal(t, wantList.Title, gotList.Title)
			assert.Equal(t, wantList.Description, gotList.Description)
			assert.Equal(t, wantList.Tasks, gotList.Tasks)

			if tt.info.Role == models.Guest {
				assert.Equal(t, wantList.OwnerID, gotList.OwnerID)
			}

			repo.AssertExpectations(t)
		})
	}
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

			tt.setup(repo, listID)

			list, err := svc.GetList(t.Context(), info, listID)
			assert.Equal(t, models.List{}, list)
			assert.ErrorIs(t, err, tt.wantErr)

			repo.AssertExpectations(t)
		})
	}
}

func TestEditList_Success(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		oldList models.List,
		wantList models.List,
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
				oldList models.List,
				wantList models.List,
			) {
				oldList.OwnerID = models.UserID(uuid.New())
				wantList.OwnerID = oldList.OwnerID

				repo.On("GetList", mock.Anything, oldList.ID).Return(oldList, nil)
				repo.On("EditList", mock.Anything, oldList.ID, wantList).Return(wantList, nil)
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
				oldList models.List,
				wantList models.List,
			) {
				repo.On("GetList", mock.Anything, oldList.ID).Return(oldList, nil)
				repo.On("EditList", mock.Anything, oldList.ID, wantList).Return(wantList, nil)
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

			oldList := models.NewList(tt.info.ID, "Test List Title", "Test List Description")
			req := dto.EditListRequest{
				Title:       "Edited " + oldList.Title,
				Description: "Edited " + oldList.Description,
			}

			wantList := models.List{
				OwnerID:     oldList.OwnerID,
				ID:          oldList.ID,
				Title:       req.Title,
				Description: req.Description,
				Tasks:       oldList.Tasks,
			}

			tt.setup(repo, oldList, wantList)

			gotList, err := svc.EditList(t.Context(), tt.info, oldList.ID, req)
			assert.NoError(t, err)
			assert.Equal(t, wantList.ID, gotList.ID)
			assert.Equal(t, wantList.Title, gotList.Title)
			assert.Equal(t, wantList.Description, gotList.Description)
			assert.Equal(t, wantList.Tasks, gotList.Tasks)

			if tt.info.Role == models.Guest {
				assert.Equal(t, wantList.OwnerID, gotList.OwnerID)
			}

			repo.AssertExpectations(t)
		})
	}
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

			tt.setup(repo, listID)

			gotList, err := svc.EditList(t.Context(), info, listID, dto.EditListRequest{})
			assert.Equal(t, models.List{}, gotList)
			assert.ErrorIs(t, err, tt.wantErr)

			repo.AssertExpectations(t)
		})
	}
}

func TestDeleteList_Success(t *testing.T) {
	type mockBehavior func(
		repo *mocks.Repository,
		list models.List,
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
				list models.List,
			) {
				list.OwnerID = models.UserID(uuid.New())

				repo.On("GetList", mock.Anything, list.ID).Return(list, nil)
				repo.On("DeleteList", mock.Anything, list.ID).Return(nil).Once()
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
				list models.List,
			) {
				repo.On("GetList", mock.Anything, list.ID).Return(list, nil)
				repo.On("DeleteList", mock.Anything, list.ID).Return(nil).Once()
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

			tt.setup(repo, list)

			err := svc.DeleteList(t.Context(), tt.info, list.ID)
			assert.NoError(t, err)

			repo.AssertExpectations(t)
		})
	}
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

			tt.setup(repo, listID)

			err := svc.DeleteList(t.Context(), info, listID)
			assert.ErrorIs(t, err, tt.wantErr)

			repo.AssertExpectations(t)
		})
	}
}
