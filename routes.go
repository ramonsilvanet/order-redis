package main

import (
	"net/http"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"merchantSaveOrUpdate",
		"POST",
		"/merchant",
		merchantSaveOrUpdate,
	},
	Route{
		"getMerchantLastUpdate",
		"GET",
		"/merchant/{merchant_id}",
		getMerchantLastUpdate,
	},
	Route{
		"merchantSaveDummyData",
		"GET",
		"/dummy",
		merchantSaveDummyData,
	},
}
