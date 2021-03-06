package conn

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbErr error

func GetConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Singapore"
	db, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil{
		panic(dbErr.Error())
	}else{
		fmt.Println("Connection Established")
	}
	return db
}