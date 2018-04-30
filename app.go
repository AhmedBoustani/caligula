package main

import (
	"log"
  "net/http"
	"fmt"

	"caligula/database"
)

func main() {
	db := database.Init()
	defer db.Close()

  router := setupRouter(AllRoutes())

  fmt.Println("Running on port 3000")
  if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
  }
}
