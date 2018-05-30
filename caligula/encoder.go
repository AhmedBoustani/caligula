package caligula

import (
  "crypto/md5"
  "encoding/base64"
  "encoding/hex"
  "fmt"
  "io"
  "log"
  "math/big"
)

func encode(url string) string{
  log.SetFlags(log.LstdFlags | log.Lmicroseconds)
  log.Println("entering encode...")
  urlHex := hex.EncodeToString(computeMD5(url))
  enc := encodeHex(urlHex)
  if !isKeyExists(enc) {
    return enc
  }
  return encode(enc)
}

func encodeHex(urlHex string) string {
  b := toBinary(urlHex, 16)
  bytes := []byte(b[0:12])

  bi := big.NewInt(0)
  bi.SetString(string(bytes), 0)

  return base64.URLEncoding.EncodeToString(bi.Bytes())
}

func computeMD5(str string) []byte {
  h := md5.New()
  io.WriteString(h, str)
  return h.Sum(nil)
}

func toBinary(v string, base int) string {
  bi := big.NewInt(0)
  bi.SetString(v, base)
  return fmt.Sprintf("%b", bi)
}
