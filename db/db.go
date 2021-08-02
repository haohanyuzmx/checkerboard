package db

import (
	"checkerboard/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	DB = db
	err = DB.AutoMigrate(&model.ChessIn{})
	if err != nil {
		log.Println(err)
		return
	}
}
func Creat(in model.ChessIn) {
	DB.Create(&in)
}
