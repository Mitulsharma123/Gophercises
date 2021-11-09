// creating our first server
package main

import (
	"gophercises/todo-api/controller"
	"gophercises/todo-api/model"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" //mysql driver
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
