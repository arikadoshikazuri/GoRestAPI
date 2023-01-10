package main 

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
	"github.com/saurabh/carApi/package/routes"
)

func main() {
	r := mux.NewRouter()
	routes.CarRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}