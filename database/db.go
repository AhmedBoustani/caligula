package database

import (
  "log"

  "github.com/dgraph-io/badger"
)

type Database struct {
  opts  badger.Options
  name  string
}

func New(name string) *Database {
  return &Database{ name: name }
}

func (d *Database) Init() {
  opts := badger.DefaultOptions
  opts.Dir = "badger/" + d.name
  opts.ValueDir = "badger/" + d.name
}

func (d *Database) Find(key string) (string, error) {
  db, err := badger.Open(d.opts)
  if err != nil {
    log.Println(d.name, "open error: ", err.Error())
    return "", err
  }
  defer db.Close()

  hit := ""
  err = db.View(func(txn *badger.Txn) error {
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

func (d *Database) Add(key string, value string) error {
  db, err := badger.Open(d.opts)
  if err != nil {
    log.Println(d.name, "open error: ", err.Error())
    return err
  }
  defer db.Close()

  err = db.Update(func(tx *badger.Txn) error {
    txn := db.NewTransaction(true)

    defer txn.Discard()

    err := txn.Set([]byte(key), []byte(value))
    if err != nil {
      return err
    }

    if err := txn.Commit(nil); err != nil {
      return err
    }
    return nil
  })

  return err
}