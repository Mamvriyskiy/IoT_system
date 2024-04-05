package repository

import (
	"fmt"

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
	var id int
	query := fmt.Sprintf("INSERT INTO %s (password, login, email) values ($1, $2, $3) RETURNING clientid", "client")
	row := r.db.QueryRow(query, user.Password, user.Username, user.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) GetUserByEmail(emal string) (int, error) {
	return 0, nil
}

func (r *UserPostgres) GetPasswordByID(id int) (string, error) {
	return "", nil
}
