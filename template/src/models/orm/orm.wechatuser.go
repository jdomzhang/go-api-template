package orm

// WechatUser is the wechat user object
type WechatUser struct {
	OrmModel
	AvatarURL string `json:"avatarUrl"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Gender    uint64 `json:"gender"`
	Language  string `json:"language"`
	NickName  string `json:"nickname"`
	Province  string `json:"province"`
	OpenID    string `json:"openid"`
	UnionID   string `json:"unionid"`
}

func init() {
	// Migrate the schema
	db.AutoMigrate(&WechatUser{})
}

func (*WechatUser) UpdateOrCreateByOpenID(v *WechatUser) error {
	// important: Assign() would only work with (*v) which is not a pointer
	return db.Where(WechatUser{OpenID: v.OpenID}).Assign(*v).FirstOrCreate(v).Error
}
