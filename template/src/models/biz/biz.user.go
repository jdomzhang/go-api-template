package biz

import (
	"{{name}}/src/models/orm"
	"{{name}}/src/models/primitive"
	"{{name}}/src/models/wechat"
)

// User is a biz processor
type User struct{}

// default biz object
var bizUser User

func (obj *User) get(ormObj *orm.User, id uint64, preloads ...string) error {
	// do not change below line, it is for simple get
	// you may change it in Get method instead
	return ormObj.Get(ormObj, id, preloads...)
}

// Get will get an ormObj
func (obj *User) Get(ormObj *orm.User, id uint64) error {
	return ormObj.Get(ormObj, id)
}

// GetAllByPage will return data by page
func (obj *User) GetAllByPage(list *[]orm.User, pagination *primitive.Pagination) error {
	var ormEmpty orm.EmptyOrmModel
	return ormEmpty.GetAllByPage(list, pagination)
}

// Create will create an ormObj
func (obj *User) Create(ormObj *orm.User) error {
	if err := obj.validateAdd(ormObj); err != nil {
		return err
	}

	if err := ormObj.Create(ormObj); err != nil {
		return err
	}

	return obj.Get(ormObj, ormObj.ID)
}

// Update will update an ormObj
func (obj *User) Update(ormObj *orm.User) error {
	if err := obj.validateUpdate(ormObj); err != nil {
		return err
	}

	if err := ormObj.Update(ormObj); err != nil {
		return err
	}

	return obj.Get(ormObj, ormObj.ID)
}

// Delete will delete an ormObj
func (obj *User) Delete(id uint64) error {
	ormObj := &orm.User{}
	if err := obj.Get(ormObj, id); err != nil {
		return err
	}

	if err := obj.validateDelete(ormObj); err != nil {
		return err
	}

	return ormObj.Delete(ormObj)
}

// LoginUserWithWechatJSCode will login wechat user with given code
func (*User) LoginUserWithWechatJSCode(code string) (*orm.User, *orm.WechatUser, error) {
	// var v enc.Values
	var result wechat.LoginResponse
	if err := wechat.JSCode2Session(&result, code); err != nil {
		return nil, nil, err
	}

	return bizUser.GetRealOrFakeUser(&result)

	// return GetOrCreateUser(user, result.SessionKey, result.OpenID, result.UnionID)
}

func (*User) GetRealOrFakeUser(result *wechat.LoginResponse) (*orm.User, *orm.WechatUser, error) {
	user := orm.User{
		WechatOpenID: result.OpenID,
	}
	var wechatUser orm.WechatUser

	// 1. get or create wechat user by openid
	var ormWechatUser orm.WechatUser
	if retVal, err := ormWechatUser.ExistByOpenID(result.OpenID); err != nil {
		// not existing, then create
		wechatUser = orm.WechatUser{
			OpenID:     result.OpenID,
			UnionID:    result.UnionID,
			SessionKey: result.SessionKey,
		}

		return &user, &wechatUser, ormWechatUser.Create(&wechatUser)
	} else {
		wechatUser = *retVal
		// update wechat user session key
		wechatUser.SessionKey = result.SessionKey
		if err := ormWechatUser.Update(&wechatUser); err != nil {
			return nil, nil, err
		}
	}

	// 2. existing wechat user, then check if existing user
	var ormUser orm.User
	if retVal, err := ormUser.ExistByOpenID(result.OpenID); err == nil {
		// existing
		user = *retVal
	}

	// return, (userID could be 0)
	return &user, &wechatUser, nil
}

/*
 handle business logic here
*/

func (obj *User) validateAdd(ormObj *orm.User) error {
	return obj.validate(ormObj)
}

func (obj *User) validateUpdate(ormObj *orm.User) error {
	return obj.validate(ormObj)
}

func (obj *User) validateDelete(ormObj *orm.User) error {
	return nil
}

func (obj *User) validate(ormObj *orm.User) error {
	return nil
}
