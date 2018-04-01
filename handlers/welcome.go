package handlers

import (
  "fmt"
  "net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "not implemented yet !")
}
