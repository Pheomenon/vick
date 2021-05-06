package vick

import "hash/crc32"

func hashing(key []byte) uint32{
	c := crc32.New(CrcTable)
	c.Write(key)
	return c.Sum32()
}
