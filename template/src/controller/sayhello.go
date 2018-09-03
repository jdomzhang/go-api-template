package controller

import (
	"github.com/gin-gonic/gin"
)

// SayHello is a simple method to check if the API is working.
func SayHello(c *gin.Context) {
	c.String(200, "Hello API")
}
