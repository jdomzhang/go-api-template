package orm

// Appt is appointment
type Appt struct {
	OrmModel
	Mobile string `json:"mobile"`
	Code   string `json:"code"`
	Age    string `json:"age"`
}

func init() {
	// Migrate the schema
	db.AutoMigrate(&Appt{})
}
