package main

import (
	"log"
	"net/http"

	"github.com/jeanbouz/wagen/config"
	"github.com/jeanbouz/wagen/models"
	"github.com/jeanbouz/wagen/router/"
)

func main() {
	config.DatabaseInit()
	rout := router.InitializeRouter()

	// Populates database
	models.NewCar(&models.Car{Manufacturer: "Citroen", Design: "DS3", Style: "sport", Doors: 4})

	log.Fatal(http.ListenAndServe(":8080", rout))
}
