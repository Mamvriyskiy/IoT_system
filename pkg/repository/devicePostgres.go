package repository

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type DevicePostgres struct {
	db *sqlx.DB
}

func NewDevicePostgres(db *sqlx.DB) *DevicePostgres {
	return &DevicePostgres{db: db}
}


func (r *DevicePostgres) CreateDevice(device *pkg.Devices) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, TypeDevice, Status, Brand, PowerConsumption, MinParameter, MaxParameter) values ($1, $2, $3, $4, $5, $6, $7) RETURNING deviceID", "device")
	row := r.db.QueryRow(query, device.Name, device.TypeDevice, device.Status, device.Brand, device.PowerConsumption, device.MinParameter, device.MaxParameter)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
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
