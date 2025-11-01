package memory

import (
	"github.com/Xiof22/ToDoList/internal/models"
)

func (repo *Repository) CreateList(title, description string) error {
	list := &models.List{
		ID:          repo.nextID,
		Title:       title,
		Description: description,
		Tasks:       make(map[int]*models.Task),
		NextID:      1,
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.Lists[repo.nextID] = list
	repo.nextID++

	return nil
}

func (repo *Repository) GetLists() ([]models.List, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	listCount := len(repo.Lists)
	lists := make([]models.List, 0, listCount)

	for _, list := range repo.Lists {
		lists = append(lists, *list)
	}

	return sortLists(lists), nil
}

func (repo *Repository) GetList(listID int) (*models.List, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	return repo.Lists[listID], nil
}

func (repo *Repository) EditList(listID int, title, description string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list, found := repo.Lists[listID]
	if !found {
		return ErrListNotFound
	}

	list.Title = title
	list.Description = description

	return nil
}

func (repo *Repository) DeleteList(listID int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, found := repo.Lists[listID]
	if !found {
		return ErrListNotFound
	}

	delete(repo.Lists, listID)
	return nil
}
