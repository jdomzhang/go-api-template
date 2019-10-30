package ctrl

import (
	"strconv"
	"{{name}}/src/models/biz"
	"{{name}}/src/models/orm"
	"{{name}}/src/models/primitive"
	"{{name}}/src/models/vo"
	"{{name}}/src/util"

	"github.com/gin-gonic/gin"
)

// User is a controller
type User struct{}

// GetByID is gin controller
func (*User) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var ormObj orm.User
	var bizObj biz.User
	if err := bizObj.Get(&ormObj, id); err != nil {
		renderError(c, err)
		return
	}

	renderJSON(c, &ormObj)
}

// Create is an action
func (*User) Create(c *gin.Context) {
	// create(c, &orm.User{}, &biz.User{})
	// validate
	var ormObj orm.User
	if err := c.ShouldBind(&ormObj); err != nil {
		renderError(c, err)
		return
	}

	// add
	var bizObj biz.User
	if err := bizObj.Create(&ormObj); err != nil {
		renderError(c, err)
		return
	}

	renderJSON(c, &ormObj)
}

// Update is gin controller
func (*User) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		renderError(c, err)
		return
	}

	var bizObj biz.User

	var origin orm.User
	if err := bizObj.Get(&origin, id); err != nil {
		renderError(c, err)
		return
	}

	var filter vo.User
	if err := util.BindWithFilter(c, &filter, &origin); err != nil {
		renderError(c, err)
		return
	}

	if err := bizObj.Update(&origin); err != nil {
		renderError(c, err)
	} else {
		renderJSON(c, &origin)
	}
}

// Delete is gin controller
func (*User) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		renderError(c, err)
		return
	}

	var bizObj biz.User

	if err := bizObj.Delete(id); err != nil {
		renderError(c, err)
		return
	}

	renderSuccessMessage(c, "done")
}

// GetAllByPage return all by page
func (*User) GetAllByPage(c *gin.Context) {
	var pagination primitive.Pagination
	if err := c.ShouldBind(&pagination); err != nil {
		renderError(c, err)
		return
	}

	var bizObj biz.User

	var list []orm.User
	if err := bizObj.GetAllByPage(&list, &pagination); err != nil {
		renderError(c, err)
		return
	}

	dataByPage := vo.DataByPage{
		Pagination: pagination,
		Items:      &list,
	}
	renderJSON(c, &dataByPage)
}
