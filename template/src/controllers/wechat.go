package controllers

import (
	"fmt"

	"{{name}}/src/models/vo"

	"{{name}}/src/models/auth"
	"{{name}}/src/models/biz"
	"{{name}}/src/models/orm"

	"github.com/gin-gonic/gin"
)

// Wechat controller
type Wechat struct{}

// WeChatLogin will return the openid by given code
// @Summary 小程序登录
// @Description 小程序登录
// @Description 更多信息，请参考 https://developers.weixin.qq.com/ebook?action=get_post_info&volumn=1&lang=zh_CN&book=miniprogram&docid=000cc48f96c5989b0086ddc7e56c0a
// @Accept json
// @Produce json
// @Param code query string true "code"
// @Success 200 {object} orm.User
// @Failure 400 {object} vo.Error
// @Router /wechat/login [get]
// @Tags weapp,login
func (*Wechat) WeChatLogin(c *gin.Context) {
	// appid := config.All["wechat.app.appid"]
	// secret := config.All["wechat.app.secret"]
	code := c.Query("code")

	if ok := code != ""; !ok {
		renderErrorMessage(c, "code is required in query string")
		return
	}

	var user orm.User
	if err := biz.LoginUserWithWechatJSCode(&user, code); err != nil {
		renderError(c, err)
		return
	}

	// set token
	token := auth.GenUserJwt(&user)
	SetHeaderToken(c, token)

	renderJSON(c, &user)
}

// WechatGetPhoneNumber will get phone number from encrypted wechat data
// @Summary 获取授权的微信用户手机号
// @Description 获取授权的微信用户手机号
// @Description 更多信息，请参考 https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/getPhoneNumber.html
// @Accept json
// @Produce json
// @Param data body vo.WXData true "HTTP request body"
// @Success 200 {object} orm.User
// @Failure 400 {object} vo.Error
// @Failure 401 string string "Unauthorized"
// @Security ApiKeyAuth
// @Router /wechat/getphonenumber [post]
// @Tags weapp
func (*Wechat) WechatGetPhoneNumber(c *gin.Context) {
	var wxData vo.WXData
	if err := c.ShouldBind(&wxData); err != nil {
		renderError(c, err)
		return
	}

	userID := getLoginContext(c).UserID
	fmt.Printf("current user is [%d]\n", userID)

	if user, err := biz.AuthorizeWXPhoneNumber(wxData, userID); err != nil {
		renderError(c, err)
	} else {
		renderJSON(c, user)
	}
}

// WechatGetUserInfo will get user info from encrypted wechat data
// @Summary 获取授权的微信用户信息
// @Description 获取授权的微信用户信息
// @Description 更多信息，请参考 https://developers.weixin.qq.com/miniprogram/dev/api/wx.getUserInfo.html
// @Accept json
// @Produce json
// @Param data body vo.WXData true "HTTP request body"
// @Success 200 {object} orm.User
// @Failure 400 {object} vo.Error
// @Failure 401 string string "Unauthorized"
// @Security ApiKeyAuth
// @Router /wechat/getuserinfo [post]
// @Tags weapp
func (*Wechat) WechatGetUserInfo(c *gin.Context) {
	var wxData vo.WXData
	if err := c.ShouldBind(&wxData); err != nil {
		renderError(c, err)
		return
	}

	lc := getLoginContext(c)
	userID := lc.UserID
	openID := wxData.WechatOpenID
	if openID == "" {
		openID = lc.OpenID
	}
	fmt.Printf("current user is [%d]\n", userID)

	if user, err := biz.SaveWXUserInfo(wxData, userID, openID); err != nil {
		renderError(c, err)
	} else {
		// set token, in case the user object is changed
		token := auth.GenUserJwt(user)
		SetHeaderToken(c, token)

		renderJSON(c, user)
	}
}

// StoreWechatFormID will store the formid for later usage
// @Summary 新增微信用户的FormID
// @Description 新增微信用户的FormID
// @Accept json
// @Produce json
// @Param data body vo.FormID true "HTTP request body"
// @Success 200 {object} orm.WechatForm
// @Failure 400 {object} vo.Error
// @Failure 401 string string "Unauthorized"
// @Security ApiKeyAuth
// @Router /wechat/storeformid [post]
// @Tags weapp
func (*Wechat) StoreWechatFormID(c *gin.Context) {
	userID := getLoginContext(c).UserID

	var data vo.FormID
	if err := c.ShouldBind(&data); err != nil {
		renderError(c, err)
		return
	}

	if form, err := biz.StoreWechatFormID(userID, data.FormID); err != nil {
		renderError(c, err)
	} else {
		renderJSON(c, form)
	}
}
