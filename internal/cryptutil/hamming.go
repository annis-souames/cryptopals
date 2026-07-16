package cryptutil

import (
	"math/bits"
)

// This function computes the hamming distance between two bit strings
func HammingDist(a, b []byte) int {
	// check that a and b have same length
	if len(a) != len(b) {
		panic("HammingDist: a and b must have same length")
	}
	dist := 0
	for i := 0; i < len(a); i++ {
		diff := a[i] ^ b[i]
		// count the number of 1s in diff - use fast SWAR technique
		dist += bits.OnesCount8(diff)
	}
	return dist
}
