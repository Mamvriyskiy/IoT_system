package repositoryCH

import (
	//"fmt"

	//"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/logger"
	//pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type HomePostgres struct {
	db driver.Conn
}

func NewHomePostgres(db driver.Conn) *HomePostgres {
	return &HomePostgres{db: db}
}
