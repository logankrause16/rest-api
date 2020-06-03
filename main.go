package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandleMain)
	r.HandleFunc("/index", HandleIndex)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func HandleMain(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Reached HandleMain")
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index!")
}
