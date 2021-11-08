// creating our first server
package main

import (
	"gophercises/todo-api/controller"
	"net/http"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe("localhost:3000", mux)
}
