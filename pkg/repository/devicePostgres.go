package repository

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/jmoiron/sqlx"
)

type DevicePostgres struct {
	db *sqlx.DB
}

func NewDevicePostgres(db *sqlx.DB) *DevicePostgres {
	return &DevicePostgres{db: db}
}

func (r *DevicePostgres) CreateDevice(device *pkg.Devices) (int, error) {
	return 0, nil
}

func (r *DevicePostgres) DeleteDevice(idDevice int, device *pkg.Devices) error {
	return nil
}

func (r *DevicePostgres) UpdateDevice(idDevice int, device *pkg.Devices) error {
	return nil
}

func (r *DevicePostgres) AddHomeDevice(idHome, idDevice int, input *pkg.Devices) error {
	return nil
}

func (r *DevicePostgres) DeleteHomeDevice(idHome, idDevice int, input *pkg.Devices) error {
	return nil
}
