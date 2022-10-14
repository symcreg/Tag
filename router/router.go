package router

import (
	"Tag/controller"
	"Tag/utility"
	"Tag/utility/middlerware"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine = gin.Default()

func SetupRouter() {
	Router.Use(middlerware.Authorization)
	Router.POST("/api/user/reg", controller.Register) //注册
	Router.POST("/api/user/login", controller.Login)  //登录
	Router.GET("/api/user/email", utility.EmailVerifyHandler)
	Router.POST("/api/user/rsa", utility.RSAHandler)               //rsa公钥
	Router.GET("/api/user/auth", utility.GenerateTokenHandler)     //token
	Router.POST("/api/img/uploadImg", controller.UploadImgHandler) //上传img
	//router.GET("/api/img/getImg", controller.GetImgHandler)           //获取img
	Router.POST("/api/tag/addTag", controller.AddTagHandler)          //添加一条tag
	Router.GET("/api/tag/getOne", controller.GetOne)                  //获取一条tag
	Router.GET("/api/tag/GetAllMyTag", controller.GetAllMyTagHandler) //获取用户自身上传的tags
	Router.DELETE("/api/tag/deleteTag", controller.DeleteTagHandler)  //删除一条tag
}
