package main

// RepeatingKeyXOR encrypts a plaintext using a Repeating XOR Key (aka Vigenere)
func RepeatingKeyXOR(plaintext []byte, key []byte) []byte {
	keystream := make([]byte, len(plaintext))
	for i := range keystream {
		keystream[i] = key[i%len(key)]
	}

	return XOR(plaintext, keystream)
}
