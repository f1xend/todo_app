package repository

import (
	"github.com/f1xend/todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type List interface{}

type Item interface{}

type Reposotory struct {
	Authorization
	List
	Item
}

func NewReposotory(db *sqlx.DB) *Reposotory {
	return &Reposotory{
		Authorization: NewAuthPostgres(db),
	}
}
