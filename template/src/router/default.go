package router

import (
	"../controllers"
	"../router/middleware"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Authorization())

	r.GET("/", controllers.SayHello)
	r.GET("/api", controllers.SayHello)
	r.GET("/api/v1", controllers.SayHello)
	r.Static("/api/v1/static", "./static")
	r.StaticFile("/favicon.ico", "./static/img/favicon.ico")

	// r.Static("/userdata", "./userdata")

	// Simple group: v1
	{
		// v1 := r.Group("/api/v1")
		// {
		// 	adminR := v1.Group("/admin")
		// 	routeAdmin(adminR)
		// }

		// {
		// 	userR := v1.Group("/users")

		// 	// temp
		// 	userR.POST("/sendsms", controllers.SendVerifySms)
		// }

	}

	return r
}
