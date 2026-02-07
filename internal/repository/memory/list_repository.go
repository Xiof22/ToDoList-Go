package memory

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
)

func (repo *Repository) CreateList(ctx context.Context, list models.List) (models.List, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.Lists[list.ID] = &list

	return list, nil
}

func (repo *Repository) GetLists(ctx context.Context) ([]models.List, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	lists := make([]models.List, 0, len(repo.Lists))
	for _, list := range repo.Lists {
		lists = append(lists, *list)
	}

	return lists, nil
}

func (repo *Repository) GetListsByUserID(ctx context.Context, userID models.UserID) ([]models.List, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var lists []models.List
	for _, list := range repo.Lists {
		if list.OwnerID == userID {
			lists = append(lists, *list)
		}
	}

	return lists, nil
}

func (repo *Repository) GetList(ctx context.Context, listID models.ListID) (models.List, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list, ok := repo.Lists[listID]
	if !ok {
		return models.List{}, errorsx.ErrListNotFound
	}

	return *list, nil
}

func (repo *Repository) EditList(ctx context.Context, listID models.ListID, list models.List) (models.List, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.Lists[listID] = &list
	return list, nil
}

func (repo *Repository) DeleteList(ctx context.Context, listID models.ListID) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	delete(repo.Lists, listID)
	return nil
}
