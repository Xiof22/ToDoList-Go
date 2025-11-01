package memory

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/models"
)

func (repo *Repository) CreateList(ctx context.Context, req dto.CreateListRequest) models.List {
	list := models.List{
		ID:          repo.nextID,
		Title:       req.Title,
		Description: req.Description,
		Tasks:       make(map[int]*models.Task),
		NextID:      1,
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.Lists[repo.nextID] = &list
	repo.nextID++

	return list
}

func (repo *Repository) GetLists(ctx context.Context) []models.List {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	listCount := len(repo.Lists)
	lists := make([]models.List, 0, listCount)

	for _, list := range repo.Lists {
		lists = append(lists, *list)
	}

	return sortLists(lists)
}

func (repo *Repository) GetList(ctx context.Context, req dto.ListIdentifier) (*models.List, bool) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list, found := repo.Lists[req.ID]
	return list, found
}

func (repo *Repository) EditList(ctx context.Context, req dto.EditListRequest) models.List {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[req.ListID]
	list.Title = req.Title
	list.Description = req.Description

	return *list
}

func (repo *Repository) DeleteList(ctx context.Context, req dto.ListIdentifier) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	delete(repo.Lists, req.ID)
}
