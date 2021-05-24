package vick

func hashing(key []byte) uint64 {
	_, err := h64.Write(key)
	if err != nil {
		panic("hash error!")
	}
	h64.Reset()
	return h64.Sum64()
}
