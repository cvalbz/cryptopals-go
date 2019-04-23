package main

import (
	"io/ioutil"
	"math"
	"strings"
)

// check panics if there is an error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ReadFile reads data from a file
func ReadFile(filepath string) []byte {
	data, err := ioutil.ReadFile(filepath)
	check(err)

	return data
}

// SplitLines splits a byte-slice in lines and converts each line to string
func SplitLines(data []byte) []string {
	result := strings.Split(string(data), "\n")
	return result
}

// DetectSingleByteXOR detects which line has been encrypted with a Single Byte XOR
func DetectSingleByteXOR(lines []string) int {
	detectedLine := 0
	minChi := math.MaxFloat64
	for i, line := range lines {
		lineDecoded := HexToBytes([]byte(line))
		_, chi := SingleByteXOR(lineDecoded)

		if chi < minChi {
			minChi = chi
			detectedLine = i
		}
	}

	return detectedLine
}
