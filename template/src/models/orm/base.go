package orm

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"
	"{{name}}/src/models/enc"
	"{{name}}/src/models/primitive"

	// this line must be here, it's to reference pq driver
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"{{name}}/src/config"

	"github.com/jinzhu/gorm"
)

// EmptyOrmModel is the empty orm type
type EmptyOrmModel struct{}

// OrmViewModel is the view model
type OrmViewModel struct {
	EmptyOrmModel
	ID        uint64     `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
	Sig       string     `json:"sig"` // sig, the unique field instead of id, no change is allowed
}

type dbViewInterface interface {
	TableName() string
	ViewSQL() string
}

// RecreateView will recreate the given view
func RecreateView(obj dbViewInterface) {
	viewName := obj.TableName()
	viewSQL := obj.ViewSQL()

	db.Exec(fmt.Sprintf("drop view %s cascade", viewName))

	db.Exec(fmt.Sprintf("create or replace view %s as %s", viewName, viewSQL))
}

// OrmModel is the base of orm types
type OrmModel struct {
	EmptyOrmModel
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Sig       string     `json:"sig"` // sig, the unique field instead of id, no change is allowed
}

var (
	db *gorm.DB
)

const (
	// DefaultPageSize is the default page size for get by page
	DefaultPageSize = 10
)

func init() {
	// init db connection
	// and keep it open
	// don't have to open & close everytime
	db = OpenDB()
	fmt.Println("***********DB initialized with LogMode:", config.All["logsql"], "*****************")
}

// InitDbConnection should be called in the beginning of the application
func InitDbConnection() {
	// when this method is called first time
	// the db variable will be set by init() function
}

// OpenDB to get the db connection
func OpenDB() *gorm.DB {
	logSQL, _ := strconv.ParseBool(config.All["logsql"])
	driver := config.All["dbdriver"]
	connectionString := config.All["connectionstring"]

	db, err := gorm.Open(driver, connectionString)

	db.LogMode(logSQL)

	if err != nil {
		log.Println(err.Error())
		panic("failed to connect database")
	}

	return db
}

// BeforeCreate will set order sig before create
func (model *OrmModel) BeforeCreate() error {
	if ev, err := (enc.Values{}).ParseObject(model); err != nil {
		return err
	} else {
		ev["salt"] = fmt.Sprintf("%v", time.Now().UnixNano())
		sig := enc.SignSHA1(ev)
		model.Sig = sig
	}

	return nil
}

// Create orm entity
func (*EmptyOrmModel) Create(obj interface{}) error {
	if rowsAffected := db.Create(obj).RowsAffected; rowsAffected == 0 {
		typeName := reflect.TypeOf(obj)
		return fmt.Errorf("Could not Create the obj[%s]", typeName.String())
	}

	// // save children
	// db.Save(obj)

	return nil
}

// CreateOrUpdate orm with id
func (*EmptyOrmModel) CreateOrUpdate(obj interface{}) error {
	if rowsAffected := db.Save(obj).RowsAffected; rowsAffected == 0 {
		typeName := reflect.TypeOf(obj)
		return fmt.Errorf("CreateOrUpdate obj[%s] has error", typeName.String())
	}

	return nil
}

// Update orm entity
func (*EmptyOrmModel) Update(obj interface{}) error {
	if rowsAffected := db.Save(obj).RowsAffected; rowsAffected == 0 {
		typeName := reflect.TypeOf(obj)
		return fmt.Errorf("Could not Update the obj[%s]", typeName.String())
	}

	return nil
}

// Get orm entity by id
func (*EmptyOrmModel) Get(obj interface{}, id uint64, preloads ...string) error {
	tmpDb := db
	for _, preLoad := range preloads {
		tmpDb = tmpDb.Preload(preLoad)
	}
	if rowsAffected := tmpDb.First(obj, id).RowsAffected; rowsAffected == 0 {
		typeName := reflect.TypeOf(obj)
		return fmt.Errorf("Could not Get the obj[%s] by ID [%v]", typeName.String(), id)
	}

	return nil
}

// Delete the orm entity
func (*EmptyOrmModel) Delete(obj interface{}) error {
	if rowsAffected := db.Delete(obj).RowsAffected; rowsAffected == 0 {
		typeName := reflect.TypeOf(obj)
		return fmt.Errorf("Could not Delete the obj[%s]", typeName.String())
	}

	return nil
}

// DeleteAll the orm entity
func (*EmptyOrmModel) DeleteAll(obj interface{}, where ...interface{}) error {
	return db.Delete(obj, where...).Error
}

// GetAll the orm entities
func (*EmptyOrmModel) GetAll(objs interface{}) error {
	return db.Find(objs).Error
}

// GetAllByPage the orm entities
func (*EmptyOrmModel) GetAllByPage(objs interface{}, pagination *primitive.Pagination, preloads ...string) error {
	var offset uint64
	var limits uint64
	if pagination.RowsPerPage <= -1 {
		offset = 0
		limits = 0
	} else {
		if pagination.RowsPerPage == 0 {
			pagination.RowsPerPage = DefaultPageSize
		}
		limits = uint64(pagination.RowsPerPage)
		offset = uint64((pagination.Page - 1) * limits)
	}

	// get total count first
	var totalCount uint64
	whereSQL := pagination.GetWhereClause()

	if err := db.Where(whereSQL).Find(objs).Count(&totalCount).Error; err != nil {
		return err
	} else {
		pagination.TotalItems = totalCount
	}

	tmpDb := db
	for _, preLoad := range preloads {
		tmpDb = tmpDb.Preload(preLoad)
	}
	// get data later
	tmpDb = tmpDb.Where(whereSQL).Order(pagination.GetOrderClause())
	if offset > 0 {
		tmpDb = tmpDb.Offset(offset)
	}
	if limits > 0 {
		tmpDb = tmpDb.Limit(limits)
	}
	tmpDb = tmpDb.Find(objs)
	if tmpDb.Error != nil {
		return tmpDb.Error
	}

	return nil
}

// Find the orm entities with conditions
func (*EmptyOrmModel) Find(objs interface{}, where ...interface{}) error {
	return db.Find(objs, where...).Error
}

// FindAll will return entities with where conditions
func (*EmptyOrmModel) FindAll(objs interface{}, query interface{}, args ...interface{}) error {
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
