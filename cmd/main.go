package main

import (
	// "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository".
	"log"

	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	handler "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/handler"
)

func main() {
	// *TODO repository.NewPostgresDB
	// repos := repository.NewRepository()
	// services := service.NewServices(repos)
	// handlers := handler.NewHandler(services)
	handlers := new(handler.Handler)

	srv := new(pkg.Server)
	if err := srv.Run("8889", handlers.InitRouters()); err != nil {
		log.Fatal("error occurred while running http server: ", err)
	}
}
