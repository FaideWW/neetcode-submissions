func isAnagram(s string, t string) bool {
	sSeen  := make(map[rune]int)

	for _, c := range s {
		sSeen[c]++
	}
	for _, c := range t {
		if v, ok := sSeen[c]; !ok  || v <= 0 {
			return false
		}
		sSeen[c]--
	}
	
	for _, v := range sSeen {
		if v > 0 {
			return false
		}
	}
	return true
}
