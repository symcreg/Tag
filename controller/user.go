package controller

import (
	"Tag/utility"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	isVerify  bool
	VerifyUrl string
}

func Register(c *gin.Context) {
	var user User
	c.ShouldBindJSON(&user)                         //密码应为rsa加密后的密文
	user.Password, _ = utility.UnRSA(user.Password) //私钥解密
	if len(user.Username) < 4 || len(user.Username) > 16 || len(user.Password) < 6 || len(user.Password) > 16 {
		c.JSON(403, gin.H{
			"code": 403,
			"msg":  "wrong format",
		})
		return
	} //判断是否符合格式，应为4至16位
	user.Password, _ = utility.HashPassword(user.Password) //hash加密密码,存入数据库的为hash值
	url := Md5(user.Email)
	user.VerifyUrl = url      //验证url
	utility.EmailVerify(user) //发送验证邮件
	db, err := gorm.Open("sqlite3", "tag.wall")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Create(&user) //存入数据库
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"user": user.Username,
	}) //返回数据

}
func Login(c *gin.Context) {
	var user User
	var userFromDB User
	c.ShouldBindJSON(&user)                                //密码应为rsa加密后的密文
	user.Password, _ = utility.UnRSA(user.Password)        //私钥解密
	user.Password, _ = utility.HashPassword(user.Password) //hash
	db, err := gorm.Open("sqlite3", "tag.db")
	if err != nil {
		panic(err)
	}
	db.Where("username=?", user.Username).First(&userFromDB)
	if user.Username != userFromDB.Username || !utility.VerifyPassword(user.Password, userFromDB.Password) {
		c.JSON(403, gin.H{
			"code": 403,
			"msg":  "用户名或密码错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"user": user.Username,
	}) //登陆成功,返回数据,前端应请求token
	return
}
