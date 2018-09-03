package orm

// Admin is the user to login admin site
type Admin struct {
	OrmModel
	Name      string `json:"name"`
	Password  string `json:"-"`
	Mobile    string `json:"mobile"`
	AdminType string `json:"adminType"` // 管理员类型，client:商家管理员, system:系统管理员
}

func init() {
	// Migrate the schema
	db.AutoMigrate(&Admin{})
}
