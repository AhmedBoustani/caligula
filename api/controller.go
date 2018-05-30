package api

import (
  "encoding/json"
  "net/http"

  "github.com/gorilla/mux"

  "caligula/caligula"
)

func AddShortUrl(w http.ResponseWriter, r *http.Request) {
  var u reqBody
  if r.Body == nil {
     http.Error(w, "Please send a request body", 400)
     return
  }

  err := json.NewDecoder(r.Body).Decode(&u)
  if err != nil {
    http.Error(w, err.Error(), 400)
    return
  }

  s := resBody{u.URL, caligula.Shorten(u.URL)}

  res, err := json.Marshal(s)
  w.WriteHeader(201)
  w.Header().Set("Content-Type", "application/json")
  w.Write(res)
}

func FetchLongUrl(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  val, err := caligula.Db.Find(vars["url"])
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }

  if val == "" {
    http.Error(w, "Url not found", 404)
    return
  }

  http.Redirect(w, r, val, http.StatusSeeOther)
}
