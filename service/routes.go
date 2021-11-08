package service

import "net/http"

type Route struct {
	Name, Method, Pattern string
	HandlerFunc           http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"getReminder",
		"GET",
		"/reminder/{id}",
		func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("content-type", "application/json")
			writer.Write([]byte("{\n \"Result\" : \"OK\" \n}"))
		},
	},
}
