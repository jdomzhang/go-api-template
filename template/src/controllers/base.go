package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"{{name}}/src/models/auth"
	"{{name}}/src/models/vo"

	"github.com/gin-gonic/gin"
)

const bearerPrefix = "Bearer "
const authorizationHeader = "Authorization"
const loginContextKey = "LoginContextKey"

// GetTokenFromHeaderOrQuery will get token from header or query string
func GetTokenFromHeaderOrQuery(c *gin.Context) string {
	token := GetHeaderToken(c)
	if token == "" {
		q, _ := url.ParseQuery(c.Request.URL.RawQuery)
		token = q.Get("jwt")
	}

	return token
}

// GetHeaderToken will get authorization header
func GetHeaderToken(c *gin.Context) string {
	value := c.GetHeader(authorizationHeader)
	if index := strings.Index(value, bearerPrefix); index == 0 {
		value = value[index+len(bearerPrefix):]
	}
	return value
}

// SetHeaderToken will set authorization header with given token
func SetHeaderToken(c *gin.Context, token string) {
	value := bearerPrefix + token
	c.Header(authorizationHeader, value)
}

// LoginContext contains current API user info
type LoginContext struct {
	UserID      uint64
	OpenID      string
	Visitor     string
	IsValid     bool
	IsValidUser bool
	IsVisitor   bool
}

// ShouldBeUser will check current login context, and make sure it's an authorized user
func ShouldBeUser(c *gin.Context) {
	loginContext := getLoginContext(c)
	if ok := loginContext.IsValidUser; !ok {
		c.AbortWithError(http.StatusUnauthorized, errors.New("current user has no proper right to access this api"))
	}
}

// ShouldBeVisitor will check current login context, and make sure it's a visitor (has token only)
func ShouldBeVisitor(c *gin.Context) {
	loginContext := getLoginContext(c)
	if ok := loginContext.IsVisitor; !ok {
		c.AbortWithError(http.StatusUnauthorized, errors.New("current user has no proper right to access this api"))
	}
}

func getLoginContext(c *gin.Context) LoginContext {

	if value, ok := c.Keys[loginContextKey]; ok {
		return value.(LoginContext)
	}

	// not existing, so set payload
	// commented this, since passing url token scenario has been removed
	// token := GetTokenFromHeaderOrQuery(c)
	token := GetHeaderToken(c)
	if token == "" {
		// generate token if it's empty
		token = RefreshTokenForUserOrVisitor("")
	}

	ok, payload := auth.Check(token)
	loginContext := LoginContext{
		IsValid:       ok,
		UserID:        payload.UserID,
		AdminID:       payload.AdminID,
		OpenID:        payload.WechatOpenID,
		IsSystemAdmin: payload.IsSystemAdmin,
		IsOperator:    payload.IsOperator,
		IsSalesPerson: payload.IsSalesPerson,
		IsValidAdmin:  payload.AdminID > 0 && ok,
		IsValidUser:   payload.UserID > 0 && ok,
		Visitor:       payload.Visitor,
		IsVisitor:     payload.Visitor != "",
	}

	c.Set(loginContextKey, loginContext)

	return loginContext
}

// RefreshTokenForUserOrVisitor will generate a new token
func RefreshTokenForUserOrVisitor(token string) string {
	// will always refresh for visitor
	// will convert an expired user to a visitor
	if len(token) > 0 {
		if ok, payload := auth.Check(token); ok {
			return payload.Gen()
		}
	}

	return auth.GenVisitorJwt()
}

func renderErrorMessage(c *gin.Context, message string, a ...interface{}) {
	if a != nil {
		message = fmt.Sprintf(message, a...)
	}
	c.IndentedJSON(http.StatusBadRequest, &vo.Error{Error: message})

	// log
	fmt.Println("API Error:", message)

	// stop running other handler in the chain
	c.Abort()
}

func renderUnauthorizedError(c *gin.Context) {
	c.IndentedJSON(http.StatusUnauthorized, &vo.Error{Error: "user is not authorized"})
	// renderErrorMessage(c, "current user has no permission to call this method [%s]", util.GetCurrentMethodName())

	// stop running other handler in the chain
	c.Abort()
}

func renderSuccessMessage(c *gin.Context, message string, a ...interface{}) {
	if a != nil {
		message = fmt.Sprintf(message, a...)
	}
	c.IndentedJSON(http.StatusOK, &vo.Success{Success: message})
}

func renderError(c *gin.Context, err error) {
	renderErrorMessage(c, err.Error())

	// stop running other handler in the chain
	c.Abort()
}

func renderJSON(c *gin.Context, obj interface{}) {
	c.IndentedJSON(http.StatusOK, obj)
}

func renderString(c *gin.Context, format string, values ...interface{}) {
	c.String(http.StatusOK, format, values...)
}
