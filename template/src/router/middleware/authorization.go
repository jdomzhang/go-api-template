package middleware

import (
	"../../controllers"
	"github.com/gin-gonic/gin"
)

// const authorizationHeader = "Authorization"

// Authorization returns the middle ware to refresh token
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := controllers.GetHeaderToken(c)
		newToken := controllers.RefreshTokenForUserOrVisitor(token)
		controllers.SetHeaderToken(c, newToken)
	}
}
