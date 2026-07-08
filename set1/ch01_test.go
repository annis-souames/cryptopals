// Package set1 solves Cryptopals Set 1.
// Challenge 1: hex to base64 — https://cryptopals.com/sets/1/challenges/1
package set1

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
	"fmt"
)

func ch1(inputHex string) (string, error) {
	decoded, err := hex.DecodeString(inputHex)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(decoded), nil
}

func TestCh01(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d" 
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	got, err := ch1(input)
	fmt.Printf("Got: %s\n",got)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
