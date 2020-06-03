package main

import "github.com/gorilla/mux"

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandleMain)
}
