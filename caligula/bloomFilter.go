package caligula

import (
  "hash/fnv"
  "io"
  "strings"

  "github.com/tidwall/murmur3"
)

type Bloom struct {
  HashTable []int
  HashFuncs []interface{}
}

func New(size int) *Bloom {
  return &Bloom{
    HashTable: make([]int, size),
    HashFuncs: []interface{}{fnvHash, primeHash, murmurHash}
  }
}

func (*Bloom) Add(input string) {
  // TODO hash
  // mod 10^9 + 7

}

func (*Bloom) IsDefinitelyNew() {
  // TODO hash, if one of them yields 0 in HashTable, then true, else false

}

func fnvHash(s string) uint32 {
  h := fnv.New32a()
  h.Write([]byte(s))
  return h.Sum32()
}

func primeHash(s string) uint32 {
  acc := 0
  for _, e := range s {
    acc = acc + int(e) * 1000000007
  }
  return uint32(acc)
}

func murmurHash(s string) uint32 {
  return murmur3.Sum32(s)
}

func xxhashHash(s string) uint32 {
  h := xxhash.New64()
  io.Copy(h, s)
  h.Sum64()
}