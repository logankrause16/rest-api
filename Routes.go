// Package classification of Product API
//
// Documentation for Test API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
// swagger:meta
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

// Creates Route for all endpoints in API.
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	sh := http.StripPrefix("/docs/", http.FileServer(http.Dir("./swaggerui/")))
	router.PathPrefix("/docs/").Handler(sh)

	return router
}

// Defines the Routes. If HTTP type is not directly define, default to "GET"
var routes = Routes{

	Route{
		"HandleMain",
		"GET",
		"/v2/data/",
		HandleMain,
	},

	Route{
		"HandleIndex",
		"GET",
		"/v2/data/index",
		HandleIndex,
	},

	Route{
		"HandleUser",
		"GET",
		"/v2/data/userId/{id}",
		HandleUser,
	},

	Route{
		"HandleDump",
		"GET",
		"/data-dump",
		HandleDump,
	},
}
