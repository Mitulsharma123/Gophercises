package controller

import (
	"net/http"
)

//servemux is router
func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	return mux
}
