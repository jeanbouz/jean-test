package models

// This file makes the link between the controller and the database and
// creates the struct Customer and all the functions that allow to manipulate it

import (
	"fmt"
	"log"
	"time"

	// import driver
	// _ "..." allows to only import (without initialisation), in order to avoid side effects
	_ "github.com/go-sql-driver/mysql"
	"github.com/jeanbouz/wagen/config"
)

// Customer contains all the information of a customer
type Customer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// String is a method used to display a Customer
func (c Customer) String() string {
	return fmt.Sprintf("Id : %v, Name : %v %v, Email : %v -- Created at %v, Updated at %v",
		c.ID, c.Surname, c.Name, c.Email, c.CreatedAt, c.UpdatedAt)
}

// Customers is a slice of Customer which contains all the customers
type Customers []Customer

// NewCustomer allow us to stock a new Customer in the database
func NewCustomer(c *Customer) {
	if c == nil {
		log.Fatal(c)
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	insert, err := config.Db().Query("INSERT INTO customers (name, surname, email, created_at, updated_at) VALUES (?,?,?,?,?);",
		c.Name, c.Surname, c.Email, c.CreatedAt, c.UpdatedAt)

	if err != nil {
		panic(err.Error())
	}
	// We defer the closure of insert
	defer insert.Close()

	if err != nil {
		panic(err.Error())
	}
}

// FindCustomerByID allows us to pick customers up from the database with their ids
func FindCustomerByID(id int) *Customer {
	var customer Customer

	row := config.Db().QueryRow("SELECT * FROM customers WHERE id = ?;", id)

	err := row.Scan(&customer.ID, &customer.Name, &customer.Surname, &customer.Email, &customer.CreatedAt, &customer.UpdatedAt)

	if err != nil {
		panic(err.Error())
	}

	return &customer
}

// AllCustomers allows us to pick up all the customers from the database
func AllCustomers() *Customers {
	var customers Customers

	rows, err := config.Db().Query("SELECT * FROM customers")

	if err != nil {
		panic(err.Error())
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var c Customer

		err := rows.Scan(&c.ID, &c.Name, &c.Surname, &c.Email, &c.CreatedAt, &c.UpdatedAt)

		if err != nil {
			panic(err.Error())
		}

		customers = append(customers, c)
	}

	return &customers
}

// UpdateCustomer allows us to update a value stored in the database
func UpdateCustomer(customer *Customer) {
	customer.UpdatedAt = time.Now()

	// The function Prepare, prepares the SQL query
	stmt, err := config.Db().Prepare("UPDATE customers SET name=?, surname=?, email=?, updated_at=? WHERE id=?;")

	if err != nil {
		log.Fatal(err)
	}

	// The function Exec, initiates the SQL query
	_, err = stmt.Exec(customer.Name, customer.Surname, customer.Email, customer.UpdatedAt, customer.ID)

	if err != nil {
		panic(err.Error())
	}
}

// DeleteCustomerByID allows us to delete a value stored in the database
func DeleteCustomerByID(id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM customers WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)

	return err
}
