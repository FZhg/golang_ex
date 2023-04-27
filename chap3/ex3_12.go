package chap3

// Assume UTF-8 encoding

func isAnagram(s1 string, s2 string) bool {
	charFreqS1 := getCharFreq(s1)

	for _, r := range s2 {
		charFreqS1[r]--
		if charFreqS1[r] < 0 {
			return false
		}
	}
	return true
}

func getCharFreq(s string) map[rune]int {
	charFreq := make(map[rune]int, len(s))
	for _, r := range s {
		charFreq[r]++
	}
	return charFreq
}
