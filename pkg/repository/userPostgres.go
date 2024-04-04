package repository

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user pkg.User) (int, error) {
	return 0, nil
}

func (r *UserPostgres) GetUserByEmail(emal string) (int, error) {
	return 0, nil
}

func (r *UserPostgres) GetPasswordByID(id int) (string, error) {
	return "", nil
}
