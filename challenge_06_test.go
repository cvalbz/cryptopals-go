package main

import (
	"bytes"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	a := []byte("this is a test")
	b := []byte("wokka wokka!!!")

	editDistance := HammingDistance(a, b)
	expected := int64(37)

	if expected != editDistance {
		t.Errorf("Expected hamming distance to be %v, but got %v", expected, editDistance)
	}
}

func TestDetectKeySize(t *testing.T) {
	ciphertext := Base64ToBytes(Unlines(string(ReadFile("./data/6.txt"))))

	keySize := DetectKeySize(ciphertext)
	expected := 29

	if expected != keySize {
		t.Errorf("Expected key size to be %v, but got %v", expected, keySize)
	}
}

func TestBreakRepeatingKeyXOR(t *testing.T) {
	ciphertext := Base64ToBytes(Unlines(string(ReadFile("./data/6.txt"))))

	key := BreakRepeatingKeyXOR(ciphertext)
	expected := []byte("Terminator X: Bring the noise")

	if !bytes.Equal(key, expected) {
		t.Errorf("Expected key to be %s, but got %s", expected, key)
	}
}
