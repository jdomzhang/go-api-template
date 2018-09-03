package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"../models/auth"
	"github.com/gin-gonic/gin"
)

// type Controller struct{}

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
	return c.GetHeader(authorizationHeader)
}

// SetHeaderToken will set authorization header with given token
func SetHeaderToken(c *gin.Context, token string) {
	c.Writer.Header().Set(authorizationHeader, token)
}

type LoginContext struct {
	UserID    uint64
	AdminID   uint64
	AdminType string
	// OpenID              string
	// EncryptedSessionKey string
	// ClientID     uint64
	// ShopID       uint64
	Visitor      string
	IsValid      bool
	IsValidAdmin bool
	IsValidUser  bool
	IsVisitor    bool
}

type AdminContex struct {
	AdminID            uint64
	AdminType          string
	ClientID           uint64
	ShopID             uint64
	IsValidAdmin       bool
	IsValidSystemAdmin bool
	IsValidClientAdmin bool
	IsValidShopAdmin   bool
}

func (lg *LoginContext) GetAdminContext() *AdminContex {
	ac := AdminContex{
		AdminID:   lg.AdminID,
		AdminType: lg.AdminType,
		// ClientID:           lg.ClientID,
		// ShopID:             lg.ShopID,
		IsValidAdmin:       lg.IsValidAdmin,
		IsValidSystemAdmin: lg.IsValidAdmin && lg.AdminType == auth.AdminTypeSystem,
		IsValidClientAdmin: lg.IsValidAdmin && lg.AdminType == auth.AdminTypeClient,
		IsValidShopAdmin:   lg.IsValidAdmin && lg.AdminType == auth.AdminTypeShop,
	}

	return &ac
}

func IsUserAuthorized(c *gin.Context) (bool, uint64) {
	loginContext := getLoginContext(c)
	return loginContext.IsValidUser, loginContext.UserID
}

func IsVisitor(c *gin.Context) (bool, string) {
	loginContext := getLoginContext(c)
	return loginContext.IsVisitor, loginContext.Visitor
}

func IsAdminAuthorized(c *gin.Context) (bool, uint64, *AdminContex) {
	loginContext := getLoginContext(c)
	return loginContext.IsValidAdmin, loginContext.AdminID, loginContext.GetAdminContext()
}

func IsSystemAdminAuthorized(c *gin.Context) (bool, uint64) {
	loginContext := getLoginContext(c)
	return loginContext.IsValidAdmin && loginContext.AdminType == auth.AdminTypeSystem, loginContext.AdminID
}

func IsClientAdminAuthorized(c *gin.Context) (bool, uint64) {
	loginContext := getLoginContext(c)
	return loginContext.IsValidAdmin && loginContext.AdminType == auth.AdminTypeClient, loginContext.AdminID
}

func getLoginContext(c *gin.Context) LoginContext {

	if value, ok := c.Keys[loginContextKey]; ok {
		return value.(LoginContext)
	}

	// loginContext := LoginContext{
	// 	UserID:      0,
	// 	AdminID:     0,
	// 	Visitor:     "",
	// 	IsValid:     false,
	// 	IsValidUser: false,
	// 	IsVisitor:   false,
	// }

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
		IsValid:   ok,
		UserID:    payload.UserID,
		AdminID:   payload.AdminID,
		AdminType: payload.AdminType,
		// ClientID:  payload.ClientID,
		// ShopID:       payload.ShopID,
		IsValidAdmin: payload.AdminID > 0 && ok,
		IsValidUser:  payload.UserID > 0 && ok,
		Visitor:      payload.Visitor,
		IsVisitor:    payload.Visitor != "",
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
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": message})
}

func renderUnauthorizedError(c *gin.Context) {
	c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "user is not authorized"})
}

func renderSuccessMessage(c *gin.Context, message string, a ...interface{}) {
	if a != nil {
		message = fmt.Sprintf(message, a...)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": message})
}

func renderError(c *gin.Context, err error) {
	renderErrorMessage(c, err.Error())
}

func renderJSON(c *gin.Context, obj interface{}) {
	c.IndentedJSON(http.StatusOK, obj)
}
