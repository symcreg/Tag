package utility

import (
	"Tag/controller"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	email2 "github.com/jordan-wright/email"
	"net/smtp"
)

func EmailVerify(user controller.User) {
	Email := email2.NewEmail()
	Email.From = "Tag"
	Email.To = []string{user.Email}
	Email.Subject = "Verify your email"
	Email.Text = []byte("Please visit the url to verify your email address" + user.VerifyUrl)
	err = Email.Send("smtp.163.com", smtp.PlainAuth("", "pointreg@163.com", "ZXRUMUULGFZZYWHO", "smtp.163.com"))
	if err != nil {
		panic("send email error")
	}
}
func EmailVerifyHandler(c *gin.Context) {
	url := c.Query("url")
	db, err := gorm.Open("sqlite3", "tag.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Where("url=?", url).Update("isVerify", true)
}
