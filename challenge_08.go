package main

// DetectECBMode detects if the ciphertext has been encrypted using AES-128-ECB
func DetectECBMode(hexCiphertext string) bool {
	blocks := ChunksOf([]byte(hexCiphertext), 32)

	blockCounts := map[string]int{}
	for _, block := range blocks {
		sBlock := string(block)
		c, exists := blockCounts[sBlock]
		if !exists {
			blockCounts[sBlock] = 1
		} else {
			blockCounts[sBlock] = c + 1
		}
	}

	for _, count := range blockCounts {
		if count > 1 {
			return true
		}
	}

	return false
}

// FindECBLineNumber finds the line that is encrypted using AES-128-ECB
func FindECBLineNumber(lines []string) int {
	for i, line := range lines {
		if DetectECBMode(line) {
			return i
		}
	}

	return -1
}
