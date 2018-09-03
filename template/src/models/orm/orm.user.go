package orm

import (
	"fmt"
)

type User struct {
	OrmModel
	NickName         string `json:"nickName"`
	AvatarURL        string `json:"avatarUrl"`
	Mobile           string `json:"mobile"`
	IsMobileValid    bool   `sql:"default:false" json:"isMobileValid"`
	WechatOpenID     string `json:"wechatOpenID"`
	WechatSessionKey string `json:"-"`
}

func init() {
	// Migrate the schema
	db.AutoMigrate(&User{})
}

// FindByWechatOpenID will return user by given wechat openid
func (*User) FindByWechatOpenID(user *User, wechatOpenID string) error {
	if rowsAffected := db.Model(user).Where("wechat_open_id = ?", wechatOpenID).First(user).RowsAffected; rowsAffected == 0 {
		return fmt.Errorf("could not find user by given openid [%s]", wechatOpenID)
	}
	return nil
}
