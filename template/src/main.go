package main

import (
	"log"
	"os"
)

func main() {

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := Route()

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = Config["port"]
	}

	log.Println("API listening at http://localhost:" + port)

	r.Run()

}
