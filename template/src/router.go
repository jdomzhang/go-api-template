package main

import (
	"github.com/gin-gonic/gin"
)

// Route contains all the API mapping
func Route() *gin.Engine {
	r := gin.Default()

	r.GET("/", SayHello)
	r.StaticFile("/favicon.ico", "./static/img/favicon.ico")

	r.GET("/api", SayHello)
	r.GET("/api/v1", SayHello)
	r.GET("/api/v1/hello", SayHello)
	r.Static("/api/v1/static", "./static")

	return r
}
