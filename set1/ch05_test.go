package set1

import (
	"encoding/hex"
	"testing"

	"github.com/annis-souames/cryptopals/internal/cryptutil"
)

func TestCh05(t *testing.T) {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	// Repeating-key XOR cycles the key across the input
	encrypted := cryptutil.RepeatingKeyXOR([]byte(input), []byte(key))
	got := hex.EncodeToString(encrypted)

	if got != expected {
		t.Fatalf("got %s, want %s", got, expected)
	}

	t.Logf("Encrypted output is: %s", got)
}
