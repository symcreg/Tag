package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"os"
	"path"
	"strings"
	"time"
)

type Tag struct {
	Id          int      `json:"id"`
	UserId      int      `json:"userId"`
	Img         string   `json:"img"`
	tags        []string `json:"tags"`
	Description string   `json:"description"`
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func UploadImgHandler(c *gin.Context) {
	img, err := c.FormFile("imgFile")
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "upload error",
		})
		return
	}
	imgExt := strings.ToLower(path.Ext(img.Filename))
	if imgExt != "png" && imgExt != "jpg" && imgExt != "jpeg" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "only allow .jpg|.jpeg|.png",
		})
		return
	}
	imgName := Md5(fmt.Sprintf("%s%s", img.Filename, time.Now().String()))
	imgDir := "img/"
	_, osErr := os.Stat(imgDir)
	if osErr != nil {
		os.Mkdir(imgDir, os.ModePerm)
	}
	imgPath := fmt.Sprintf("%s%s%s", imgDir, imgName, imgExt)
	c.SaveUploadedFile(img, imgPath)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"img":  imgName + imgExt,
	})
}
func AddTagHandler(c *gin.Context) { //另需加以验证身份
	var tag Tag
	c.ShouldBindJSON(&tag)
	db, err := gorm.Open("sqlite3", "tag.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Create(&tag)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
func DeleteTagHandler(c *gin.Context) { //另需加以验证身份
	id := c.Query("id")
	db, err := gorm.Open("sqlite3", "tag.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Where("id=?", id).Delete(&Tag{})
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
func GetOne(c *gin.Context) {
	var tag Tag
	db, err := gorm.Open("sqlite3", "tag.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Take(&tag) //从数据库中拿出一条数据
	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "success",
		"tag":  tag,
	})
	imgPath := "img/"
	c.File(imgPath + tag.Img)
}
func GetAllMyTagHandler(c *gin.Context) { //另需加以验证身份
	userid := c.Query("userid")
	var tags []Tag
	db, err := gorm.Open("sqlite3", "tag.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Where("userid=?", userid).Find(&tags)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"tags": tags,
	})
}

//func GetImgHandler(c *gin.Context) {
//	imgPath := "img/"
//	img := c.Query("img")
//	c.File(imgPath + img)
//}
