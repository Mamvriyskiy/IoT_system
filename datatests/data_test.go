package tests_test

import (
	"testing"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/repository"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/configs"
)

func TestDeviceData(t *testing.T) {
	if err := configs.initConfig(); err != nil {
		return err
	}

	if err := godotenv.Load(); err != nil {
		return err
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	homePostgres := NewHomePostgres(db)

	home := pkg.home{
		Name : "home",
	}
	homeID := homePostgres.CreateHome(1, home)

	res, err = homePostgres.GetHomeByID(homeID)

	assert.Equal(t, home.Name, res.Name)
	assert.Equal(t, home.OwnerID, 1)
}
