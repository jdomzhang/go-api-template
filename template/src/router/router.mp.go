package router

import (
	"{{name}}/src/ctrl"

	"github.com/gin-gonic/gin"
)

func routeMP(r *gin.RouterGroup) {
	var wechat ctrl.Wechat
	r.GET("/mp/wechat/login", wechat.WeChatLogin)
	r.POST("/mp/wechat/getuserinfo", wechat.WechatGetUserInfo)
	r.POST("/mp/wechat/getphonenumber", ctrl.ShouldBeUser, wechat.WechatGetPhoneNumber)
	r.POST("/mp/wechat/storeformid", ctrl.ShouldBeUser, wechat.StoreWechatFormID)

}
