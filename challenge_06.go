package main

import (
	"encoding/base64"
	"math"
	"strings"
)

// CountSetBits counts number of 1 bits
func CountSetBits(v byte) byte {
	v = (v & 0x55) + ((v >> 1) & 0x55)
	v = (v & 0x33) + ((v >> 2) & 0x33)
	return (v + (v >> 4)) & 0xF
}

// HammingDistance computes Hamming (edit) distance between two byte-slices
func HammingDistance(a []byte, b []byte) int64 {
	xored := XOR(a, b)

	nrOneBits := int64(0)
	for _, x := range xored {
		nrOneBits += int64(CountSetBits(x))
	}

	return nrOneBits
}

// Base64ToBytes decodes a base64 encoded string to a byte-slice
func Base64ToBytes(s string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(s)
	check(err)

	return decoded
}

// Unlines concatenates all lines in one long line
func Unlines(s string) string {
	return strings.Replace(s, "\n", "", -1)
}

// NormalizedHammingDistance computes a normalized Hamming distance based on the length of input
func NormalizedHammingDistance(doubleBlock []byte, keySize int) float64 {
	if len(doubleBlock) != 2*keySize {
		panic("Block size is wrong")
	}

	a := doubleBlock[:keySize]
	b := doubleBlock[keySize:]
	return float64(HammingDistance(a, b)) / float64(keySize)
}

// DetectKeySize detects the probable key size used for a ciphertext
func DetectKeySize(ciphertext []byte) int {
	minNormalizedDist := math.MaxFloat64
	keySize := 1

	for sizeCandidate := 2; sizeCandidate < 40; sizeCandidate++ {
		nrDoubleBlocks := len(ciphertext) / (2 * sizeCandidate)

		meanNormalizedDistance := float64(0)
		for i := 0; i < nrDoubleBlocks; i++ {
			doubleBlock := ciphertext[i*2*sizeCandidate : (i+1)*2*sizeCandidate]
			meanNormalizedDistance += NormalizedHammingDistance(doubleBlock, sizeCandidate)
		}
		meanNormalizedDistance /= float64(nrDoubleBlocks)

		if meanNormalizedDistance < minNormalizedDist {
			minNormalizedDist = meanNormalizedDistance
			keySize = sizeCandidate
		}
	}

	return keySize
}

// ChunksOf splits a byte-slice into equal chunks
func ChunksOf(b []byte, size int) [][]byte {
	result := make([][]byte, len(b)/size)

	for i := range result {
		//result[i] = make([]byte, size)
		result[i] = b[i*size : (i+1)*size]
	}

	return result
}

// TransposeMatrix transposes a two dimensional byte-slice
func TransposeMatrix(m [][]byte) [][]byte {
	rowsNr := len(m)
	if rowsNr == 0 {
		return [][]byte{}
	}

	columnsNr := len(m[0])

	result := make([][]byte, columnsNr)
	for i := range result {
		result[i] = make([]byte, rowsNr)
		for j := range result[i] {
			result[i][j] = m[j][i]
		}
	}

	return result
}

// BreakRepeatingKeyXOR finds the key of a ciphertext encrypted using a Repeating XOR Cipher
func BreakRepeatingKeyXOR(ciphertext []byte) []byte {
	keySize := DetectKeySize(ciphertext)

	blocks := ChunksOf(ciphertext, keySize)
	transposedBlocks := TransposeMatrix(blocks)

	result := make([]byte, keySize)
	for i, tb := range transposedBlocks {
		keyByte, _ := SingleByteXOR(tb)
		result[i] = keyByte
	}

	return result
}
