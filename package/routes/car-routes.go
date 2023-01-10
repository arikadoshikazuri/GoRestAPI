package routes

import (
	"github.com/gorilla/mux"
	"github.com/saurabh/carApi/package/controllers"
)

var CarRoutes = func(router *mux.Router){
	router.HandleFunc("/car", controllers.CreateCar).Methods("POST")
	router.HandleFunc("/car", controllers.GetCars).Methods("GET")
	router.HandleFunc("/car/{carId}", controllers.GetCarById).Methods("GET")
	router.HandleFunc("/car/{carId}", controllers.UpdateCar).Methods("PUT")
	router.HandleFunc("/car/{carId}", controllers.DeleteCar).Methods("DELETE")
	
}