package main

import (
	"fmt"
	"github.com/akaletr/yandex/internal/app"
	"log"
)

func main() {
	server, err := app.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	log.Fatal(server.Start())
}
