package repositoryCH

import (
	"fmt"
	//_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/jmoiron/sqlx"
	//"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/logger"
	"github.com/ClickHouse/clickhouse-go"
	// Импорт драйвера PostgreSQL для его регистрации.
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewClickHouseDB(cfg *Config) (*sqlx.DB, error) {
	//fmt.Println(fmt.Sprintf("http://%s:%s/?user=%s&password=%s", cfg.Host, cfg.Port, cfg.Username, ""))
	//connStr := fmt.Sprintf("http://%s:%s/?user=%s&password=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password)		
	//connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s",
		//cfg.Host, cfg.Port, cfg.Username, "")
	connect, err := clickhouse.Open("tcp://localhost:8321?database=default&username=default&password=")
	if err != nil {
		panic(err)
	}
	fmt.Println("DB: ", connect)

	//driver := "clickhouse"	
	// clickhouse.NewConfig() // need this - compile will fall cause import not used
	// connect, err := sqlx.Open(driver, connStr)
	// connect, err := sqlx.Open("clickhouse", "tcp://127.0.0.1:8123?database=default&username=default&password=")
	// fmt.Println(connect, err)

	// err = connect.Ping()
	// if err != nil {
	// 	logger.Log("Error", " sqlx.Open", "Error connect DB:", err, "clickhouse", fmt.Sprintf(
	// 		"host=%s port=%s user=%s dbname=%s password='' sslmode=%s",
	// 		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.SSLMode))
	// 	logger.Log("Error", "Ping()", "Error check connection:", err, "")
	// 	return nil, err
	// }

	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

    return db, err
}
