package main

import (
	"io/ioutil"
	"math"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filepath string) []byte {
	data, err := ioutil.ReadFile(filepath)
	check(err)

	return data
}

func SplitLines(data []byte) []string {
	result := strings.Split(string(data), "\n")
	return result
}

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
