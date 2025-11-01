package service

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/models"
)

func (svc *Service) CreateList(ctx context.Context, req dto.CreateListRequest) models.List {
	return svc.repo.CreateList(ctx, req)
}

func (svc *Service) GetLists(ctx context.Context) []models.List {
	return svc.repo.GetLists(ctx)
}

func (svc *Service) GetList(ctx context.Context, req dto.ListIdentifier) (*models.List, bool) {
	return svc.repo.GetList(ctx, req)
}

func (svc *Service) EditList(ctx context.Context, req dto.EditListRequest) (models.List, error) {
	if _, found := svc.repo.GetList(ctx, dto.ListIdentifier{
		ID: req.ListID,
	}); !found {
		return models.List{}, ErrListNotFound
	}

	return svc.repo.EditList(ctx, req), nil
}

func (svc *Service) DeleteList(ctx context.Context, req dto.ListIdentifier) error {
	if _, found := svc.repo.GetList(ctx, req); !found {
		return ErrListNotFound
	}

	svc.repo.DeleteList(ctx, req)
	return nil
}
