package database
import (
  "log"
  "github.com/dgraph-io/badger"
)

func Init() *badger.DB {
  opts := badger.DefaultOptions
  opts.Dir = "badger"
  opts.ValueDir = "badger"
  db, err := badger.Open(opts)
  if err != nil {
    log.Fatal(err)
  }

  return db
}
