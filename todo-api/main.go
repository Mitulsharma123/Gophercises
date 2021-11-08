// creating our first server
package main

import (
	"encoding/json"
	"gophercises/todo-api/structs"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)
		}
		//w.Write([]byte("hello-world"))
	})
	//mux takes request from web appliaction and routes it to API
	http.ListenAndServe("localhost:3000", mux)

}
