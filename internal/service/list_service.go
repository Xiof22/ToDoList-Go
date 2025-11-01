package service

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/models"
)

func (svc *Service) CreateList(ctx context.Context, req dto.CreateListRequest) (models.List, error) {
	list := models.NewList(req.Title, req.Description)
	return svc.repo.CreateList(ctx, list)
}

func (svc *Service) GetLists(ctx context.Context) ([]models.List, error) {
	return svc.repo.GetLists(ctx)
}

func (svc *Service) GetList(ctx context.Context, listID models.ListID) (models.List, error) {
	list, err := svc.repo.GetList(ctx, listID)
	if err != nil {
		return list, err
	}

	return list, nil
}

func (svc *Service) EditList(ctx context.Context, listID models.ListID, req dto.EditListRequest) (models.List, error) {
	list, err := svc.repo.GetList(ctx, listID)
	if err != nil {
		return list, err
	}

	list.Title = req.Title
	list.Description = req.Description

	return svc.repo.EditList(ctx, listID, list)
}

func (svc *Service) DeleteList(ctx context.Context, listID models.ListID) error {
	if _, err := svc.repo.GetList(ctx, listID); err != nil {
		return err
	}

	svc.repo.DeleteList(ctx, listID)
	return nil
}
