package set1

import (
	"encoding/hex"
	"testing"
	"github.com/annis-souames/cryptopals/internal/cryptutil"
)


func TestCh03(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	data,_ := hex.DecodeString(input)
	// We all try all combination of a 1 byte key: 256 possibilities (0-255)
	for key := 0; key < 256; key++ {
		got := cryptutil.SingleByteXOR(data, byte(key))
		// Here we need to check if got is english-like
		// For simplicity, we will just print the result for now
		t.Logf("Key: %d, Result: %s, Score: %d", key, got, scoreEnglishLike(got))
	}
}



// This function scores how "English-like" a byte slice is.
func scoreEnglishLike(data []byte) int {
	score := 0
	for _, b := range data {
		switch b {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			score += 1
		case ' ', '\n', '\t':
			score += 1
		}
	}
	return score	
}
