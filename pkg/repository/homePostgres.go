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
	row := r.db.QueryRow(query, ownerID, home.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *HomePostgres) DeleteHome(homeID int) error {
	query1 := fmt.Sprintf(`DELETE FROM access 
		WHERE accessid IN (SELECT accessid 
			FROM accesshome WHERE homeid = $1);`)
	_, err := r.db.Exec(query1, homeID)
	if err != nil {
		return err
	}

	query2 := fmt.Sprintf(`DELETE FROM historydev 
		WHERE historydevid IN (SELECT historydevid 
			FROM historydevice WHERE deviceid 
				IN (SELECT deviceid FROM devicehome WHERE homeid = $1));`)
	_, err = r.db.Exec(query2, homeID)
	if err != nil {
		return err
	}

	query3 := fmt.Sprintf(`DELETE FROM device 
		WHERE deviceid IN (SELECT deviceid 
			FROM devicehome WHERE homeid = $1);`)
	_, err = r.db.Exec(query3, homeID)
	if err != nil {
		return err
	}

	query4 := fmt.Sprintf(`DELETE FROM home
		WHERE homeid = $1;`)
	_, err = r.db.Exec(query4, homeID)
	if err != nil {
		return err
	}

	return err
}

func (r *HomePostgres) UpdateHome(home pkg.Home) error {
	query := fmt.Sprintf(`
	UPDATE home
		SET name = $1
		WHERE homeid = $2;`)
	result, err := r.db.Exec(query, home.Name, home.ID)
	if err != nil {
		// Обработка ошибки, если запрос не удалось выполнить

		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Обработка ошибки, если не удалось получить количество затронутых строк

		return err
	}

	if rowsAffected == 0 {
		return nil
	}

	return err
}

func (r *HomePostgres) GetHomeByID(homeID int) (pkg.Home, error) {
	var home pkg.Home
	query := fmt.Sprintf("SELECT * from %s where homeid = $1", "home")
	err := r.db.Get(&home, query, homeID)

	return home, err
}
