package repository

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user pkg.User) (int, error)
}


type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization : NewAuthPostgres(db),
	}
}
