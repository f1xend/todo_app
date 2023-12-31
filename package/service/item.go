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

func (s *ItemService) GetById(userId, itemId int) (todo.Item, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *ItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *ItemService) Update(userId, id int, input todo.UpdateItemInput) error {
	return s.repo.Update(userId, id, input)
}
