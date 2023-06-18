package service

import "github.com/f1xend/todo-app/package/repository"

type Authorization interface {
}

type List interface{}

type Item interface{}

type Service struct {
	Authorization
	List
	Item
}

func NewService(repos *repository.Reposotory) *Service {
	return &Service{}
}
