package models

import (
  "gorm.io/gorm"
  _ "gorm.io/driver/mysql"
  "github.com/saurabh/carApi/package/config"
)

var db *gorm.DB

type Car struct {
  gorm.Model
  Name  string `gorm:""json:"name"`
  Company string `json:"company"`
  
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Car{})
}

func (b *Car) CreateCar() *Car{
  db.Create(&b)
  return b
}

func GetAllCars() []Car{
  var cars []Car
  db.Find(&cars)
  return cars
}

func GetCarById(carId int64) (*Car, *gorm.DB) {
  var car Car
  db := db.Where("ID=?", carId).Find(&car)
  return &car, db
}


func DeleteCar(carId int64) Car {
  var car Car
  db.Where("ID=?", carId).Delete(&car)
  return car
}



