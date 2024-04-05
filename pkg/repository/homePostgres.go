package repository

import (
	"fmt"
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
	var id int
	query := fmt.Sprintf("INSERT INTO %s (ownerid, name) values ($1, $2) RETURNING homeID", "home")
	row := r.db.QueryRow(query, ownerID, home.name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *HomePostgres) DeleteHome(ownerID int, home pkg.Home) error {
	return nil
}

func (r *HomePostgres) UpdateHome(ownerID int, home pkg.Home) error {
	return nil
}
