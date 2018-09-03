package router

import (
	"../controller"
	"../router/middleware"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Authorization())

	r.GET("/", controller.SayHello)
	r.GET("/api", controller.SayHello)
	r.GET("/api/v1", controller.SayHello)
	r.Static("/api/v1/static", "./static")
	r.StaticFile("/favicon.ico", "./static/img/favicon.ico")

	// r.Static("/userdata", "./userdata")

	// Simple group: v1
	{

		{
			userR := v1.Group("/hello")

			// temp
			userR.POST("/hello", controller.SayHello)
		}

	}

	return r
}
