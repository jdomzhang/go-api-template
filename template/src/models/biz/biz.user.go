package biz

import (
	"encoding/json"
	"fmt"

	"../../models/enc"
	"../../models/orm"
	"../../models/vo"
)

// LoginUser will log user on with given name and password
func LoginUser(wechatOpenID, wechatSessionKey string) (*orm.User, error) {
	var user orm.User
	if err := user.FindByWechatOpenID(&user, wechatOpenID); err != nil {
		user = orm.User{
			WechatOpenID:     wechatOpenID,
			WechatSessionKey: wechatSessionKey,
		}

		if err := user.Create(&user); err != nil {
			return nil, err
		}

		return &user, nil
	}

	// update wechat session key
	user.WechatSessionKey = wechatSessionKey
	if err := user.Update(&user); err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

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
			user.IsMobileValid = true

			if err := user.Update(&user); err != nil {
				return nil, err
			}

			return &user, nil
		}
	}
}

func SaveWXUserInfo(rawData vo.WXData, userID uint64) (*orm.User, error) {
	var user orm.User
	if err := user.Get(&user, userID); err != nil {
		return nil, nil
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

// // CheckUser will find user by email or mobile
// func CheckUser(name string) (orm.User, error) {
// 	var user orm.User
// 	if err := user.FindByEmailOrMobile(&user, name); err == nil {
// 		return user, nil
// 	}
// 	return user, fmt.Errorf("The user doesn't exist")
// }

// // VerifyMobileCode will verify the given mobile and verification code
// func VerifyMobileCode(mobile, code string) (orm.User, error) {
// 	var user orm.User
// 	var verificationCode orm.VerificationCode
// 	if ok := verificationCode.DoesSmsCodeMatch(mobile, code); ok {
// 		if err := user.FindByEmailOrMobile(&user, mobile); err != nil {
// 			user.Name = mobile
// 			user.PwdHash = util.MD5(util.GetRandomString(8))
// 			user.Mobile = mobile
// 			user.Avatar = "/static/img/avatar.jpg"
// 			user.Create(&user)
// 		}

// 		user.IsMobileValid = true
// 		user.Update(&user)
// 	} else {
// 		return user, fmt.Errorf("The code doesn't match")
// 	}

// 	return user, nil
// }

// // LoadOrdersByUser will return all user's orders
// func LoadOrdersByUser(orders *[]orm.Order, clientID uint64, userID uint64) error {
// 	var o orm.Order
// 	return o.LoadAllByClientAndUser(orders, clientID, userID)
// }

// // LoadOneOrderByUserAndID will return one order by given user id and order id
// func LoadOneOrderByUserAndID(o *orm.Order, userID, id uint64) error {
// 	// var o orm.Order
// 	return o.LoadOneOrderByUserAndID(o, userID, id)
// }

// // GetOrCreateClientUser will check if the client user existing, if no, create.
// // then return one client user
// func GetOrCreateClientUser(clientUser *orm.ClientUser, clientID uint64, userID uint64) error {
// 	if err := clientUser.LoadOne(clientUser, clientID, userID); err != nil {
// 		// err means the client user doesn't exist
// 		// so create one

// 		// check client existing
// 		var c orm.Client
// 		if err := c.Get(&c, clientID); err != nil {
// 			return err
// 		}

// 		// check user existing
// 		var u orm.User
// 		if err := u.Get(&u, userID); err != nil {
// 			return err
// 		}

// 		// create client user
// 		clientUser.ClientID = clientID
// 		clientUser.UserID = userID

// 		if err := clientUser.Create(clientUser); err != nil {
// 			return err
// 		}

// 		// load client user
// 		return clientUser.LoadOne(clientUser, clientID, userID)
// 	}

// 	return nil
// }
