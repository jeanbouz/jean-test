package controllers

// This file handles the relations between the http queries and the Car struct

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeanbouz/wagen/models"
)

// CarsIndex allows us to send all the cars in the database
func CarsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.AllCars())
}

// CarsCreate allows us to create a new car
func CarsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// We pick up the Body
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err.Error())
	}

	var car models.Car

	// Unmarshal parses the content of body and puts it in car
	err = json.Unmarshal(body, &car)

	if err != nil {
		panic(err.Error())
	}

	// Persists the variable car
	models.NewCar(&car)

	json.NewEncoder(w).Encode(car)
}

// CarsShow allows us to show a car
func CarsShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Vars allow us to pick
	vars := mux.Vars(r)
	fmt.Printf("In cars.go, function CarsShow -> Vars(r) = %v\n", vars)
	// Converts the id from string to int
	// strconv.Atoi is shorthand for ParsInt
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err.Error())
	}

	car := models.FindCarByID(id)

	json.NewEncoder(w).Encode(car)
}

// CarsUpdate allows us to update a car
func CarsUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Picks up the id of the car in the url of the query
	vars := mux.Vars(r)
	// Converts the id from string to int
	// strconv.Atoi is shorthand for ParsInt
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err.Error())
	}

	car := models.FindCarByID(id)

	// Unmarshal parses the content of body and puts it in car
	err = json.Unmarshal(body, &car)

	if err != nil {
		panic(err.Error())
	}

	models.UpdateCar(car)

	json.NewEncoder(w).Encode(car)
}

// CarsDelete allows us to delete a car
func CarsDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Picks up the id of the car in the url of the query
	vars := mux.Vars(r)
	// Converts the id from string to int
	// strconv.Atoi is shorthand for ParsInt
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err.Error())
	}

	err = models.DeleteCarByID(id)

	if err != nil {
		panic(err.Error())
	}
}
