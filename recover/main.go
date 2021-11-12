package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	//mux.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":3000", recoverMw(mux, true)))
}

func recoverMw(app http.Handler, dev bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				stack := debug.Stack()
				log.Println(string(stack))
				if !dev {
					http.Error(w, "something went wrong", http.StatusInternalServerError)
					return
				}
				fmt.Fprint(w, "<h1>panic: %v<h1><pre>%s</pre>", err, string(stack))
			}
		}()

		app.ServeHTTP(w, r)
	}
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
	panic("oh NO!")
}

/*
func hello() {
	fmt.Println("Hello")
}
*/
