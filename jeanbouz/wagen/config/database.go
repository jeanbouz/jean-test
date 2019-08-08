// Package config : is used in order to write the code related to the database management
package config

import (
	"database/sql"

	// import driver
	// _ "..." allows to only import (without initialisation), in order to avoid side effects
	_ "github.com/go-sql-driver/mysql"
)

var (
	// db will contain the instance of SQL connexion to the database
	db *sql.DB
)

const (
	// String used in the createCarsTable function, for readability reasons
	carsTableCreationString string = "CREATE TABLE IF NOT EXISTS `cars`(" +
		"id serial," +
		"owner int," +
		"brand varchar(20)," +
		"model varchar(20)," +
		"created_at TIMESTAMP default CURRENT_TIMESTAMP," +
		"updated_at TIMESTAMP default CURRENT_TIMESTAMP," +
		"constraint pk primary key(id))"

	// String used in the createCustomersTable function, for readability reasons
	customersTableCreationString string = "CREATE TABLE IF NOT EXISTS `customers`(" +
		"id serial," +
		"name varchar(20)," +
		"surname varchar(20)," +
		"email varchar(30)," +
		"created_at TIMESTAMP default CURRENT_TIMESTAMP," +
		"updated_at TIMESTAMP default CURRENT_TIMESTAMP," +
		"constraint pk primary key(id))"

	// String used in the createUsersTable function, for readability reasons
	usersTableCreationString string = "CREATE TABLE IF NOT EXISTS `users`(" +
		"id serial," +
		"pseudo varchar(20)," +
		"password varchar(20)," +
		"email varchar(30)," +
		//"user_type enum('admin', 'normal')," + or use another table in the database with all the types of users
		"user_type varchar(20)," +
		"created_at TIMESTAMP default CURRENT_TIMESTAMP," +
		"updated_at TIMESTAMP default CURRENT_TIMESTAMP," +
		"constraint pk primary key(id))"
)

// Db is a public getter for db var
func Db() *sql.DB {
	return db
}

// DatabaseInit initialize the database
func DatabaseInit() {
	var err error

	// First argument "mysql" is the name of our driver
	// Second argument concatenation of : user name: root, password: ,
	//protocol:tcp, IP adress: 127.0.0.1:3306, name of the data base: garage
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/garage?parseTime=true")
	// If we cannot connect to the database
	if err != nil {
		panic(err.Error())
	}

	// Creates Table cars if not exists
	createCarsTable()

	// Creates Table customers (owners of the cars) if not exists
	createCustomersTable()

	// Creates Table users (people using the API) if not exists
	createUsersTable()
}

// createCarsTable creates Table cars if not exists
func createCarsTable() {
	_, err := db.Exec(carsTableCreationString)

	if err != nil {
		panic(err.Error())
	}
}

// createCustomersTable creates Table customers if not exists
func createCustomersTable() {
	_, err := db.Exec(customersTableCreationString)

	if err != nil {
		panic(err.Error())
	}
}

// createUsersTable creates Table users if not exists
func createUsersTable() {
	_, err := db.Exec(usersTableCreationString)

	if err != nil {
		panic(err.Error())
	}
}
