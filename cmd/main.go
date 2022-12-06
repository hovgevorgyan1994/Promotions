package main

import (
	server "awesomeProject"
	"log"
)

func main() {
	srv := new(server.Server)
	if err := srv.Run("1321"); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
