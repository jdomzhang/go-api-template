package main

import (
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	r := gin.Default()

	r.GET("/", SayHello)
	r.GET("/api", SayHello)
	r.GET("/api/v1", SayHello)
	r.Static("/api/v1/static", "./static")
	r.StaticFile("/favicon.ico", "./static/img/favicon.ico")

	// r.Static("/userdata", "./userdata")

	// Simple group: v1
	{
		v1 := r.Group("/api/v1")

		{
			say := v1.Group("/say")

			// temp
			say.POST("/hello", SayHello)
		}

	}

	return r
}
