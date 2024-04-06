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

func (r *DevicePostgres) CreateDevice(homeID int, device pkg.Devices) (int, error) {
	var id int
	fmt.Println("+")
	query := fmt.Sprintf("INSERT INTO %s (name, TypeDevice, Status, Brand, PowerConsumption, MinParametr, MaxParametr) values ($1, $2, $3, $4, $5, $6, $7) RETURNING deviceID", "device")
	row := r.db.QueryRow(query, device.Name, device.TypeDevice, device.Status, device.Brand, device.PowerConsumption, device.MinParameter, device.MaxParameter)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	query1 := fmt.Sprintf("INSERT INTO %s (homeID, deviceId) VALUES ($1, $2)", "deviceHome")
	r.db.QueryRow(query1, homeID, id)
	if err != nil {
    	return 0, err
	}

	return id, nil
}

func (r *DevicePostgres) DeleteDevice(idDevice int, device pkg.Devices) error {
	return nil
}
