// Package set1 solves Cryptopals Set 1.
// Challenge 2: XOR  — https://cryptopals.com/sets/1/challenges/2
package set1

import (
    "encoding/hex"
    "testing"

    "github.com/annis-souames/cryptopals/internal/cryptutil"
)

func TestChallenge02(t *testing.T) {
    a, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
    b, _ := hex.DecodeString("686f6c6c6f7768696465736e6f6f6b")
    got, _ := cryptutil.FixedXOR(a, b)
    want := "746865206b696420646f6e277420706c6179"
    if hex.EncodeToString(got) != want {
        t.Errorf("got %x, want %s", got, want)
    }
}
