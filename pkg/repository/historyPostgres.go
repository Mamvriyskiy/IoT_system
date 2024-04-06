package repository

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type DeviceHistoryPostgres struct {
	db *sqlx.DB
}

func NewDeviceHistoryPostgres(db *sqlx.DB) *DeviceHistoryPostgres {
	return &DeviceHistoryPostgres{db: db}
}


func (r *DeviceHistoryPostgres) CreateDeviceHistory(deviceID int, device pkg.DevicesHistory) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (timeWork, AverageIndicator, EnergyConsumed) values ($1, $2, $3) RETURNING historyDevID", "historyDev")
	row := r.db.QueryRow(query, device.TimeWork, device.AverageIndicator, device.EnergyConsumed)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	query2 := fmt.Sprintf("INSERT INTO %s (deviceID, historydevID) VALUES ($1, $2)", "historydevice")
	r.db.QueryRow(query2, deviceID, id)

	return id, nil
}

func (r *DeviceHistoryPostgres) UpdateDeviceHistory(id int, history pkg.DevicesHistory) error {
	return nil
}

func (r *DeviceHistoryPostgres) GetDeviceHistory(idDevice int) ([]pkg.DevicesHistory, error) {
	return nil, nil
}
