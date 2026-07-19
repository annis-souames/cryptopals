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
		// determine how many blocks of size keySize we can take from the input
		numBlocks := len(input) / keySize
		// iterate through the blocks per keysize worth of bytes, calculate normalized hamming dist.
		totalNormalizedH := 0.0
		for i := 0; i < numBlocks-1; i++ {
			H := cryptutil.HammingDist(input[i*keySize:(i+1)*keySize], input[(i+1)*keySize:(i+2)*keySize])
			normalizedH := float64(H) / float64(keySize)
			totalNormalizedH += normalizedH
		}
		// compute average of distances array for this pass/keysize:
		averageNormalizedH := totalNormalizedH / float64(numBlocks-1)
		t.Logf("Key size: %d, Average normalized Hamming distance: %f", keySize, averageNormalizedH)
	}
}
