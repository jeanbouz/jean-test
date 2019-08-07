package models

// This file makes the link between the controller and the database and
// creates the struct User and all the functions that allow to manipulate it

import (
	"fmt"
	"log"
	"time"

	// import driver
	// _ "..." allows to only import (without initialisation), in order to avoid side effects
	_ "github.com/go-sql-driver/mysql"
	"github.com/jeanbouz/wagen/config"
)

// User contains all the information of a user
type User struct {
	ID        int       `json:"id"`
	Pseudo    string    `json:"pseudo"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// String is a method used to display a User
func (u User) String() string {
	return fmt.Sprintf("Id : %v, Pseudo : %v, Type : %v, Password : %v, Email : %v -- Created at %v, Updated at %v",
		u.ID, u.Pseudo, u.UserType, u.Password, u.Email, u.CreatedAt, u.UpdatedAt)
}

// Users is a slice of User which contains all the users
type Users []User

// NewUser allow us to stock a new USer in the database
func NewUser(u *User) {
	if u == nil {
		log.Fatal(u)
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	insert, err := config.Db().Query("INSERT INTO users (pseudo, password, email, user_type, created_at, updated_at) VALUES (?,?,?,?,?,?);",
		u.Pseudo, u.Password, u.Email, u.UserType, u.CreatedAt, u.UpdatedAt)

	if err != nil {
		panic(err.Error())
	}
	// We defer the closure of insert
	defer insert.Close()

	if err != nil {
		panic(err.Error())
	}
}

// FindUserByID allows us to pick users up from the database with their ids
func FindUserByID(id int) *User {
	var user User

	row := config.Db().QueryRow("SELECT * FROM users WHERE id = ?;", id)

	err := row.Scan(&user.ID, &user.Pseudo, &user.Password, &user.Email, &user.UserType, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		panic(err.Error())
	}

	return &user
}

// AllUsers allows us to pick up all the users from the database
func AllUsers() *Users {
	var users Users

	rows, err := config.Db().Query("SELECT * FROM users")

	if err != nil {
		panic(err.Error())
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var u User

		err := rows.Scan(&u.ID, &u.Pseudo, &u.Password, &u.Email, &u.UserType, &u.CreatedAt, &u.UpdatedAt)

		if err != nil {
			panic(err.Error())
		}

		users = append(users, u)
	}

	return &users
}

// UpdateUser allows us to update a value stored in the database
func UpdateUser(user *User) {
	user.UpdatedAt = time.Now()

	// The function Prepare, prepares the SQL query
	stmt, err := config.Db().Prepare("UPDATE users SET pseudo=?, password=?, email=?, user_type=?, updated_at=? WHERE id=?;")

	if err != nil {
		log.Fatal(err)
	}

	// The function Exec, initiates the SQL query
	_, err = stmt.Exec(user.Pseudo, user.Password, user.Email, user.UserType, user.UpdatedAt, user.ID)

	if err != nil {
		panic(err.Error())
	}
}

// DeleteUserByID allows us to delete a value stored in the database
func DeleteUserByID(id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM users WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)

	return err
}
