package caligula

import "log"

func isKeyExists(key string) bool {
  hit, err := Db.Find(key)
  if err != nil {
    log.Fatal(err)
  }
  return hit != ""
}
