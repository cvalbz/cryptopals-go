package main

import (
	"bytes"
	"testing"
)

func TestDetectSingleByteXOR(t *testing.T) {
	lines := SplitLines(ReadFile("./data/4.txt"))

	line := DetectSingleByteXOR(lines)
	expectedLine := 170

	if line != expectedLine {
		t.Errorf("Expected key to be %v, but got %v", expectedLine, line)
	}

	expectedMessage := []byte("Now that the party is jumping\n")
	ciphertext := HexToBytes([]byte(lines[expectedLine]))
	key, _ := SingleByteXOR(ciphertext)
	message := DecryptSingleByteXOR(key, ciphertext)

	if !bytes.Equal(message, expectedMessage) {
		t.Errorf("Expected key to be %v, but got %v", string(expectedMessage), string(message))
	}
}
