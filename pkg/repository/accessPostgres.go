package repository

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type AccessHomePostgres struct {
	db *sqlx.DB
}

func NewAccessHomePostgres(db *sqlx.DB) *AccessHomePostgres {
	return &AccessHomePostgres{db: db}
}

func (r *AccessHomePostgres) AddUser(access pkg.AccessHome) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (accessStatus, accessLevel) values ($1, $2) RETURNING accessID", "access")
	row := r.db.QueryRow(query, access.AccessStatus, access.Username, access.AccessLevel)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return 0, nil
}

func (r *AccessHomePostgres) DeleteUser(idUser int, access pkg.AccessHome) error {
	return nil
}

func (r *AccessHomePostgres) UpdateLevel(idUser int, access pkg.AccessHome) error {
	return nil
}

func (r *AccessHomePostgres) UpdateStatus(idUser int, access pkg.AccessHome) error {
	return nil
}

func (r *AccessHomePostgres) GetListUserHome(idHome int, home pkg.AccessHome) ([]pkg.User, error) {
	return nil, nil
}
