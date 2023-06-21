package repository

import (
	"github.com/f1xend/todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type List interface {
	Create(userId int, list todo.List) (int, error)
	GetAll(userId int) ([]todo.List, error)
	GetById(userId, id int) (todo.List, error)
	Delete(userId, id int) error
	Update(userId, id int, input todo.UpdateListInput) error
}

type Item interface {
	Create(listId int, item todo.Item) (int, error)
	GetAll(userId, listId int) ([]todo.Item, error)
}

type Reposotory struct {
	Authorization
	List
	Item
}

func NewReposotory(db *sqlx.DB) *Reposotory {
	return &Reposotory{
		Authorization: NewAuthPostgres(db),
		List:          NewListPostgres(db),
		Item:          NewItemPostgres(db),
	}
}
