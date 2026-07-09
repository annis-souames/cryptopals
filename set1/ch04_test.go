package set1

import (
	"bufio"
	"encoding/hex"
	"os"
	"testing"

	"github.com/annis-souames/cryptopals/internal/cryptutil"
)

func TestCh04(t *testing.T) {
	file, err := os.Open("testdata/ch4.txt")
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	// Read file line by line and store in lines string slice
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	// The key used for XORing is a single byte, so we can try all 256 possible values XORed with each line and check if the result is a valid English text. We can use a scoring function to determine how likely the result is to be English text.
	bestScore := 0
	var bestLine string
	var bestResult []byte
	var actualKey byte
	for _, line := range lines {
		for i := 0; i < 256; i++ {
			// lineData is a byte slice of the hex-decoded line
			lineData, _ := hex.DecodeString(line)
			result := cryptutil.SingleByteXOR(lineData, byte(i))
			//t.Logf("key %d, result: %s", i, result)
			score := cryptutil.ScoreEnglishText(result)
			if score > bestScore {
				bestScore = score
				bestLine = line
				bestResult = result
				actualKey = byte(i)
			}
		}
	}

	t.Logf("The best scored hashed line is: %s , key in hex is: %s / Score: %d", bestLine, hex.EncodeToString([]byte{actualKey}), bestScore)
	t.Logf("Original input is %s", bestResult)
}
