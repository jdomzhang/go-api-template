package auth

import (
	"fmt"
	"strconv"
	"time"

	"{{name}}/src/config"
	"{{name}}/src/util"

	"{{name}}/src/models/orm"

	"github.com/dgrijalva/jwt-go"
)

// Payload is the structure to store user related information
type Payload struct {
	UserID              uint64 `json:"uid,omitempty"`
	WechatOpenID        string `json:"wid,omitempty"`
	EncryptedSessionKey string `json:"esk,omitempty"`
	Visitor             string `json:"visitor,omitempty"`
	jwt.StandardClaims
}

// CreateJwt will create token
func CreateJwt(payload *Payload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	ss, err := token.SignedString([]byte(config.All["JwtSignKey"]))

	return ss, err
}

// ParseJwt will parse jwt token to payload
func ParseJwt(tokenString string) (bool, *Payload) {
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.All["JwtSignKey"]), nil
	})

	if err != nil {
		fmt.Println(err)
		return false, &Payload{}
	}

	payload, ok := token.Claims.(*Payload)
	if ok && token.Valid {
		return ok, payload
	}

	fmt.Printf("Result: %v, Valid: %v", ok, token.Valid)
	return false, payload
}

// Gen will generate jwt token
func (base *Payload) Gen() string {
	expireMinutes, _ := strconv.Atoi(config.All["token.expire.minutes"])

	base.ExpiresAt = time.Now().Add(time.Minute * time.Duration(expireMinutes)).Unix()

	tokenString, _ := CreateJwt(base)

	return tokenString
}

// Check will validate jwt token
func Check(token string) (bool, *Payload) {
	return ParseJwt(token)
}

// GenVisitorJwt will generate a new visitor jwt
func GenVisitorJwt() string {
	// generate a vistor token
	randomUser := util.MD5(fmt.Sprintf("%v", time.Now().UnixNano()))
	payload := Payload{
		Visitor: randomUser,
	}

	return payload.Gen()
}

// GenUserJwt will generate user jwt token by a given user
func GenUserJwt(user *orm.User) string {
	payload := Payload{
		UserID:       user.ID,
		WechatOpenID: user.WechatOpenID,
		// EncryptedSessionKey: user.WechatSessionKey,
	}

	return payload.Gen()
}
