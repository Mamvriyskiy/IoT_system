package repositoryCH

import (
	//"fmt"

	// "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/logger"
	// pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	// "github.com/jmoiron/sqlx"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type DeviceHistoryPostgres struct {
	db driver.Conn
}

func NewDeviceHistoryPostgres(db driver.Conn) *DeviceHistoryPostgres {
	return &DeviceHistoryPostgres{db: db}
}
