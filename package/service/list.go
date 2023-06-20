package service

import (
	"github.com/f1xend/todo-app"
	"github.com/f1xend/todo-app/package/repository"
)

type ListService struct {
	repo repository.List
}

func NewListService(repo repository.List) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) Create(userId int, list todo.List) (int, error) {
	return s.repo.Create(userId, list)
}
