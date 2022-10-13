package utility

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var secret = "SYMC"                 //盐
var ExpireTime = time.Hour * 24 * 2 //过期时间2d
type Jwt struct {
	Id   int    `json:"id"`
	User string `json:"user"`
	jwt.StandardClaims
}

func GenerateTokenHandler(c gin.Context) {
	var claim Jwt
	c.ShouldBindJSON(&claim)
	claim.ExpiresAt = time.Now().Add(ExpireTime).Unix() //过期时间
	claim.Issuer = "SYMC"                               //签发人
	token, err := GenToken(claim)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "error",
		})
	}
	c.JSON(200, gin.H{
		"code":  200,
		"msg":   "success",
		"token": token,
	})
}
func GenToken(claim Jwt) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString("SYMC") //加盐
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
func ParseToken(tokenStr string) (*Jwt, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Jwt{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	claim, ok := token.Claims.(*Jwt)
	if ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("invalid token")
}
