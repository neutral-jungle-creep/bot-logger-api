package main

import (
	"bot-logger-front/pkg"
	"log"
)

func main() {
	srv := new(pkg.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error http server, %s", err.Error())
	}
}
