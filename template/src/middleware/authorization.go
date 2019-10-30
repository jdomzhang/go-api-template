package middleware

import (
	"{{name}}/src/ctrl"

	"github.com/gin-gonic/gin"
)

// Authorization returns the middle ware to refresh token
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := ctrl.GetHeaderToken(c)
		newToken := ctrl.RefreshTokenForUserOrVisitor(token)
		ctrl.SetHeaderToken(c, newToken)
	}
}
