// creating our first server
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received")
		fmt.Println(r.Method)
		w.Write([]byte("hello-world"))
	})
	//mux takes request from web appliaction and routes it to API
	http.ListenAndServe("localhost:3000", mux)

}
