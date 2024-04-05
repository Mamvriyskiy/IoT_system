package repository

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/jmoiron/sqlx"
)

type DeviceHistoryPostgres struct {
	db *sqlx.DB
}

func NewDeviceHistoryPostgres(db *sqlx.DB) *DeviceHistoryPostgres {
	return &DeviceHistoryPostgres{db: db}
}


func (r *DeviceHistoryPostgres) CreateDeviceHistory(device pkg.DevicesHistory) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (timeWork, AverageIndicator, EnergyConsumed) values ($1, $2, $3) RETURNING historyDevID", "historyDev")
	row := r.db.QueryRow(query, device.TimeWork, device.AverageIndicator, device.AccessLevel)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return 0, nil
}

func (r *DeviceHistoryPostgres) UpdateDeviceHistory(id int, history pkg.DevicesHistory) error {
	return nil
}

func (r *DeviceHistoryPostgres) GetDeviceHistory(idDevice int) ([]pkg.DevicesHistory, error) {
	return nil, nil
}
