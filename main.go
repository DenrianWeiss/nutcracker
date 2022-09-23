package main

import (
	"log"
	"time"
)

func main() {
	crackTarget := []byte{}
	log.Println(time.Now())
	rM := BuildPkMapRange(0, 0xFFFFFFFF)
	for i := 0; i < 0xFFFFFFFFFFFF; i++ {
		r := [33]byte{}
		dpk := SubPk(crackTarget, GenDiffInOrder(uint64(i)))
		copy(r[:], dpk)
		if v, ok := rM[r]; ok {
			println("Successfully found private key components:", v, i)
		}
	}
	log.Println(time.Now())
}
