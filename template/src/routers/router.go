package routers

import (
	"{{name}}/src/controllers"
	"{{name}}/src/middleware"

	"github.com/gin-gonic/gin"
)

// Route contains all the API routes mapping
func Route() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Options)

	r.GET("/", controllers.SayHello)
	r.GET("/api", controllers.SayHello)
	r.GET("/api/v1", controllers.SayHello)
	r.GET("/api/v1/hello", controllers.SayHello)
	r.Static("/api/v1/static", "./static")
	r.Static("/userdata", "./userdata")
	r.StaticFile("/favicon.ico", "./static/img/favicon.ico")

	v1 := r.Group("/api/v1")
	v1.Use(middleware.Authorization())

	var wechat controllers.Wechat
	v1.GET("/wechat/login", wechat.WeChatLogin)
	v1.POST("/wechat/getuserinfo", wechat.WechatGetUserInfo)
	v1.POST("/wechat/getphonenumber", controllers.ShouldBeUser, wechat.WechatGetPhoneNumber)
	v1.POST("/wechat/storeformid", controllers.ShouldBeUser, wechat.StoreWechatFormID)

	return r
}
