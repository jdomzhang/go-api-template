package controllers

import (
	"{{name}}/src/models/biz"

	"github.com/gin-gonic/gin"
)

// GetMyUserInfo will return user object of current login user
func GetMyUserInfo(c *gin.Context) {
	userID := getLoginContext(c).UserID

	if user, err := biz.GetUserByID(userID); err != nil {
		renderError(c, err)
	} else {
		renderJSON(c, user)
	}
}
