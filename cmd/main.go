package main

import (
	"log"
	handler "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/handler"
	src "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(src.Server)
	if err := srv.Run("8889", handlers.InitRouters()); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}
