package orm

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	// this line must be here, it's to reference pq driver
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"../../config"
	"github.com/jinzhu/gorm"
)

// Model is the base of orm types
type OrmModel struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

var (
	db *gorm.DB
)

func init() {
	// init db connection
	// and keep it open
	// don't have to open & close everytime
	db = OpenDB()
}

// InitDbConnection should be called in the beginning of the application
func InitDbConnection() {
	// when this method is called first time
	// the db variable will be set by init() function
	fmt.Println("***********DB initialized with LogMode:", config.All["logsql"], "*****************")

}

// OpenDB to get the db connection
func OpenDB() *gorm.DB {
	logSQL, _ := strconv.ParseBool(config.All["logsql"])
	driver := config.All["dbdriver"]
	connectionString := config.All["connectionstring"]

	db, err := gorm.Open(driver, connectionString)

	db.LogMode(logSQL)

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	return db
}

// Create orm entity
func (model *OrmModel) Create(obj interface{}) error {
	if rowsAffected := db.Create(obj).RowsAffected; rowsAffected == 0 {
		typeName := reflect.TypeOf(obj)
		return fmt.Errorf("Could not Create the obj[%s]", typeName.String())
	}

	// // save children
	// db.Save(obj)

	return nil
}

// CreateOrUpdate orm with id
func (model *OrmModel) CreateOrUpdate(obj interface{}) error {
	// db.Save(obj)

	if rowsAffected := db.Save(obj).RowsAffected; rowsAffected == 0 {
		typeName := reflect.TypeOf(obj)
		return fmt.Errorf("CreateOrUpdate obj[%s] has error", typeName.String())
	}

	return nil
}

// Update orm entity
func (model *OrmModel) Update(obj interface{}) error {
	if rowsAffected := db.Save(obj).RowsAffected; rowsAffected == 0 {
		typeName := reflect.TypeOf(obj)
		return fmt.Errorf("Could not Update the obj[%s]", typeName.String())
	}

	return nil
}

// Get orm entity by id
func (model *OrmModel) Get(obj interface{}, id uint64) error {
	if rowsAffected := db.First(obj, id).RowsAffected; rowsAffected == 0 {
		typeName := reflect.TypeOf(obj)
		return fmt.Errorf("Could not Get the obj[%s] by ID [%v]", typeName.String(), id)
	}

	return nil
}

// Delete the orm entity
func (model *OrmModel) Delete(obj interface{}) error {
	if rowsAffected := db.Delete(obj).RowsAffected; rowsAffected == 0 {
		typeName := reflect.TypeOf(obj)
		return fmt.Errorf("Could not Delete the obj[%s]", typeName.String())
	}

	return nil
}

// DeleteAll the orm entity
func (model *OrmModel) DeleteAll(obj interface{}, where ...interface{}) error {
	return db.Delete(obj, where...).Error
}

// GetAll the orm entities
func (model *OrmModel) GetAll(objs interface{}) error {
	return db.Find(objs).Error
}

// Find the orm entities with conditions
func (model *OrmModel) Find(objs interface{}, where ...interface{}) error {
	return db.Find(objs, where...).Error
}

// FindAll will return entities with where conditions
func (model *OrmModel) FindAll(objs interface{}, query interface{}, args ...interface{}) error {
	return db.Where(query, args...).Find(objs).Error
}

// EnsureSequence will enable table sequence > the max id
// in case manual inserted sql with ID specified without increasing seq
func EnsureSequence(tableName string) {
	var currID, nextSeq struct {
		Value int64
	}

	db.Raw(fmt.Sprintf("SELECT max(id) as value from %s", tableName)).Scan(&currID)

	db.Raw(fmt.Sprintf("SELECT NEXTVAL('%s_id_seq') as value", tableName)).Scan(&nextSeq)

	if currID.Value > nextSeq.Value {
		db.Exec(fmt.Sprintf("ALTER SEQUENCE %s_id_seq restart with %d", tableName, currID.Value+100))
	}
}
