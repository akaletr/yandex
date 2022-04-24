package main

import (
	"github.com/akaletr/yandex/internal/app"
	"log"
)

func main() {
	server, err := app.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.Start())
}
