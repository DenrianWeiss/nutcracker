package main

func BuildPkMapRange(s, r uint32) map[[33]byte]uint32 {
	resp := make(map[[33]byte]uint32)
	for i := s; i < r; i++ {
		k := [33]byte{}
		pk := Seed2Uint256(i)
		cpk := Pk2Compressed(pk)
		copy(k[:], cpk)
		resp[k] = i
	}
	return resp
}
