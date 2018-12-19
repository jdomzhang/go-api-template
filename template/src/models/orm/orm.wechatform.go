package orm

import "errors"

// WechatForm stores user's lots of formId
type WechatForm struct {
	OrmModel
	UserID uint64 `json:"userID"`
	OpenID string `json:"openID"`
	FormID string `json:"formID"`
}

func init() {
	// Migrate the schema
	db.AutoMigrate(&WechatForm{})
}

// GetValidOneByUserID will return the oldest valid form id by given user id
func (*WechatForm) GetValidOneByUserID(form *WechatForm, userID uint64) error {
	// form id only valid for 7 days
	if rowsAffected := db.Where("user_id = ? and created_at + interval '7 days' > now()", userID).Order("id").First(&form).RowsAffected; rowsAffected == 0 {
		return errors.New("no form id formed")
	}

	return nil
}

// DeleteExpiredFormID will delete records if it's older than 7 days
func (*WechatForm) DeleteExpiredFormID() error {
	return db.Exec("delete from wechat_forms where created_at + interval '7 days' < now()").Error
}
