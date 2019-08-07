package models

// This file makes the link between the controller and the database and
// creates the struct Car and all the functions that allow to manipulate it

import (
	"fmt"
	"log"
	"time"

	// import driver
	// _ "..." allows to only import (without initialisation), in order to avoid side effects
	_ "github.com/go-sql-driver/mysql"
	"github.com/jeanbouz/wagen/config"
)

// Car contains all the information of a car
type Car struct {
	ID        int       `json:"id"`
	Owner     int       `json:"owner"`
	Brand     string    `json:"brand"`
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// String is a method used to display a Car
func (c Car) String() string {
	return fmt.Sprintf("Id : %v, Owner : %v, -- %v %v -- Created at %v, Updated at %v",
		c.ID, c.Owner, c.Brand, c.Model, c.CreatedAt, c.UpdatedAt)
}

// Cars is a slice of cars which contains all the cars
type Cars []Car

// NewCar allow us to stock a new Car in the database
func NewCar(c *Car) {
	if c == nil {
		log.Fatal(c)
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	insert, err := config.Db().Query("INSERT INTO cars (owner, brand, model, created_at, updated_at) VALUES (?,?,?,?,?);",
		c.Owner, c.Brand, c.Model, c.CreatedAt, c.UpdatedAt)

	if err != nil {
		panic(err.Error())
	}
	// We defer the closure of insert
	defer insert.Close()

	if err != nil {
		panic(err.Error())
	}
}

// FindCarByID allows us to pick cars up from the database with their ids
func FindCarByID(id int) *Car {
	var car Car

	row := config.Db().QueryRow("SELECT * FROM cars WHERE id = ?;", id)

	err := row.Scan(&car.ID, &car.Owner, &car.Brand, &car.Model, &car.CreatedAt, &car.UpdatedAt)

	if err != nil {
		panic(err.Error())
	}

	return &car
}

// AllCars allows us to pick up all the cars from the database
func AllCars() *Cars {
	var cars Cars

	rows, err := config.Db().Query("SELECT * FROM cars")

	if err != nil {
		panic(err.Error())
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var c Car

		err := rows.Scan(&c.ID, &c.Owner, &c.Brand, &c.Model, &c.CreatedAt, &c.UpdatedAt)

		if err != nil {
			panic(err.Error())
		}

		cars = append(cars, c)
	}

	return &cars
}

// UpdateCar allows us to update a value stored in the database
func UpdateCar(car *Car) {
	car.UpdatedAt = time.Now()

	// The function Prepare, prepares the SQL query
	stmt, err := config.Db().Prepare("UPDATE cars SET owner=?, brand=?, model=?, updated_at=? WHERE id=?;")

	if err != nil {
		log.Fatal(err)
	}

	// The function Exec, initiates the SQL query
	_, err = stmt.Exec(car.Owner, car.Brand, car.Model, car.UpdatedAt, car.ID)

	if err != nil {
		panic(err.Error())
	}
}

// DeleteCarByID allows us to delete a value stored in the database
func DeleteCarByID(id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM cars WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)

	return err
}
