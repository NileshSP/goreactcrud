package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error
var dbPath = "../database/Sample.db"
var dbType = "sqlite3"

func InitialMigrations() {
	fmt.Println("Migrations started")
	if conn, ok := getDbConn(); ok {
		conn.AutoMigrate(&User{})
	}
	fmt.Println("Migrations completed")
}

func getDbConn() (*gorm.DB, bool) {
	db, err = gorm.Open(dbType, dbPath)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	return db, true
}
