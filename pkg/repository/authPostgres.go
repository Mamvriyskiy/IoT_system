package repository

import (
	"github.com/jmoiron/sqlx"
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
)


type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db : db}
}

func (r *AuthPostgres) CreateUser(user pkg.User) (int, error) {
	return 0, nil
}
