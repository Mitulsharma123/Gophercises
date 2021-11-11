package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic", panicDemo)
	mux.HandleFunc("/panic-after", panicAfterDemo)
	//mux.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover() // that msg into panic and recover in this value
		fmt.Fprint(w, err)
	}()
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("oh NO !")
}
