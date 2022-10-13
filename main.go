package main

import (
	"Tag/db"
	"Tag/router"
)

func main() {
	db.InitDB()          //初始化数据库
	router.SetupRouter() //初始化路由
}
