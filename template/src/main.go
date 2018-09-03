package main

import (
	"log"
	"os"

	"./models/orm"
	"./router"
)

func main() {

	// this will make sure the db connection is initialized
	orm.InitDbConnection()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := router.Route()

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "9090"
	}

	log.Println("API listening at http://localhost:" + port)

	r.Run()

}
