package main

import (
  "net/http"

  "caligula/handlers"
)

type RouteHandler func(http.ResponseWriter, *http.Request)

type Route struct {
  Method      string
	Path        string
	HandlerFunc RouteHandler
}

type Routes []Route


func AllRoutes() Routes {
	routes := Routes{
    Route{"GET", "/", handlers.Welcome},
	}
	return routes
}
