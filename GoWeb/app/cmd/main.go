package main

import (
	"app/internal/application"
	"log"
)

func main() {

	// init server
	sv := application.NewServer(":8080")
	err := sv.Run()
	if err != nil {
		log.Println(err)
	}
}
