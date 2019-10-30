package ctrl

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SayHello is a simple method to check if the API is working.
func SayHello(c *gin.Context) {
	c.String(200, fmt.Sprintf("[%s]Hello API\n", time.Now().Format(time.RFC3339)))
}
