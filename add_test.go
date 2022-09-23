package main

import (
	"encoding/binary"
	"github.com/bszcz/mt19937_64"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"log"
	"math/big"
	"testing"
)

func Seed2Uint256T(seed uint32) []byte {
	r := mt19937_64.New()
	r.Seed(int64(seed))
	b := make([]byte, 32)
	binary.BigEndian.PutUint64(b[24:], r.Uint64())
	binary.BigEndian.PutUint64(b[16:], r.Uint64())
	binary.BigEndian.PutUint64(b[8:], r.Uint64())
	binary.BigEndian.PutUint64(b[0:], r.Uint64())
	return b[:]
}

func TestKeyAdd(t *testing.T) {
	key1 := Seed2Uint256T(1)
	key2 := Seed2Uint256T(2)
	key1BigInt := big.NewInt(0).SetBytes(key1)
	log.Println(key1BigInt)
	key2BigInt := big.NewInt(0).SetBytes(key2)

	key1Pkx, key1Pky := secp256k1.S256().ScalarBaseMult(key1)
	key2Pkx, key2Pky := secp256k1.S256().ScalarBaseMult(key2)
	pkSumx, pkSumy := secp256k1.S256().Add(key1Pkx, key1Pky, key2Pkx, key2Pky)
	log.Println("pkSumx:", pkSumx, "pkSumy:", pkSumy)

	keySum := big.NewInt(0).Add(key1BigInt, key2BigInt)
	keySumx, keySumy := secp256k1.S256().ScalarBaseMult(keySum.Bytes())
	log.Println("keySumx:", keySumx, "keySumy:", keySumy)

	pk1Pkx, pk1Pky := secp256k1.S256().Add(keySumx, keySumy, key2Pkx, key2Pky.Neg(key2Pky))
	log.Println("pk1kx:", pk1Pkx, "pk1Pky:", pk1Pky)
	log.Println("key1Pkx", key1Pkx, "key1pky", key1Pky)
}

func TestGenDiffInOrder(t *testing.T) {
	diff := GenDiffInOrder(0x10000002)
	log.Println(big.NewInt(0).SetBytes(diff))
}
