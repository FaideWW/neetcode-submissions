func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[uint64][]string)
	for _, s := range strs {
			code := encode(s)
			if _, ok := anagramMap[code]; !ok {
				anagramMap[code] = []string{}
			}
			anagramMap[code] = append(anagramMap[code], s)
	}

	anagrams := [][]string{}
	for _, anas := range anagramMap {
		anagrams = append(anagrams, anas)
	}
	return anagrams
}

var PRIMES []int = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

func encode(s string) uint64 {
	seen := make(map[byte]int)
	for _, c := range s {
		seen[byte(c)]++
	}

	var value uint64 = 1
	for c, n := range seen {
		idx := c - byte('a')
		value = value * pow(PRIMES[idx], n)
	}

	return value
}

func pow(base int, exp int) uint64 {
	var v uint64 = 1
	for  range exp {
		v = v * uint64(base)
	}
	return v
}