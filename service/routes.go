package service

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// Initialize our routes
var routes = Routes{
	Route{
		"info",
		"GET",
		"/info",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			_, _ = w.Write([]byte("{\"result\":\"OK\"}"))
		},
	},
}
