package main

import (
	"bytes"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	a := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	b := []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")

	actual := HexToBase64(a)
	if !bytes.Equal(b, actual) {
		t.Errorf("Expected encoding to be %v, but got %v", b, actual)
	}
}
