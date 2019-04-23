package main

import (
	"testing"
)

func TestDetectECBMode(t *testing.T) {
	hexCiphertext := "72d8a640970bf9ce9f9f3a87fb908eebf7e603fb0589a00893381cc6662726130bc564c42ccef93447bb74065650138396d1452408641b05990cb6c0e73b1a93bed3e428c70a0a073777089d555b9195fb371bcaff16a24487918737f2fe4d43597664daf610196b3dbc25ade7e68cbe3be39712a1508e3a85531bd175c6484a48058104d6d6f18fd71db2bfd61e72987b092edced52e3f848e7a39e90ed07f5"
	detected := DetectECBMode(hexCiphertext)
	expected := false
	if expected != detected {
		t.Errorf("Expected ECB detection to be %v, but got %v", expected, detected)
	}

	hexCiphertext = "72d8a640970bf9ce9f9f3a87fb908eeb72d8a640970bf9ce9f9f3a87fb908eeb72d8a640970bf9ce9f9f3a87fb908eeb72d8a640970bf9ce9f9f3a87fb908eeb72d8a640970bf9ce9f9f3a87fb908eeb72d8a640970bf9ce9f9f3a87fb908eeb72d8a640970bf9ce9f9f3a87fb908eeb72d8a640970bf9ce9f9f3a87fb908eeb72d8a640970bf9ce9f9f3a87fb908eeb72d8a640970bf9ce9f9f3a87fb908eeb"
	detected = DetectECBMode(hexCiphertext)
	expected = true
	if expected != detected {
		t.Errorf("Expected ECB detection to be %v, but got %v", expected, detected)
	}
}

func TestFindECBLineNumber(t *testing.T) {
	ciphertextLines := SplitLines(ReadFile("./data/8.txt"))
	ECBLine := FindECBLineNumber(ciphertextLines)
	expected := 132

	if expected != ECBLine {
		t.Errorf("Expected ECB line to be %v, but got %v", expected, ECBLine)
	}
}
