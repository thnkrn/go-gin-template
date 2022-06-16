package main

import (
	"log"

	di "github.com/thnkrn/go-gin-template/pkg/di"
)

func main() {
	server, diErr := di.InitializeAPI()
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
