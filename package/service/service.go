package service

import (
	"github.com/f1xend/todo-app"
	"github.com/f1xend/todo-app/package/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type List interface{}

type Item interface{}

type Service struct {
	Authorization
	List
	Item
}

func NewService(repos *repository.Reposotory) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
