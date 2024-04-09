package repository

import (
	"fmt"

	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/jmoiron/sqlx"
)

type AccessHomePostgres struct {
	db *sqlx.DB
}

func NewAccessHomePostgres(db *sqlx.DB) *AccessHomePostgres {
	return &AccessHomePostgres{db: db}
}

func (r *AccessHomePostgres) AddUser(homeID, userID int, access pkg.AccessHome) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (accessStatus, accessLevel) 
		values ($1, $2) RETURNING accessID`, "access")
	row := r.db.QueryRow(query, access.AccessStatus, access.AccessLevel)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	query2 := fmt.Sprintf("INSERT INTO %s (clientID, accessID) VALUES ($1, $2)", "accessClient")

	result, err := r.db.Exec(query2, userID, id)
	if err != nil {
		// Обработка ошибки, если запрос не удалось выполнить
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Обработка ошибки, если не удалось получить количество затронутых строк

		return 0, err
	}

	if rowsAffected == 0 {
		return 0, nil
	}

	query3 := fmt.Sprintf("INSERT INTO %s (homeID, accessID) VALUES ($1, $2)", "accessHome")
	r.db.QueryRow(query3, homeID, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AccessHomePostgres) UpdateLevel(idUser int, access pkg.AccessHome) error {
	query := fmt.Sprintf(`
	UPDATE access
	SET accesslevel = $1
	WHERE accessid = (
		SELECT accessid FROM accessclient WHERE clientid = $2
	);`)
	_, err := r.db.Exec(query, access.AccessLevel, idUser)

	return err
}

func (r *AccessHomePostgres) UpdateStatus(idUser int, access pkg.AccessHome) error {
	query := fmt.Sprintf(`
	UPDATE access
		SET accessstatus = $1
			WHERE accessid = (
				SELECT accessid FROM accessclient WHERE clientid = $2
	);`)
	_, err := r.db.Exec(query, access.AccessStatus, idUser)

	return err
}

func (r *AccessHomePostgres) GetListUserHome(idHome int) ([]pkg.ClientHome, error) {
	var lists []pkg.ClientHome
	query := fmt.Sprintf(`select c.login, a.accesslevel, a.accessstatus from client c 
							join accessclient ac on c.clientid = ac.clientid
								join access a on a.accessid = ac.accessid 
									join accesshome ah on ah.accessid = a.accessid 
										where ah.homeid = $1;`)
	err := r.db.Select(&lists, query, idHome)
	if err != nil {
		return nil, err
	}

	return lists, nil
}

func (r *AccessHomePostgres) DeleteUser(idUser int) error {
	query := fmt.Sprintf(`DELETE FROM access 
								where accessid 
									in (select accessid 
											from accessclient where clientid = $1);`)
	_, err := r.db.Exec(query, idUser)

	return err
}
