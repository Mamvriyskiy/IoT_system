package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	handler "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/handler"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		// *TODO: log
		return
	}

	if err := godotenv.Load(); err != nil {
		// *TODO: log
		return
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		fmt.Println("no connect DB", err)
		return
	}
	repos := repository.NewRepository(db)
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)

	srv := new(pkg.Server)
	if err := srv.Run("8000", handlers.InitRouters()); err != nil {
		log.Fatal("error occurred while running http server: ", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
