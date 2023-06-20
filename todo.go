package todo

type List struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id      int
	Id_user int
	Id_list int
}

type Item struct {
	Id          int
	Title       string
	Description string
	Done        bool
}

type ListItem struct {
	Id      int
	Id_list int
	Id_item int
}
