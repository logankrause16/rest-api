package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

/*
	r.HandleFunc("/", HandleMain)
	r.HandleFunc("/index", HandleIndex)
	r.HandleFunc("/user/{id}", HandleUser)
	r.HandleFunc("/data-dump", HandleDump)
*/

var routes = Routes{

	Route{
		"HandleMain",
		"GET",
		"/",
		HandleMain,
	},

	Route{
		"HandleIndex",
		"GET",
		"/index",
		HandleIndex,
	},

	Route{
		"HandleUser",
		"GET",
		"/user/{id}",
		HandleUser,
	},

	Route{
		"HandleDump",
		"GET",
		"/data-dump",
		HandleDump,
	},
}
