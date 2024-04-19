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

func (r *DeviceHistoryPostgres) CreateDeviceHistory(deviceID int,
	device pkg.DevicesHistory,
) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s 
		(timeWork, AverageIndicator, EnergyConsumed) 
			values ($1, $2, $3) RETURNING historyDevID`, "historyDev")
	row := r.db.QueryRow(query, device.TimeWork, device.AverageIndicator, device.EnergyConsumed)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (deviceID, historydevID) VALUES ($1, $2)", "historydevice")
	result, err := r.db.Exec(query, deviceID, id)
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

func (r *DeviceHistoryPostgres) GetDeviceHistory(idDevice int) ([]pkg.DevicesHistory, error) {
	var lists []pkg.DevicesHistory
	query := fmt.Sprintf(`select hi.timework, hi.averageindicator, hi.energyconsumed 
		from %s as hi join %s as hd on hi.historydevid = hd.historydevid 
			where hd.deviceid = $1`, "historydev", "historydevice")
	err := r.db.Select(&lists, query, idDevice)
	if err != nil {
		return nil, err
	}

	return lists, nil
}
