package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB //global variable connection to have access to DB layer from model

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234$@tcp(localhost:3306)/mysql")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB")
	con = db
	return db
}
