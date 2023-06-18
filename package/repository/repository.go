package repository

type Authorization interface {
}

type List interface{}

type Item interface{}

type Reposotory struct {
	Authorization
	List
	Item
}

func NewReposotory() *Reposotory {
	return &Reposotory{}
}
