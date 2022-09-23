package main

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

var curve = secp256k1.S256()

func Pk2Compressed(r []byte) []byte {
	x, y := curve.ScalarBaseMult(r)
	cpk := secp256k1.CompressPubkey(x, y)
	return cpk
}

func SubPk(pubK []byte, privKDiff []byte) []byte {
	dx, dy := curve.ScalarBaseMult(privKDiff)
	uncompX, uncompY := secp256k1.DecompressPubkey(pubK)
	diffX, diffY := curve.Add(uncompX, uncompY, dx, dy.Neg(dy))
	return secp256k1.CompressPubkey(diffX, diffY)
}

func GenDiffInOrder(diffSrc uint64) []byte {
	head := diffSrc & 0x3FFFFF
	tail := diffSrc & 0xFFFFFFc00000
	tail = tail >> 22
	b := make([]byte, 32)
	binary.BigEndian.PutUint64(b[24:], tail)
	binary.BigEndian.PutUint64(b[16:], 0)
	binary.BigEndian.PutUint64(b[8:], 0)
	binary.BigEndian.PutUint64(b[0:], head)
	return b[:]
}
