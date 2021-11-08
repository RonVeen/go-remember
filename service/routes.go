package service

import "net/http"

type Route struct {
	Name, Method, Pattern string
	HandlerFunc           http.HandlerFunc
}

type Routes []Route

var path = "/reminder"

var routes = Routes{
	Route{
		"getReminder",
		"GET",
		path + "/{id}",
		GetReminder,
	},
	Route{
		"getAllReminders",
		"GET",
		path,
		GetAllReminders,
	},
	Route{
		"CreateReminder",
		"POST",
		path,
		CreateReminder,
	},
	Route{
		"UpdateReminder",
		"PUT",
		path + "/{id}",
		UpdateReminder,
	},
	Route{
		"DeleteReminder",
		"DELETE",
		path + "/{id}",
		DeleteReminder,
	},
}
