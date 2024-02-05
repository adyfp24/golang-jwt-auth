package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error){
	dsn := "root:@tcp(127.0.0.1:3306)/db_go_jwt?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!= nil{
		return nil, err
	}
	fmt.Println("koneksi ke database sukses")
	DB = db
	return db, nil
}

// kill -9 $(lsof -i :4000 -t)