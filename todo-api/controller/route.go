package controller

import (
	"net/http"
)

//servemux is router; below func return mux
func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping()) //endpoint created
	mux.HandleFunc("/", crud())
	return mux
}
