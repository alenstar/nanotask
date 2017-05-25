package utils

import (
	"hash/crc32"
	"hash/crc64"
)

func CRC32(b []byte) uint32 {
	return crc32.Checksum(b, crc32.MakeTable(crc32.IEEE))
}

func CRC64(b []byte) uint64 {
	return crc64.Checksum(b, crc64.MakeTable(crc64.ISO))
}
