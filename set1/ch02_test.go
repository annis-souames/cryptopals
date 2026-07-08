// Package set1 solves Cryptopals Set 1.
// Challenge 2: XOR  — https://cryptopals.com/sets/1/challenges/2
package set1

import (
    "encoding/hex"
    "testing"

    "github.com/annis-souames/cryptopals/internal/cryptutil"
)

func TestCh02(t *testing.T) {
	inputA := "1c0111001f010100061a024b53535009181c"
	inputB := "686974207468652062756c6c277320657965"
	dataA, _ := hex.DecodeString(inputA)
	dataB, _ := hex.DecodeString(inputB)
    got, _ := cryptutil.FixedXOR(dataA, dataB)
    want := "746865206b696420646f6e277420706c6179"
    if hex.EncodeToString(got) != want {
        t.Errorf("got %x, want %s", got, want)
    }
}
