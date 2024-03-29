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
	return 0, nil
}

func (r *DeviceHistoryPostgres) UpdateDeviceHistory(idDevice int, history pkg.DevicesHistory) error {
	return nil
}


func (r *DeviceHistoryPostgres) GetDeviceHistory(idDevice int) ([]pkg.DevicesHistory, error) {
	return nil, nil
}
