package caligula

import (
  "log"

  "caligula/database"
)

var Db database.Database

func InitDB() *database.Database {
  Db := database.New("caligula")
  Db.Init()
  return Db
}

func Shorten(url string) string {
  log.SetFlags(log.LstdFlags | log.Lmicroseconds)
  log.Println("encoding the hex...")
  s := encode(url)
  log.Println("adding to the database...")
  Db.Add(s, url)
  log.Println("added to the database")
  return s
}
