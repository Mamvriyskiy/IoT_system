package repository

import (
	"fmt"

	logger "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3"
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
	query := fmt.Sprintf(`INSERT INTO %s (password, login, email) 
		values ($1, $2, $3) RETURNING clientid`, "client")
	row := r.db.QueryRow(query, user.Password, user.Username, user.Email)
	if err := row.Scan(&id); err != nil {
		logger.Log("Error", "Scan", "Error insert into client:", err, id)
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) GetUser(login, password string) (pkg.User, error) {
	var user pkg.User
	query := fmt.Sprintf("SELECT clientid from %s where login = $1 and password = $2", "client")
	err := r.db.Get(&user, query, login, password)

	return user, err
}
