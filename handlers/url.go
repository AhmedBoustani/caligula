package handlers

import (
  "fmt"
  "encoding/json"
  "net/http"

  "caligula/urlShortner"
)

type req_body struct {
  URL string
}

func AddShortUrl(w http.ResponseWriter, r *http.Request) {
  var u req_body
  if r.Body == nil {
     http.Error(w, "Please send a request body", 400)
     return
  }

  err := json.NewDecoder(r.Body).Decode(&u)
  if err != nil {
    http.Error(w, err.Error(), 400)
    return
  }

  s := urlShortner.Encode(u.URL)

  fmt.Println(s)
  fmt.Println(len(s))
}
