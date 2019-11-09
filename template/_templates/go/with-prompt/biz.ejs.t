---
to: src/models/biz/biz.<%=h.changeCase.lowerCase(name)%>.go
---
package biz

import (
	"{{name}}/src/models/orm"
	"{{name}}/src/models/primitive"
)

// <%=name%> is a biz processor
type <%=name%> struct{}

var biz<%=name%> <%=name%>;

// Get will get an ormObj
func (obj *<%=name%>) Get(ormObj *orm.<%=name%>, id uint64) error {
	return ormObj.Get(ormObj, id)
}

// GetAllByPage will return data by page
func (obj *<%=name%>) GetAllByPage(list *[]orm.<%=name%>, pagination *primitive.Pagination) error {
	var ormEmpty orm.EmptyOrmModel
	return ormEmpty.GetAllByPage(list, pagination)
}

// Create will create an ormObj
func (obj *<%=name%>) Create(ormObj *orm.<%=name%>) error {
	if err := obj.validateAdd(ormObj); err != nil {
		return err
	}

	if err := ormObj.Create(ormObj); err != nil {
		return err
	}

	return obj.Get(ormObj, ormObj.ID)
}

// Update will update an ormObj
func (obj *<%=name%>) Update(ormObj *orm.<%=name%>) error {
	if err := obj.validateUpdate(ormObj); err != nil {
		return err
	}

	if err := ormObj.Update(ormObj); err != nil {
		return err
	}

	return obj.Get(ormObj, ormObj.ID)
}

// Delete will delete an ormObj
func (obj *<%=name%>) Delete(id uint64) error {
	ormObj := &orm.<%=name%>{}
	if err := obj.Get(ormObj, id); err != nil {
		return err
	}

	if err := obj.validateDelete(ormObj); err != nil {
		return err
	}

	return ormObj.Delete(ormObj)
}

/*
 handle business logic here
*/

func (obj *<%=name%>) validateAdd(ormObj *orm.<%=name%>) error {
	return obj.validate(ormObj)
}

func (obj *<%=name%>) validateUpdate(ormObj *orm.<%=name%>) error {
	return obj.validate(ormObj)
}

func (obj *<%=name%>) validateDelete(ormObj *orm.<%=name%>) error {
	return nil
}

func (obj *<%=name%>) validate(ormObj *orm.<%=name%>) error {
	return nil
}
