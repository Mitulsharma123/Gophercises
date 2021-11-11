// creating our first server
package main

import (
	"fmt"
	"gophercises/todo-api/controller"
	"gophercises/todo-api/model"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" //mysql driver
)

func main() {
	mux := controller.Register() //create a new mux
	db := model.Connect()        //connect to DB
	defer db.Close()             //whenever we tear down server,close the DB
	fmt.Println("serving.....")
	log.Fatal(http.ListenAndServe(":3000", mux)) //listen and serve at port 3000 and pass mux
}
