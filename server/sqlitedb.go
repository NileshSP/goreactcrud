package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

var sqlitedb *gorm.DB
var err error
var dbPath = "../database/Sample.db"
var dbType = "sqlite3"

func InitializeSqlite() {
	fmt.Println("Migrations started")
	if conn, ok := getDbConn(); ok {
		conn.AutoMigrate(&User{})
	}
	fmt.Println("Migrations completed")
}

func getDbConn() (*gorm.DB, bool) {
	sqlitedb, err = gorm.Open(dbType, dbPath)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	return sqlitedb, true
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All users endpoint initialized")

	var users []User

	if conn, ok := getDbConn(); ok {

		conn.Find(&users)
		conn.Close()
		fmt.Println("Query output: ", users)
		json.NewEncoder(w).Encode(users)
		fmt.Println("All users endpoint finished")

	} else {
		fmt.Fprintf(w, "All users endpoint couldn't get db connection")
	}
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New user endpoint initialized")

	var user User

	if conn, ok := getDbConn(); ok {

		reqVals := mux.Vars(r)
		name := reqVals["name"]
		email := reqVals["email"]

		conn.Where(User{Name: name, Email: email}).FirstOrCreate(&user)
		conn.Close()

		fmt.Println(user)
		fmt.Println("New user endpoint finished")

	} else {
		fmt.Fprintf(w, "New user endpoint couldn't get db connection")
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete user endpoint initialized")

	if conn, ok := getDbConn(); ok {

		reqVals := mux.Vars(r)
		name := reqVals["name"]
		email := reqVals["email"]

		conn.Delete(User{}, "name = ? and email = ?", name, email)
		conn.Close()

		fmt.Println("Delete user endpoint finished")

	} else {
		fmt.Fprintf(w, "Delete user endpoint couldn't get db connection")
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update user endpoint initialized")

	var user User

	if conn, ok := getDbConn(); ok {

		reqVals := mux.Vars(r)
		name := reqVals["name"]
		email := reqVals["email"]

		conn.Where(User{Name: name}).Find(&user)
		user.Email = email
		sqlitedb.Save(&user)
		conn.Close()

		fmt.Println(user)
		fmt.Println("Update user endpoint finished")

	} else {
		fmt.Fprintf(w, "Update user endpoint couldn't get db connection")
	}
}
