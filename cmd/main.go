package main

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
	handler "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/handler"
	"log"
)

func main() {
	// *TODO repository.NewPostgresDB
	repos := repository.NewRepository()
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)

	srv := new(pkg.Server)
	if err := srv.Run("8889", handlers.InitRouters()); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}
