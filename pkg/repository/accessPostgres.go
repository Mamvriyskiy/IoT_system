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

	query2 := fmt.Sprintf("INSERT INTO %s (clientID, accessID) VALUES ($1, $2)", "accessClient")
	r.db.QueryRow(query2, userID, id)
	if err != nil {
    	return 0, err
	}

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
	query := fmt.Sprintf(`
	UPDATE access
	SET accesslevel = $1
	WHERE accessid = (
		SELECT accessid FROM accessclient WHERE clientid = $2
	);`)
	fmt.Println(idUser, access)
	_, err := r.db.Exec(query, access.AccessLevel, idUser)
	// if err != nil {
	// 	// Обработка ошибки, если запрос не удалось выполнить
	//     fmt.Println("Ошибка выполнения запроса:", err, result)
	// 	return err
	// }

	// rowsAffected, err := result.RowsAffected()
	// if err != nil {
	// 	// Обработка ошибки, если не удалось получить количество затронутых строк
	// 	fmt.Println("Ошибка получения количества затронутых строк:", err)
	// 	return err
	// }

	// if rowsAffected == 0 {
	// 	fmt.Println("Не было обновлено ни одной строки")
	// 	return nil
	// }

	return err
}

func (r *AccessHomePostgres) UpdateStatus(idUser int, access pkg.AccessHome) error {
	query := fmt.Sprintf(`
	UPDATE access
	SET accessstatus = $1
	WHERE accessid = (
		SELECT accessid FROM accessclient WHERE clientid = $2
	);`)
	fmt.Println(idUser, access)
	_, err := r.db.Exec(query, access.AccessStatus, idUser)

	return err
}

func (r *AccessHomePostgres) GetListUserHome(idHome int, home pkg.AccessHome) ([]pkg.User, error) {
	return nil, nil
}
