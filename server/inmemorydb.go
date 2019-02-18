package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	memdb "github.com/hashicorp/go-memdb"
)

var gomemdb *memdb.MemDB

// Create a sample struct
type Person struct {
	Email string
	Name  string
}

// Create the DB schema
var schema = &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{
		"person": &memdb.TableSchema{
			Name: "person",
			Indexes: map[string]*memdb.IndexSchema{
				"id": &memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "Name"},
				},
			},
		},
	},
}

// Create a new data base
func InitializeInMemoryDB() {
	gomemdb, err = memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
}

func GetPersons(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Persons endpoint initialized")
	if gomemdb != nil {
		// Create a write transaction
		txn := gomemdb.Txn(false)
		defer txn.Abort()

		result, err := txn.Get("person", "id")
		if err != nil {
			panic(err)
		}

		var persons []Person
		for {
			obj := result.Next()
			if obj == nil {
				break
			} else {
				//fmt.Println("GetPersons => person => ", obj)
				if objPerson, ok := obj.(*Person); ok {
					person := Person(*objPerson)
					persons = append(persons, person)
					//fmt.Println("GetPersons => person => after assertion => ", person)
				} else {
					fmt.Println("GetPersons => person => type assertion for person failed")
				}
			}
		}

		fmt.Println("Query output: ", persons)

		json.NewEncoder(w).Encode(persons)

	} else {
		fmt.Println("Get Persons endpoint couldn't get db connection")
	}
	fmt.Println("Get Persons endpoint completed")
}

func AddPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Person endpoint Initialized")
	if gomemdb != nil {

		reqVals := mux.Vars(r)
		name := reqVals["name"]
		email := reqVals["email"]

		var person = &Person{email, name}
		AddPersontoDb(person)

	} else {
		fmt.Println("Add person endpoint couldn't get db connection")
	}
	fmt.Println("Add Person endpoint completed")
}

func AddPersontoDb(person *Person) {
	// Create a write transaction
	txn := gomemdb.Txn(true)

	// Insert a new person
	if err := txn.Insert("person", person); err != nil {
		panic(err)
	} else {
		// Commit the transaction
		txn.Commit()
	}
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Person endpoint Initialized")
	if gomemdb != nil {
		// Create a write transaction
		txn := gomemdb.Txn(true)

		reqVals := mux.Vars(r)
		name := reqVals["name"]
		email := reqVals["email"]

		// Update a person
		var person = &Person{email, name}
		if err := txn.Insert("person", person); err != nil {
			panic(err)
		} else {
			// Commit the transaction
			txn.Commit()
		}
	} else {
		fmt.Println("Update Person endpoint couldn't get db connection")
	}
	fmt.Println("Update Person endpoint completed")
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Person endpoint Initialized")
	if gomemdb != nil {
		// Create a write transaction
		txn := gomemdb.Txn(true)

		reqVals := mux.Vars(r)
		name := reqVals["name"]
		email := reqVals["email"]

		// Delete a person
		var person = &Person{email, name}
		if err := txn.Delete("person", person); err != nil {
			panic(err)
		} else {
			// Commit the transaction
			txn.Commit()
		}
	} else {
		fmt.Println("Delete Person endpoint couldn't get db connection")
	}
	fmt.Println("Delete Person endpoint completed")
}
