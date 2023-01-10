package config 

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
 

func Connect() {
	dsn := "root:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "newUser:password@tcp(mysql_db:3306)/dbname"
	// dsn := "user:password@tcp(db)/dbname?charset=utf8"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connection Successful")
	db = d
}

func GetDB() *gorm.DB {
	return db
}