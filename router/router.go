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
	router.POST("/api/user/reg", controller.Register)
	router.POST("/api/uer/login", controller.Login)
	router.POST("/api/user/rsa", utility.RSAHandler)
	router.GET("/api/user/auth", middlerware.Authorization)

}
