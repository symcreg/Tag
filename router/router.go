package router

import (
	"Tag/controller"
	"Tag/utility"
	"Tag/utility/middlerware"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine = gin.Default()

func SetupRouter() {
	router.Use(middlerware.Authorization)
	router.POST("/api/user/reg", controller.Register)              //注册
	router.POST("/api/uer/login", controller.Login)                //登录
	router.POST("/api/user/rsa", utility.RSAHandler)               //rsa公钥
	router.GET("/api/user/auth", utility.GenerateTokenHandler)     //token
	router.POST("/api/img/uploadImg", controller.UploadImgHandler) //上传img
	//router.GET("/api/img/getImg", controller.GetImgHandler)           //获取img
	router.POST("/api/tag/addTag", controller.AddTagHandler)          //添加一条tag
	router.GET("/api/tag/getOne", controller.GetOne)                  //获取一条tag
	router.GET("/api/tag/GetAllMyTag", controller.GetAllMyTagHandler) //获取用户自身上传的tags
	router.DELETE("/api/tag/deleteTag", controller.DeleteTagHandler)  //删除一条tag
}
