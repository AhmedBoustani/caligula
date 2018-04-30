package urlShortner

import (
  "crypto/md5"
  "io"
  "encoding/hex"
  "fmt"
  "math/big"
  "encoding/base64"
)

func Encode(url string) string {
  urlHex := hex.EncodeToString(computeMD5(url))
  b := toBinary(urlHex, 16)
  bytes := []byte(b)[0:14]

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
