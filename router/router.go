package router

import (
  "github.com/gorilla/mux"
)

func Setup() *mux.Router {
  router := mux.NewRouter()
  routes := getAllRoutes()

  for _, route := range routes {
		handle := Logger(route.HandlerFunc)

		router.HandleFunc(route.Path, handle).Methods(route.Method)
}

  return router
}
