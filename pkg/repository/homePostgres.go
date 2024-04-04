package repository

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/jmoiron/sqlx"
)

type HomePostgres struct {
	db *sqlx.DB
}

func NewHomePostgres(db *sqlx.DB) *HomePostgres {
	return &HomePostgres{db: db}
}

func (r *HomePostgres) CreateHome(ownerID int, home pkg.Home) (int, error) {
	return 0, nil
}

func (r *HomePostgres) DeleteHome(ownerID int, home pkg.Home) error {
	return nil
}

func (r *HomePostgres) UpdateHome(ownerID int, home pkg.Home) error {
	return nil
}
