package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Handles "/"
func HandleMain(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Reached HandleMain")
}

// Handles "/index"
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index!")
}

// Handles "/userId/{id}"
func HandleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	fmt.Fprintln(w, "User ID:", userId)
}

// Handles "/data-dump"
func HandleDump(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Fprintln(w, string(data))
	}
}
