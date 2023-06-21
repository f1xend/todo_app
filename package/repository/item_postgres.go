package repository

import (
	"fmt"

	"github.com/f1xend/todo-app"
	"github.com/jmoiron/sqlx"
)

type ItemPostgres struct {
	db *sqlx.DB
}

func NewItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) Create(listId int, item todo.Item) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemquery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", itemTable)
	row := tx.QueryRow(createItemquery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (id_list, id_item) VALUES ($1, $2)", listItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *ItemPostgres) GetAll(userId, listId int) ([]todo.Item, error) {
	var items []todo.Item
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.id_item = ti.id 
						  INNER JOIN %s ul ON ul.id_list = li.id_list WHERE ul.id_user=$1 AND li.id_list=$2`,
		itemTable, listItemsTable, usersListTable)

	if err := r.db.Select(&items, query, userId, listId); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ItemPostgres) GetById(userId, itemId int) (todo.Item, error) {
	var item todo.Item
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.id_item = ti.id 
						  INNER JOIN %s ul ON ul.id_list = li.id_list WHERE ul.id_user=$1 AND ti.id=$2`,
		itemTable, listItemsTable, usersListTable)

	if err := r.db.Get(&item, query, userId, itemId); err != nil {
		return item, err
	}
	return item, nil
}
