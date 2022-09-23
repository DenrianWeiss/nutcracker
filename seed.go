package main

import (
	"encoding/binary"
	"github.com/bszcz/mt19937_64"
)

func Seed2Uint256(seed uint32) []byte {
	r := mt19937_64.New()
	r.Seed(int64(seed))
	b := make([]byte, 32)
	binary.BigEndian.PutUint64(b[24:], r.Uint64())
	binary.BigEndian.PutUint64(b[16:], r.Uint64())
	binary.BigEndian.PutUint64(b[8:], r.Uint64())
	binary.BigEndian.PutUint64(b[0:], r.Uint64())
	return b[:]
}
