package db

import (
	"Tag/controller"
	"github.com/jinzhu/gorm"
)

func InitDB() {
	db, err := gorm.Open("sqlite3", "tag.db")
	if err != nil {
		panic("init database error")
	}
	defer db.Close()
	db.AutoMigrate(&controller.User{}) //结构有变动自动迁移
	db.AutoMigrate(&controller.Tag{})
}
