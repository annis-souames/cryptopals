package set1

import (
	"encoding/hex"
	"testing"

	"github.com/annis-souames/cryptopals/internal/cryptutil"
)

func TestCh03(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	data, _ := hex.DecodeString(input)
	var scores [256]int
	// We all try all combination of a 1 byte key: 256 possibilities (0-255)
	for key := 0; key < 256; key++ {
		got := cryptutil.SingleByteXOR(data, byte(key))
		scores[key] = cryptutil.ScoreEnglishText(got)
		// Here we need to check if got is english-like
		// For simplicity, we will just print the result for now
		t.Logf("Key: %d, Result: %s, Score: %d", key, got, scores[key])
	}

	keyMaxScore := maxIdx(scores[:])

	t.Logf("The key with highest score is %d and the original input is %s", keyMaxScore, cryptutil.SingleByteXOR(data, byte(keyMaxScore)))
}

// helper func to index of max in a slice
func maxIdx(slice []int) int {
	max := 0
	maxIdx := 0
	for i, v := range slice {
		if v > max {
			max = v
			maxIdx = i
		}
	}
	return maxIdx
}
