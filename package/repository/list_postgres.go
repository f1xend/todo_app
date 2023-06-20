package repository

import (
	"fmt"

	"github.com/f1xend/todo-app"
	"github.com/jmoiron/sqlx"
)

const (
	todoListsTable    = "list"
	todoUserListTable = "user_list"
)

type ListPostgres struct {
	db *sqlx.DB
}

func NewListPostgres(db *sqlx.DB) *ListPostgres {
	return &ListPostgres{db: db}
}

func (r *ListPostgres) Create(userId int, list todo.List) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, descriptiuon) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (id_user, id_list) VALUES($1,$2)", todoUserListTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *ListPostgres) GetAll(userId int) ([]todo.List, error) {
	var lists []todo.List

	query := fmt.Sprintf("SELECT l.id, t.title, t.description FROM %s l INNER JOIN %s ul on ul.id_list = l.id WHERE id_user = $1", todoListsTable, todoUserListTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *ListPostgres) GetById(userId, id int) (todo.List, error) {
	var list todo.List

	query := fmt.Sprintf("SELECT l.id, t.title, t.description FROM %s l INNER JOIN %s ul on ul.id_list = l.id WHERE id_user = $1 AND ul.id_list = $2", todoListsTable, todoUserListTable)
	err := r.db.Get(&list, query, userId, id)

	return list, err
}
