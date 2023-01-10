package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/saurabh/carApi/package/utils"
	"github.com/saurabh/carApi/package/models"
)  

var newCars models.Car

func GetCars(w http.ResponseWriter, r *http.Request) {
	cars := models.GetAllCars()
	res, _ := json.Marshal(cars)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId := vars["carId"]
	id, err := strconv.ParseInt(carId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing carId")
	}
	carDetails, _ := models.GetCarById(id)
	res, _ := json.Marshal(carDetails )
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCar (w http.ResponseWriter, r *http.Request) {
	newCar := &models.Car{}
	utils.ParseBody(r, newCar)
	b := newCar.CreateCar()
	res, _ := json.Marshal(b) 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	carId := vars["carId"]
	id, err := strconv.ParseInt(carId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing carId")
	}

	deletedCar := models.DeleteCar(id)
	res, _ := json.Marshal(deletedCar)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCar(w http.ResponseWriter, r *http.Request){
	var updateCar = &models.Car{}
	utils.ParseBody(r, updateCar)
	
	vars := mux.Vars(r)
	carId := vars["carId"]
	id, err := strconv.ParseInt(carId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing carId")
	}
	car, db := models.GetCarById(id)
	if car.Name != "" {
		car.Name = updateCar.Name
	}
	if car.Company != "" {
		car.Company = updateCar.Company
	}
	
	db.Save(&car)
	res, _ := json.Marshal(car)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}