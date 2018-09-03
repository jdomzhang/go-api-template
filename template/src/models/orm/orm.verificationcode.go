package orm

import (
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// VerificationCode is to verify code through email or mobile
type VerificationCode struct {
	OrmModel
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
	Code   string `json:"code"`
}

func init() {
	// Migrate the schema
	db.AutoMigrate(&VerificationCode{})
}

// DoesSmsCodeMatch will check if the sms code match
func (obj *VerificationCode) DoesSmsCodeMatch(emailOrMobile, code string) bool {
	var vcode VerificationCode
	emailOrMobile = strings.ToLower(emailOrMobile)

	t := time.Now().Add(-20 * time.Minute)

	if rowsAffected := db.Model(&vcode).Where("lower(mobile) = ? and created_at > ?", emailOrMobile, t).Last(&vcode).RowsAffected; rowsAffected == 0 {
		return false
	}
	logrus.Infof("db: %s, input: %s", vcode.Code, code)
	return vcode.Code == code
}
