package cryptutil

// This function scores how "English-like" a byte slice is.
func ScoreEnglishText(data []byte) int {
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
