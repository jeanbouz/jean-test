package controllers

// This file handles the relations between the http queries and the Customer struct

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeanbouz/wagen/models"
)

// CustomersIndex allows us to send all the customers in the database
func CustomersIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.AllCustomers())
}

// CustomersCreate allows us to create a new customer
func CustomersCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// We pick up the Body
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err.Error())
	}

	var customer models.Customer

	// Unmarshal parses the content of body and puts it in customer
	err = json.Unmarshal(body, &customer)

	if err != nil {
		panic(err.Error())
	}

	// Persists the variable customer
	models.NewCustomer(&customer)

	json.NewEncoder(w).Encode(customer)
}

// CustomersShow allows us to show a customer
func CustomersShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Picks up the id of the customerr in the url of the query
	vars := mux.Vars(r)
	// Converts the id from string to int
	// strconv.Atoi is shorthand for ParsInt
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err.Error())
	}

	customer := models.FindCustomerByID(id)

	json.NewEncoder(w).Encode(customer)
}

// CustomersUpdate allows us to update a customer
func CustomersUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Picks up the id of the customer in the url of the query
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

	customer := models.FindCustomerByID(id)

	// Unmarshal parses the content of body and puts it in customer
	err = json.Unmarshal(body, &customer)

	if err != nil {
		panic(err.Error())
	}

	models.UpdateCustomer(customer)

	json.NewEncoder(w).Encode(customer)
}

// CustomersDelete allows us to delete a customer
func CustomersDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Picks up the id of the customer in the url of the query
	vars := mux.Vars(r)
	// Converts the id from string to int
	// strconv.Atoi is shorthand for ParsInt
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err.Error())
	}

	err = models.DeleteCustomerByID(id)

	if err != nil {
		panic(err.Error())
	}
}
