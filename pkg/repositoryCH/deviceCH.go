package repositoryCH

import (
	//"fmt"

	// "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/logger"
	// pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	// "github.com/jmoiron/sqlx"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type DevicePostgres struct {
	db driver.Conn
}

func NewDevicePostgres(db driver.Conn) *DevicePostgres {
	return &DevicePostgres{db: db}
}
