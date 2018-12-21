package main

import (
	"fmt"
	"log"
	"os"
	"{{name}}/src/config"
	"{{name}}/src/models/bg"
	"{{name}}/src/models/orm"
	"{{name}}/src/models/wechat"
	"{{name}}/src/routers"
)

func main() {
	// this will make sure the db connection is initialized
	orm.InitDbConnection()

	// read wechat access token
	fmt.Println("wechat.enable", config.All["wechat.enable"])
	if config.All["wechat.enable"] == "true" {
		go func() { wechat.ForceRefreshGlobalAccessToken() }()
	}

	// bg tasks
	go func() { bg.DailyJob() }()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := routers.Route()

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = config.All["port"]
	}

	log.Println("API listening at http://localhost:" + port)

	r.Run(":" + port)

}
