package wagen

// This file handles all the routes of the app

import (
	"github.com/gorilla/mux"
	"github.com/jeanbouz/wagen/controllers"
)

// InitializeRouter initializes the router
// We use the package gorilla/mux in order manage the router
func InitializeRouter() *mux.Router {
	// When StrictSlash is true, routes like /cars/ are redirected to /cars
	router := mux.NewRouter().StrictSlash(true)

	// In order to define the routes we use different fucntions :
	// - Methods() defines the method which is manage by our route
	// - Path() is the path that we want to link we the action
	// - Name() is the name that we want to give to our route
	// - HandlerFucn() gives the functions that is linked with the route
	router.Methods("GET").Path("/cars").Name("Index").HandlerFunc(controllers.CarsIndex)
	router.Methods("POST").Path("/cars").Name("Create").HandlerFunc(controllers.CarsCreate)
	router.Methods("GET").Path("/cars/{id}").Name("Show").HandlerFunc(controllers.CarsShow)
	router.Methods("PUT").Path("/cars/{id}").Name("Update").HandlerFunc(controllers.CarsUpdate)
	router.Methods("DELETE").Path("/cars/{id}").Name("Delete").HandlerFunc(controllers.CarsDelete)

	router.Methods("GET").Path("/customers").Name("Index").HandlerFunc(controllers.CustomersIndex)
	router.Methods("POST").Path("/customers").Name("Create").HandlerFunc(controllers.CustomersCreate)
	router.Methods("GET").Path("/customers/{id}").Name("Show").HandlerFunc(controllers.CustomersShow)
	router.Methods("PUT").Path("/customers/{id}").Name("Update").HandlerFunc(controllers.CustomersUpdate)
	router.Methods("DELETE").Path("/customers/{id}").Name("Delete").HandlerFunc(controllers.CustomersDelete)

	router.Methods("GET").Path("/users").Name("Index").HandlerFunc(controllers.UsersIndex)
	router.Methods("POST").Path("/users").Name("Create").HandlerFunc(controllers.UsersCreate)
	router.Methods("GET").Path("/users/{id}").Name("Show").HandlerFunc(controllers.UsersShow)
	router.Methods("PUT").Path("/users/{id}").Name("Update").HandlerFunc(controllers.UsersUpdate)
	router.Methods("DELETE").Path("/users/{id}").Name("Delete").HandlerFunc(controllers.UsersDelete)

	return router
}
