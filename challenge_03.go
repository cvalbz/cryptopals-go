package main

import (
	"math"
)

// [26] - spaces
// [27] - punctuation
// [28] - unprintable
type LetterFreqDist [29]float64

func GetEnglishFreqDist() LetterFreqDist {
	return [29]float64{
		0.0651738, 0.0124248, 0.0217339, 0.0349835, 0.0941442,
		0.0197881, 0.0158610, 0.0492888, 0.0458094, 0.0009033,
		0.0050529, 0.0331490, 0.0202124, 0.0464513, 0.0496302,
		0.0137645, 0.0008606, 0.0497563, 0.0515760, 0.0629357,
		0.0225134, 0.0082903, 0.0171272, 0.0013692, 0.0145984,
		0.0007836, 0.1618182, 0.8, 0.000001,
	}
}

func ChiSquared(observed LetterFreqDist, expected LetterFreqDist) float64 {
	result := float64(0)
	for i, x := range observed {
		d := x - expected[i]
		result += (d * d) / expected[i]
	}
	return result
}

func buildFreqDist(bs []byte) LetterFreqDist {
	var result LetterFreqDist

	totalChars := 0
	punctuation := 0
	spaces := 0
	unprintable := 0
	for _, b := range bs {
		totalChars++

		if b >= 97 && b <= 122 {
			result[b-97]++
			continue
		}

		if b >= 65 && b <= 90 {
			result[b-65]++
			continue
		}

		if b == 32 {
			spaces++
			continue
		}

		if (b == 44) || (b == 46) || (b == 39) {
			punctuation++
			continue
		}

		unprintable++
	}

	result[26] = float64(spaces) / float64(totalChars)
	result[27] = float64(punctuation) / float64(totalChars)
	result[28] = float64(unprintable) / float64(totalChars)
	for i, b := range result {
		result[i] = b / float64(totalChars)
	}

	return result
}

func SingleByteXOR(ciphertext []byte) (key byte, minChi float64) {
	minChi = math.MaxFloat64

	for keyCandidate := byte(0); keyCandidate < 255; keyCandidate++ {
		messageCandidate := DecryptSingleByteXOR(byte(keyCandidate), ciphertext)

		letterFreqDist := buildFreqDist(messageCandidate)
		chi := ChiSquared(letterFreqDist, GetEnglishFreqDist())

		if chi < minChi {
			minChi = chi
			key = keyCandidate
		}
	}
	return
}

func DecryptSingleByteXOR(key byte, ciphertext []byte) []byte {
	streamKey := make([]byte, len(ciphertext))
	for i := range ciphertext {
		streamKey[i] = key
	}

	return XOR(ciphertext, streamKey)
}
