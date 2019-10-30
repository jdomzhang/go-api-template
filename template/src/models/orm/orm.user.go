package orm

import (
	"fmt"
)

// User is for front-end, may not use in Wechat Mini Program project
type User struct {
	OrmModel
	FullName          string `json:"fullName,omitempty"`
	Email             string `json:"email,omitempty"`
	Mobile            string `json:"mobile,omitempty"`
	NickName          string `json:"nickName,omitempty"`
	AvatarURL         string `json:"avatarUrl,omitempty"`
	IsMobileValidated bool   `json:"isMobileValidated,omitempty" sql:"default:0"`
	WechatOpenID      string `json:"wechatOpenID,omitempty"`
	WechatUnionID     string `json:"wechatUnionID,omitempty"`
}

func init() {
	// Migrate the schema
	db.AutoMigrate(&User{})
}

// ExistByOpenID will get the first by openID
func (*User) ExistByOpenID(openID string) (*User, error) {
	var value User
	return &value, db.Model(&User{}).Where("wechat_open_id = ?", openID).First(&value).Error
}

// FindByMobileAndPassword will return user by given mobile and password
func (*User) FindByMobileAndPassword(user *User, mobile, password string) error {
	if rowsAffected := db.Model(&User{}).Where("mobile = ? and password = ?", mobile, password).First(user).RowsAffected; rowsAffected == 0 {
		return fmt.Errorf("mobile or password is wrong")
	}
	return nil
}

// FindByWechatOpenID will return user by given wechat openid
func (*User) FindByWechatOpenID(user *User, wechatOpenID string) error {
	if rowsAffected := db.Model(user).Where("wechat_open_id = ?", wechatOpenID).First(user).RowsAffected; rowsAffected == 0 {
		return fmt.Errorf("could not find user by given openid [%s]", wechatOpenID)
	}
	return nil
}
