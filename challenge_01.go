package main

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

// HexToBase64 converts a byte slice from hex to base64
func HexToBase64(b []byte) []byte {
	decoded := HexToBytes(b)
	return BytesToBase64(decoded)
}

// BytesToBase64 encodes a byte slice as base64
func BytesToBase64(b []byte) []byte {
	encoded := base64.StdEncoding.EncodeToString(b)
	return []byte(encoded)
}

// HexToBytes decodes hex byte slice
func HexToBytes(b []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(b)))
	n, err := hex.Decode(dst, b)
	if err != nil {
		log.Fatal(err)
	}

	return dst[:n]
}
