package service

import (
	"github.com/f1xend/todo-app"
	"github.com/f1xend/todo-app/package/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type List interface {
	Create(userId int, list todo.List) (int, error)
	GetAll(userId int) ([]todo.List, error)
	GetById(userId, id int) (todo.List, error)
	Delete(userId, id int) error
	Update(userId, id int, input todo.UpdateListInput) error
}

type Item interface {
	Create(userId, listId int, item todo.Item) (int, error)
	GetAll(userId, listId int) ([]todo.Item, error)
	GetById(userId, itemId int) (todo.Item, error)
	Delete(userId, itemId int) error
}

type Service struct {
	Authorization
	List
	Item
}

func NewService(repos *repository.Reposotory) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		List:          NewListService(repos.List),
		Item:          NewItemService(repos.Item, repos.List),
	}
}
