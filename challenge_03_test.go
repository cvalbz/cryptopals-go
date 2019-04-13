package main

import (
	"bytes"
	"testing"
)

func TestSingleByteXOR(t *testing.T) {
	ciphertext := HexToBytes([]byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
	key, _ := SingleByteXOR(ciphertext)
	expectedKey := byte(88)

	expectedMessage := []byte("Cooking MC's like a pound of bacon")
	message := DecryptSingleByteXOR(key, ciphertext)

	if key != expectedKey {
		t.Errorf("Expected key to be %v, but got %v", expectedKey, key)
	}

	if !bytes.Equal(message, expectedMessage) {
		t.Errorf("Expected key to be %v, but got %v", expectedMessage, message)
	}
}
