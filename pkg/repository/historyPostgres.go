package repository

import (
	"fmt"

	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/jmoiron/sqlx"
)

type DeviceHistoryPostgres struct {
	db *sqlx.DB
}

func NewDeviceHistoryPostgres(db *sqlx.DB) *DeviceHistoryPostgres {
	return &DeviceHistoryPostgres{db: db}
}

func (r *DeviceHistoryPostgres) CreateDeviceHistory(userID int, history pkg.AddHistory) (int, error) {
	var homeID int
	queryHomeID := `select h.homeid from home h 
	where h.homeid in (select a.homeid from accesshome a 
		where a.accessid in (select a.accessid from accessclient a 
			JOIN access ac ON a.accessid = ac.accessid where clientid = $1));`
	err := r.db.Get(&homeID, queryHomeID, userID)
	fmt.Println("1", err)

	var deviceID int
	querDeviceID := `select d.deviceid from device d join devicehome dh on d.deviceid = dh.deviceid 
		where dh.homeid = $1 and d.name = $2;`
	err = r.db.Get(&deviceID, querDeviceID, homeID, history.Name)
	fmt.Println("2", err)
			
	var id int
	query := fmt.Sprintf(`INSERT INTO %s 
		(timeWork, AverageIndicator, EnergyConsumed) 
			values ($1, $2, $3) RETURNING historyDevID`, "historyDev")
	row := r.db.QueryRow(query, history.TimeWork, history.AverageIndicator, history.EnergyConsumed)
	err = row.Scan(&id)
	fmt.Println("3", err)
	if err != nil {
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (deviceID, historydevID) VALUES ($1, $2)", "historydevice")
	result, err := r.db.Exec(query, deviceID, id)
	fmt.Println("4", err)
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

	return id, nil
}

func (r *DeviceHistoryPostgres) GetDeviceHistory(userID int, name string) ([]pkg.DevicesHistory, error) {
	var homeID int
	queryHomeID := `select h.homeid from home h 
	where h.homeid in (select a.homeid from accesshome a 
		where a.accessid in (select a.accessid from accessclient a 
			JOIN access ac ON a.accessid = ac.accessid where clientid = $1));`
	err := r.db.Get(&homeID, queryHomeID, userID)
	fmt.Println("1", err)

	var deviceID int
	querDeviceID := `select d.deviceid from device d join devicehome dh on d.deviceid = dh.deviceid 
		where dh.homeid = $1 and d.name = $2;`
	err = r.db.Get(&deviceID, querDeviceID, homeID, name)
	fmt.Println("2", err)

	var lists []pkg.DevicesHistory
	query := `select hi.timework, hi.averageindicator, hi.energyconsumed 
		from historydev as hi join historydevice as hd on hi.historydevid = hd.historydevid 
			where hd.deviceid = $1`
	err = r.db.Select(&lists, query, deviceID)
	if err != nil {
		return nil, err
	}

	return lists, nil
}
