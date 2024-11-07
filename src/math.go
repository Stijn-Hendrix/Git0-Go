package main

import (
	"crypto/sha256"
	"encoding/binary"
	"strconv"
)

func hash(s string) uint64 {
	hasher := sha256.New()
	hasher.Write([]byte(s))
	hashBytes := hasher.Sum(nil)
	return binary.BigEndian.Uint64(hashBytes[:8])
}

func hashString(s string) string {
	hash := hash(s)
	hashStr := strconv.FormatUint(hash, 10)
	return hashStr
}
