---
to: src/ctrl/ctrl.<%=h.changeCase.lowerCase(name)%>.go
---
package ctrl

import (
	"{{name}}/src/models/biz"
	"{{name}}/src/models/orm"
	"{{name}}/src/models/primitive"
	"{{name}}/src/models/vo"
	"{{name}}/src/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

// <%=name%> is a controller
type <%=name%> struct{}

// GetByID is gin controller
func (*<%=name%>) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var ormObj orm.<%=name%>
	var bizObj biz.<%=name%>
	if err := bizObj.Get(&ormObj, id); err != nil {
		renderError(c, err)
		return
	}

	renderJSON(c, &ormObj)
}

// Create is an action
func (*<%=name%>) Create(c *gin.Context) {
	// create(c, &orm.<%=name%>{}, &biz.<%=name%>{})
	// validate
	var ormObj orm.<%=name%>
	if err := c.ShouldBind(&ormObj); err != nil {
		renderError(c, err)
		return
	}

	// add
	var bizObj biz.<%=name%>
	if err := bizObj.Create(&ormObj); err != nil {
		renderError(c, err)
		return
	}

	renderJSON(c, &ormObj)
}

// Update is gin controller
func (*<%=name%>) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		renderError(c, err)
		return
	}

	var bizObj biz.<%=name%>

	var origin orm.<%=name%>
	if err := bizObj.Get(&origin, id); err != nil {
		renderError(c, err)
		return
	}

	var filter vo.<%=name%>
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
func (*<%=name%>) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		renderError(c, err)
		return
	}

	var bizObj biz.<%=name%>

	if err := bizObj.Delete(id); err != nil {
		renderError(c, err)
		return
	}

	renderSuccessMessage(c, "done")
}

// GetAllByPage return all by page
func (*<%=name%>) GetAllByPage(c *gin.Context) {
	var pagination primitive.Pagination
	if err := c.ShouldBind(&pagination); err != nil {
		renderError(c, err)
		return
	}

	var bizObj biz.<%=name%>

	var list []orm.<%=name%>
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
