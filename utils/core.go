package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	dsn := "root:abcd1234@tcp(127.0.0.1:3306)/go_todolist?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
	} else {
		log.Printf("Error occoured when connecting to the data base! error: %e", err)
	}
}
