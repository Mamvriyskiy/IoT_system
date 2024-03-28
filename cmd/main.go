package main

import "log"

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8889"); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}
