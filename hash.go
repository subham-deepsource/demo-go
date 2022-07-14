package main

import (
	"github.com/pierrec/xxHash/xxHash32"
)

var xxhash = xxHash32.New(0xCAFE) // hash.Hash32

func fastHash(buf []byte) uint32 {
	xxhash.Reset()
	xxhash.Write(buf)
	return xxhash.Sum32()
}
