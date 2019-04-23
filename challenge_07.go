package main

import (
	"crypto/aes"
)

// AES128Block is an AES block with 128 bits
type AES128Block [16]byte

// BytesToBlocks converts a byte-slice to a list of AES blocks
func BytesToBlocks(bs []byte) []AES128Block {
	if len(bs)%16 != 0 {
		panic("not a multiple of 16 bytes (128 bits)")
	}

	blocks := make([]AES128Block, len(bs)/16)
	chunks := ChunksOf(bs, 16)
	for i := range blocks {
		copy(blocks[i][:], chunks[i])
	}

	return blocks
}

// BytesToAES128Key converts a byte-slice to a single AES block
func BytesToAES128Key(key []byte) AES128Block {
	if len(key) != 16 {
		panic("key size should be 16 bytes")
	}

	var result AES128Block
	copy(result[:], key)
	return result
}

// AES128ECBEncryptBlocks encrypts AES blocks using AES-128-ECB
func AES128ECBEncryptBlocks(plaintext []AES128Block, key AES128Block) []AES128Block {
	c, err := aes.NewCipher(key[:])
	check(err)

	ciphertext := make([]AES128Block, len(plaintext))
	for i := range plaintext {
		c.Encrypt(ciphertext[i][:], plaintext[i][:])
	}

	return ciphertext
}

// AES128ECBDecryptBlocks decrypts AES blocks using AES-128-ECB
func AES128ECBDecryptBlocks(ciphertext []AES128Block, key AES128Block) []AES128Block {
	c, err := aes.NewCipher(key[:])
	check(err)

	plaintext := make([]AES128Block, len(ciphertext))
	for i := range plaintext {
		c.Decrypt(plaintext[i][:], ciphertext[i][:])
	}

	return plaintext
}

// AES128EncryptECB encrypts byte-slice using AES-128-ECB
func AES128EncryptECB(plaintext []byte, key []byte) []byte {
	plaintextBlocsk := BytesToBlocks(plaintext)
	keyBlock := BytesToAES128Key(key)

	encryptedBlocks := AES128ECBEncryptBlocks(plaintextBlocsk, keyBlock)
	result := make([]byte, len(plaintext))

	for i, encBlock := range encryptedBlocks {
		bs := [16]byte(encBlock)
		copy(result[i*16:(i+1)*16], bs[:])
	}

	return result
}

// AES128DecryptECB decrypts byte-slice using AES-128-ECB
func AES128DecryptECB(ciphertext []byte, key []byte) []byte {
	ciphertextBlocsk := BytesToBlocks(ciphertext)
	keyBlock := BytesToAES128Key(key)

	decryptedBlocks := AES128ECBDecryptBlocks(ciphertextBlocsk, keyBlock)
	result := make([]byte, len(ciphertext))

	for i, decBlock := range decryptedBlocks {
		bs := [16]byte(decBlock)
		copy(result[i*16:(i+1)*16], bs[:])
	}

	return result
}
