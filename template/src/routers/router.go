package routers

import (
	"{{name}}/src/controllers"
	"{{name}}/src/middleware"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Options)

	r.GET("/", controllers.SayHello)
	r.GET("/api", controllers.SayHello)
	r.GET("/api/v1", controllers.SayHello)
	r.GET("/api/v1/hello", controllers.SayHello)
	r.Static("/api/v1/static", "./static")
	r.StaticFile("/favicon.ico", "./static/img/favicon.ico")

	return r
}
