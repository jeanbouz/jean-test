package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jeanbouz/wagen/config"
	"github.com/jeanbouz/wagen/models"
	router "github.com/jeanbouz/wagen/router"
)

func main() {
	config.DatabaseInit()
	rout := router.InitializeRouter()

	/* Populates database
	models.NewCar(&models.Car{Owner: 1, Brand: "Peugeot", Model: "201"})
	models.NewCustomer(&models.Customer{Name: "Fau", Surname: "Lele", Email: "lele.fau@gmail.com"})
	models.NewUser(&models.User{Pseudo: "admin", Password: "admin", Email: "admin.admin@gmail.com", UserType: "admin"})
	*/

	// Print data from the database
	fmt.Println(models.FindCarByID(11).String())
	fmt.Println(models.FindCustomerByID(4).String())
	fmt.Println(models.FindUserByID(2).String())

	/* Print all the data base
	fmt.Println(models.AllCars())
	fmt.Println(models.AllCustomers())
	fmt.Println(models.AllUsers())
	*/

	// Update data in the database
	models.UpdateCar(&models.Car{ID: 11, Owner: 3, Brand: "MG", Model: "B"})
	models.UpdateCustomer(&models.Customer{ID: 4, Name: "And", Surname: "Popo", Email: "popo.and@gmail.com"})
	models.UpdateUser(&models.User{ID: 2, Pseudo: "lambda", Password: "password", Email: "lambda.password@gmail.com", UserType: "normal"})

	// Print data from the database
	fmt.Println(models.FindCarByID(11).String())
	fmt.Println(models.FindCustomerByID(4).String())
	fmt.Println(models.FindUserByID(2).String())

	// Delete data in the database
	err := models.DeleteCarByID(10)
	if err != nil {
		panic(err.Error())
	}
	err = models.DeleteCustomerByID(3)
	if err != nil {
		panic(err.Error())
	}
	err = models.DeleteUserByID(1)
	if err != nil {
		panic(err.Error())
	}

	// Fatal is equivalent to Print() followed by a call to os.Exit(1)
	log.Fatal(http.ListenAndServe(":8080", rout))

}
