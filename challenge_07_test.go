package main

import (
	"bytes"
	"crypto/rand"
	"testing"
)

func TestAES128Decrypt(t *testing.T) {
	ciphertext := Base64ToBytes(Unlines(string(ReadFile("./data/7.txt"))))
	key := []byte("YELLOW SUBMARINE")

	plaintext := AES128DecryptECB(ciphertext, key)
	plaintextBlock := plaintext[:16]
	expected := []byte("I'm back and I'm")

	if !bytes.Equal(expected, plaintextBlock) {
		t.Errorf("Expected first plaintext block to be %v, but got %v", plaintextBlock, expected)
	}
}

func TestAES128EncryptDecrypt(t *testing.T) {
	message := make([]byte, 160)
	_, err := rand.Read(message)
	check(err)

	key := make([]byte, 16)
	_, err = rand.Read(key)
	check(err)

	encrypted := AES128EncryptECB(message, key)
	decrypted := AES128DecryptECB(encrypted, key)

	if !bytes.Equal(message, decrypted) {
		t.Errorf("Assertion dec(enc(message)) == message failed.")
	}
}
