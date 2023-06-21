package service

import (
	"github.com/f1xend/todo-app"
	"github.com/f1xend/todo-app/package/repository"
)

type ItemService struct {
	repo     repository.Item
	listRepo repository.List
}

func NewItemService(repo repository.Item, listRepo repository.List) *ItemService {
	return &ItemService{repo: repo, listRepo: listRepo}
}

func (s *ItemService) Create(userId, listId int, item todo.Item) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, nil
	}
	return s.repo.Create(listId, item)
}

func (s *ItemService) GetAll(userId, listId int) ([]todo.Item, error) {
	return s.repo.GetAll(userId, listId)
}
