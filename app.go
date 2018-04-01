package main

import (
	"log"
  "net/http"
)

func main() {
  r := setupRouter(AllRoutes())

  if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
  }
}
