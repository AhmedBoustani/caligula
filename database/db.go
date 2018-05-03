package database

import (
  "log"

  "github.com/dgraph-io/badger"
)

var db *badger.DB

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

func Find(key string) (string, error) {
  hit := ""
  db := Init()
  defer db.Close()
  err := db.View(func(txn *badger.Txn) error {
    item, err := txn.Get([]byte(key))
    if err != nil {
      if err == badger.ErrKeyNotFound {
        return nil
      }
      return err
    }
    val, err := item.Value()
    if err != nil {
      return err
    }
    hit = string(val)
    return nil
  })

  if err != nil {
    log.Fatal(err)
    return "", err
  }

  return hit, nil
}

func Add(key string, value string) error {
  db := Init()
  defer db.Close()
  err := db.Update(func(tx *badger.Txn) error {
    // Start a writable transaction.
    txn := db.NewTransaction(true)

    defer txn.Discard()

    // Use the transaction...
    err := txn.Set([]byte(key), []byte(value))
    if err != nil {
      return err
    }

    // Commit the transaction and check for error.
    if err := txn.Commit(nil); err != nil {
      return err
    }
    return nil
  })

  return err
}