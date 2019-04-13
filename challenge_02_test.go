package main

import (
	"bytes"
	"testing"
)

func TestXOR(t *testing.T) {
	a := HexToBytes([]byte("1c0111001f010100061a024b53535009181c"))
	b := HexToBytes([]byte("686974207468652062756c6c277320657965"))
	expected := HexToBytes([]byte("746865206b696420646f6e277420706c6179"))

	actual := XOR(a, b)
	if !bytes.Equal(actual, expected) {
		t.Errorf("Expected result to be %v, but got %v", expected, actual)
	}
}
