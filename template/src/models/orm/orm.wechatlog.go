package orm

// WechatLog saves all the Wechat Pay log
type WechatLog struct {
	OrmModel
	Category       string
	IO             string
	URL            string
	Method         string
	Request        string `sql:"type:text"`
	Response       string `sql:"type:text"`
	ResponseStatus int64
	RefKey         string
	RefName        string
	Ext1           string
	Ext2           string
}

func init() {
	// Migrate the schema
	db.AutoMigrate(&WechatLog{})
}
