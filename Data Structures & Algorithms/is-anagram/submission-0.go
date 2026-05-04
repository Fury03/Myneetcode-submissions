func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	wordfreq := make(map[rune]int)

	for _, ch := range s {
		wordfreq[ch]++
	}

	for _, ch := range t {
		wordfreq[ch]--
	}

	for _, count := range wordfreq {
		if count != 0 {
			return false
		}
	}
	return true
}
