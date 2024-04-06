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

func (r *AccessHomePostgres) AddUser(homeID, userID int, access pkg.AccessHome) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (accessStatus, accessLevel) values ($1, $2) RETURNING accessID", "access")
	row := r.db.QueryRow(query, access.AccessStatus, access.AccessLevel)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("1")
	query2 := fmt.Sprintf("INSERT INTO %s (clientID, accessID) VALUES ($1, $2)", "accessClient")
	r.db.QueryRow(query2, userID, id)
	if err != nil {
    	return 0, err
	}

	fmt.Println("2")
	query3 := fmt.Sprintf("INSERT INTO %s (homeID, accessID) VALUES ($1, $2)", "accessHome")
	r.db.QueryRow(query3, homeID, id)
	if err != nil {
    	return 0, err
	}
	
	return id, nil
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
