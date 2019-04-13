package main

// XOR performs element-by-element xor operation
func XOR(a []byte, b []byte) []byte {
	if len(a) != len(b) {
		panic("Different lenghts are not supported for XOR operation!")
	}

	result := make([]byte, len(a))
	for i, x := range a {
		result[i] = b[i] ^ x
	}

	return result
}
