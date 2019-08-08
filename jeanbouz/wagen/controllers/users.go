package controllers

// This file handles the relations between the http queries and the User struct

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeanbouz/wagen/models"
)

// UsersIndex allows us to send all the users in the database
func UsersIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.AllUsers())
}

// UsersCreate allows us to create a new user
func UsersCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// We pick up the Body
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err.Error())
	}

	var user models.User

	// Unmarshal parses the content of body and puts it in user
	err = json.Unmarshal(body, &user)

	if err != nil {
		panic(err.Error())
	}

	// Persists the variable user
	models.NewUser(&user)

	json.NewEncoder(w).Encode(user)
}

// UsersShow allows us to show a user
func UsersShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Picks up the id of the userr in the url of the query
	vars := mux.Vars(r)
	// Converts the id from string to int
	// strconv.Atoi is shorthand for ParsInt
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err.Error())
	}

	user := models.FindUserByID(id)

	json.NewEncoder(w).Encode(user)
}

// UsersUpdate allows us to update a user
func UsersUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Picks up the id of the user in the url of the query
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

	user := models.FindUserByID(id)

	// Unmarshal parses the content of body and puts it in user
	err = json.Unmarshal(body, &user)

	if err != nil {
		panic(err.Error())
	}

	models.UpdateUser(user)

	json.NewEncoder(w).Encode(user)
}

// UsersDelete allows us to delete a user
func UsersDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Picks up the id of the user in the url of the query
	vars := mux.Vars(r)
	// Converts the id from string to int
	// strconv.Atoi is shorthand for ParsInt
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err.Error())
	}

	err = models.DeleteUserByID(id)

	if err != nil {
		panic(err.Error())
	}
}
