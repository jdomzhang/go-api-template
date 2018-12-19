package wechat

import (
	"encoding/json"
	"fmt"
	"time"
	"{{name}}/src/models/orm"
	"{{name}}/src/util"

	"{{name}}/src/config"

	"github.com/jdomzhang/resty"
)

var globalAccessToken string

// ForceRefreshGlobalAccessToken will set globalAccessToken
func ForceRefreshGlobalAccessToken() (string, error) {
	var gat orm.WechatGlobalAccessToken
	if err := getGlobalAccessTokenFromWeb(&gat); err != nil {
		return "", err
	}

	// assume gat get token successfully
	if gat.ErrCode == 0 {
		changeGlobalTokenVariable(gat.AccessToken)
		return globalAccessToken, nil
	}

	return "", fmt.Errorf("ErrorCode: %d, ErrorMsg: %s", gat.ErrCode, gat.ErrMsg)
}

// GetOrRefreshGlobalAccessToken will check and refresh wechat access token
func GetOrRefreshGlobalAccessToken() error {
	var gat orm.WechatGlobalAccessToken
	if err := getGlobalAccessTokenFromDB(&gat); err != nil {
		if err := getGlobalAccessTokenFromWeb(&gat); err != nil {
			return err
		}
	}

	// assume gat get token successfully
	if gat.ErrCode == 0 {
		changeGlobalTokenVariable(gat.AccessToken)
	}

	// fmt.Println("access token:", globalAccessToken)
	return nil
}

func getGlobalAccessTokenFromDB(gat *orm.WechatGlobalAccessToken) error {
	var entity orm.WechatGlobalAccessToken
	return entity.GetValidLatest(gat)
}

func getGlobalAccessTokenFromWeb(gat *orm.WechatGlobalAccessToken) error {
	apiBase := fmt.Sprintf(config.All["wechat.getaccesstoken.url"],
		config.All["wechat.app.appid"],
		config.All["wechat.app.secret"])
	resp, _ := resty.R().Get(apiBase)

	// save log
	log := orm.WechatLog{
		Category:       "GlobalToken",
		IO:             "Send",
		URL:            apiBase,
		Method:         resp.Request.Method,
		Request:        util.GetRawRequest(resp.Request),
		Response:       util.GetRawResponse(resp),
		ResponseStatus: int64(resp.StatusCode()),
	}
	log.Create(&log)

	fmt.Println(string(resp.Body()))

	json.Unmarshal(resp.Body(), &gat)

	// save to database
	return gat.Create(&gat)
}

func changeGlobalTokenVariable(token string) {
	if different := token != globalAccessToken; different {
		fmt.Println("changing varialbe to: ", token)
		globalAccessToken = token
	}
}
