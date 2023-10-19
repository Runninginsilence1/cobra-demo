package database

import (
	"Forensics_Equipment_Plugin_Manager/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {
	// 1. 获取连接
	var err error
	Db, err = gorm.Open(sqlite.Open("resources/test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database, no sqlite database file.")
	}

	fmt.Println("SQLite initialize success")

	//	2. 迁移数据库
	err = Db.AutoMigrate(&models.Plugin{})
	if err != nil {
		fmt.Println("[Error]: AutoMigrate")
		return
	}
}
