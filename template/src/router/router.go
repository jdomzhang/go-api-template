package router

import (
	"{{name}}/src/ctrl"
	"{{name}}/src/middleware"

	"github.com/gin-gonic/gin"
)

// Route contains all the API routes mapping
func Route() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Options)

	r.GET("/", ctrl.SayHello)
	r.GET("/api", ctrl.SayHello)
	r.GET("/api/v1", ctrl.SayHello)
	r.GET("/api/v1/hello", ctrl.SayHello)
	r.Static("/api/v1/static", "./static")
	r.Static("/userdata", "./userdata")
	r.StaticFile("/favicon.ico", "./static/img/favicon.ico")

	v1 := r.Group("/api/v1")
	v1.Use(middleware.Authorization())

	// route Admin
	routeAdmin(v1)
	routeMP(v1)

	// var wechat ctrl.Wechat
	// v1.GET("/wechat/login", wechat.WeChatLogin)
	// v1.POST("/wechat/getuserinfo", wechat.WechatGetUserInfo)
	// v1.POST("/wechat/getphonenumber", ctrl.ShouldBeUser, wechat.WechatGetPhoneNumber)
	// v1.POST("/wechat/storeformid", ctrl.ShouldBeUser, wechat.StoreWechatFormID)

	return r
}
