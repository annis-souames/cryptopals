package set1

import (
	"testing"

	"github.com/annis-souames/cryptopals/internal/cryptutil"
)

// TestHammingDistance tests the Hamming distance function
func TestHammingDist(t *testing.T) {
	dist := cryptutil.HammingDist([]byte("this is a test"), []byte("wokka wokka!!!"))
	t.Logf("Hamming distance between 'this is a test' and 'wokka wokka!!!' is: %d", dist)
}
