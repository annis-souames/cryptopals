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

func TestCh06(t *testing.T) {
	// Read the base64-encoded input from the file
	input, err := cryptutil.ReadBase64File("testdata/ch6.txt")
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	keySizeMax := 40

	for keySize := 2; keySize <= keySizeMax; keySize++ {
		// hamming dist for first 2 blocks of keysize
	}
}
