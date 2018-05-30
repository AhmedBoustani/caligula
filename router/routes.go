package router

import (
  "net/http"

  "caligula/api"
)

type RouteHandler func(http.ResponseWriter, *http.Request)

type Route struct {
  Method      string
	Path        string
	HandlerFunc RouteHandler
}

type Routes []Route

func getAllRoutes() Routes {
	routes := Routes{
    Route{ "POST", "/", api.AddShortUrl },
    Route{ "GET", "/{url}", api.FetchLongUrl },
	}
	return routes
}
