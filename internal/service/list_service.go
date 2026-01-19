package service

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
)

func (svc *Service) CreateList(ctx context.Context, info models.UserInfo, req dto.CreateListRequest) (models.List, error) {
	list := models.NewList(info.ID, req.Title, req.Description)
	return svc.repo.CreateList(ctx, list), nil
}

func (svc *Service) GetLists(ctx context.Context, info models.UserInfo) []models.List {
	if info.Role == models.Admin {
		return svc.repo.GetLists(ctx)
	}

	return svc.repo.GetListsByUserID(ctx, info.ID)
}

func (svc *Service) GetList(ctx context.Context, info models.UserInfo, listID int) (models.List, error) {
	if listID <= 0 {
		return models.List{}, errorsx.ErrInvalidListID
	}

	list, err := svc.repo.GetList(ctx, listID)
	if err != nil {
		return list, err
	}

	if list.OwnerID != info.ID && info.Role != models.Admin {
		return models.List{}, errorsx.ErrForbidden
	}

	return list, nil
}

func (svc *Service) EditList(ctx context.Context, info models.UserInfo, listID int, req dto.EditListRequest) (models.List, error) {
	if listID <= 0 {
		return models.List{}, errorsx.ErrInvalidListID
	}

	list, err := svc.repo.GetList(ctx, listID)
	if err != nil {
		return list, err
	}

	if list.OwnerID != info.ID && info.Role != models.Admin {
		return models.List{}, errorsx.ErrForbidden
	}

	list.Title = req.Title
	list.Description = req.Description

	err = svc.repo.EditList(ctx, listID, list)
	return list, err
}

func (svc *Service) DeleteList(ctx context.Context, info models.UserInfo, listID int) error {
	if listID <= 0 {
		return errorsx.ErrInvalidListID
	}

	if list, err := svc.repo.GetList(ctx, listID); err != nil {
		return err
	} else if list.OwnerID != info.ID && info.Role != models.Admin {
		return errorsx.ErrForbidden
	}

	svc.repo.DeleteList(ctx, listID)
	return nil
}
