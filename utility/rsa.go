package utility

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"github.com/gin-gonic/gin"
)

var PrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
var PublicKey = PrivateKey.Public()

func RSAHandler(c *gin.Context) {
	if err != nil {
		panic(err)
		return
	}
	c.JSON(200, gin.H{
		"code":      200,
		"msg":       "success",
		"PublicKey": PublicKey.(string),
	})
} //公钥

func UnRSA(EncryptedData string) (string, error) {
	DecryptedData, err := PrivateKey.Decrypt(nil, []byte(EncryptedData), &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
		return "", err
	}
	return string(DecryptedData), nil
} //私钥解密
