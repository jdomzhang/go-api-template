package wechat

import (
	"encoding/json"
	"fmt"
	"{{name}}/src/config"
	"{{name}}/src/models/orm"
	"{{name}}/src/util"

	"github.com/jdomzhang/resty"
)

// LoginResponse contains the response data when call sns/jscode2session.
// for more detail, please check https://developers.weixin.qq.com/miniprogram/dev/api/open-api/login/code2Session.html
type LoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// JSCode2Session will get session value with given code
func JSCode2Session(result *LoginResponse, code string) error {
	appid := config.All["wechat.app.appid"]
	secret := config.All["wechat.app.secret"]

	// construct url
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		appid, secret, code)

	resp, err := resty.R().Get(url)

	// save log
	log := orm.WechatLog{
		Category:       "JSCode2Session",
		IO:             "Send",
		URL:            resp.Request.URL,
		Method:         resp.Request.Method,
		Request:        util.GetRawRequest(resp.Request),
		Response:       util.GetRawResponse(resp),
		ResponseStatus: int64(resp.StatusCode()),
	}
	log.Create(&log)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(resp.Body(), result); err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return fmt.Errorf("Error: %v", result)
	}

	return nil
}
