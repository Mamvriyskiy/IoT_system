package main

import (
	"log"
	src "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/src/handler"
)

func main() {
	handlers := new(src.Handler)
	srv := new(src.Server)
	if err := srv.Run("8889", handlers.InitRouters()); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}
