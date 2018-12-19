package orm

import "fmt"

// WechatGlobalAccessToken is the entity to store wechat global access token
// see https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140183
type WechatGlobalAccessToken struct {
	OrmModel
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   uint64 `json:"expires_in"`   // 凭证有效时间，单位：秒
	ErrCode     uint64 `json:"errcode"`      // -1: 系统繁忙, 0: 请求成功, 40001: AppSecret错误, 40002: 请确保grant_type字段值为client_credential, 40013: 错误，详见errmsg
	ErrMsg      string `json:"errmsg"`       // invalid appid
	// Ticket      string `json:"ticket"`       // invalid appid
}

func init() {
	// Migrate the schema
	db.AutoMigrate(&WechatGlobalAccessToken{})
}

// GetValidLatest returns valid latest entity
func (*WechatGlobalAccessToken) GetValidLatest(entity *WechatGlobalAccessToken) error {
	rowsAffected := db.Where("err_code = 0 and created_at + cast(expires_in || ' seconds' as interval) > now()").Order("id desc").First(&entity).RowsAffected

	if rowsAffected == 0 {
		return fmt.Errorf("Could not find the WechatGlobalAccessToken")
	}

	return nil
}
