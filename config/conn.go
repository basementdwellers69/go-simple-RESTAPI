package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbErr error

func GetConnection() *gorm.DB {
	dsn := "host="+host+" user="+user+" password="+pass+" dbname="+dbName+" port="+port+" sslmode=disable TimeZone=Asia/Singapore"
	db, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil{
		panic(dbErr.Error())
	}else{
		fmt.Println("Connection Established")
	}
	return db
}
