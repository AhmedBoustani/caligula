package main

import (
  "fmt"
  "log"
  "net/http"

  "caligula/caligula"
  "caligula/router"
)

func main() {
  caligula.InitDB()

  fmt.Println("Running on port 3000")
  if err := http.ListenAndServe(":3000", router.Setup()); err != nil {
		log.Fatal(err)
  }
}
