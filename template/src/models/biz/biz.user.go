package biz

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"{{name}}/src/models/enc"
	"{{name}}/src/models/orm"
	"{{name}}/src/models/vo"
)

// GetOrCreateUser will get user by given wechat openid, or create if not found
func GetOrCreateUser(user *orm.User, wechatSessionKey, wechatOpenID, wechatUnionID string) error {
	// check wechat session key & openid
	if ok := wechatSessionKey != "" && wechatOpenID != ""; !ok {
		return fmt.Errorf("wechat sessionkey [%s] and openid [%s] should both have value", wechatSessionKey, wechatOpenID)
	}

	var entity orm.User
	if err := entity.FindByWechatOpenID(user, wechatOpenID); err != nil {
		user.WechatSessionKey = wechatSessionKey
		user.WechatOpenID = wechatOpenID
		user.WechatUnionID = wechatUnionID

		return entity.Create(user)
	}

	// update wechat session key
	user.WechatSessionKey = wechatSessionKey
	user.WechatUnionID = wechatUnionID
	return entity.Update(user)
}

// AuthorizeWXPhoneNumber will save mobile phone number to corresponding user
func AuthorizeWXPhoneNumber(rawData vo.WXData, userID uint64) (*orm.User, error) {
	var user orm.User
	if err := user.Get(&user, userID); err != nil {
		return nil, nil
	}

	if jsonStr, err := enc.DecryptWXData(user.WechatSessionKey, rawData.IV, rawData.EncryptedData); err != nil {
		return nil, err
	} else {
		fmt.Println(jsonStr)
		// return jsonStr, nil
		// get phone number from jsonStr
		if vl, err := (enc.Values{}).ParseJSON(jsonStr); err != nil {
			return nil, err
		} else {
			// write to database
			mobile := vl["phoneNumber"]
			user.Mobile = mobile
			user.IsMobileValidated = 1

			if err := user.Update(&user); err != nil {
				return nil, err
			}

			return &user, nil
		}
	}
}

// SaveWXUserInfo will save wechat user info to corresponding user
func SaveWXUserInfo(rawData vo.WXData, userID uint64, openID string) (*orm.User, error) {
	var user orm.User
	if err := user.Get(&user, userID); err != nil {
		// could not find by user id, then find by openid
		if err := user.FindByWechatOpenID(&user, openID); err != nil {
			return nil, err
		}
	}

	if jsonStr, err := enc.DecryptWXData(user.WechatSessionKey, rawData.IV, rawData.EncryptedData); err != nil {
		return nil, err
	} else {
		fmt.Println(jsonStr)
		// return jsonStr, nil
		// convert to wechat user
		var wechatUser orm.WechatUser
		if err := json.Unmarshal([]byte(jsonStr), &wechatUser); err != nil {
			return nil, err
		}

		// set openid
		wechatUser.OpenID = user.WechatOpenID
		wechatUser.UnionID = user.WechatUnionID
		if err := wechatUser.UpdateOrCreateByOpenID(&wechatUser); err != nil {
			return nil, err
		}

		// set nickname and avatar
		user.NickName = wechatUser.NickName
		user.AvatarURL = wechatUser.AvatarURL

		if err := user.Update(&user); err != nil {
			return nil, err
		}

		return &user, nil
		// }
	}
}

// GetUserByID will return user object by given id
func GetUserByID(id uint64) (*orm.User, error) {
	var ormUser orm.User
	if err := ormUser.Get(&ormUser, id); err != nil {
		return nil, err
	}

	return &ormUser, nil
}

// StoreWechatFormID will store user formId for later usage
func StoreWechatFormID(userID uint64, formID string) (*orm.WechatForm, error) {
	// check formID
	if formID == "" {
		return nil, errors.New("formID should not be empty")
	}

	if strings.Index(formID, " ") != -1 {
		return nil, errors.New("formID should not contain space")
	}

	// get user
	var user orm.User
	if err := user.Get(&user, userID); err != nil {
		return nil, err
	}

	form := orm.WechatForm{
		UserID: userID,
		OpenID: user.WechatOpenID,
		FormID: formID,
	}

	return &form, form.Create(&form)
}

// DeleteExpiredFormID will delete the form id that has expired (older than 7 days)
func DeleteExpiredFormID() error {
	var form orm.WechatForm
	return form.DeleteExpiredFormID()
}
