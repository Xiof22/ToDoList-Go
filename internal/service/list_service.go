package service

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
)

func (svc *Service) CreateList(ctx context.Context, info models.UserInfo, req dto.CreateListRequest) (models.List, error) {
	list := models.NewList(info.ID, req.Title, req.Description)
	return svc.repo.CreateList(ctx, list)
}

func (svc *Service) GetLists(ctx context.Context, info models.UserInfo) ([]models.List, error) {
	if info.Role == models.Admin {
		return svc.repo.GetLists(ctx)
	}

	return svc.repo.GetListsByUserID(ctx, info.ID)
}

func (svc *Service) GetList(ctx context.Context, info models.UserInfo, listID models.ListID) (models.List, error) {
	list, err := svc.repo.GetList(ctx, listID)
	if err != nil {
		return list, err
	}

	if list.OwnerID != info.ID && info.Role != models.Admin {
		return models.List{}, errorsx.ErrForbidden
	}

	return list, nil
}

func (svc *Service) EditList(ctx context.Context, info models.UserInfo, listID models.ListID, req dto.EditListRequest) (models.List, error) {
	list, err := svc.repo.GetList(ctx, listID)
	if err != nil {
		return list, err
	}

	if list.OwnerID != info.ID && info.Role != models.Admin {
		return models.List{}, errorsx.ErrForbidden
	}

	list.Title = req.Title
	list.Description = req.Description

	return svc.repo.EditList(ctx, listID, list)
}

func (svc *Service) DeleteList(ctx context.Context, info models.UserInfo, listID models.ListID) error {
	if list, err := svc.repo.GetList(ctx, listID); err != nil {
		return err
	} else if list.OwnerID != info.ID && info.Role != models.Admin {
		return errorsx.ErrForbidden
	}

	svc.repo.DeleteList(ctx, listID)
	return nil
}
