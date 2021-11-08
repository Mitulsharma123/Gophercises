// creating our first server
package main

import (
	"gophercises/todo-api/controller"
	"net/http"
)

func main() {
	mux := controller.Register()
	mux.HandleFunc("/ping", ping())
	//mux takes request from web appliaction and routes it to API
	http.ListenAndServe("localhost:3000", mux)
}
