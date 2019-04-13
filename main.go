package main

import "fmt"

func main() {
	lines := SplitLines(ReadFile("./data/4.txt"))

	line := DetectSingleByteXOR(lines)
	fmt.Println("detected line is", line)
}
