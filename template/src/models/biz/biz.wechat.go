package biz

import (
	"{{name}}/src/models/orm"
	"{{name}}/src/models/wechat"
)

// LoginUserWithWechatJSCode will login wechat user with given code
func LoginUserWithWechatJSCode(user *orm.User, code string) error {
	// var v enc.Values
	var result wechat.LoginResponse
	if err := wechat.JSCode2Session(&result, code); err != nil {
		return err
	}

	return GetOrCreateUser(user, result.SessionKey, result.OpenID, result.UnionID)
}
