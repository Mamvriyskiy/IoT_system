package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"fmt"
)

type Config struct {
	Host string
	Port string
	Username string
	Password string
	DBName string
	SSLMode string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
        return nil, err
    }
	fmt.Println("1) ",err)
	err = db.Ping()
	fmt.Println("2) ",err)
	if err != nil {
		return nil, err
	}

	return db, nil
}
