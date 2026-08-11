
type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	rleString := ""
	for _, s := range strs {
		rleString = rleString + fmt.Sprintf("%03d%s", len(s),s)
	}
	return rleString
}

func (s *Solution) Decode(encoded string) []string {
	idx := 0
	strs := []string{}
	for idx < len(encoded) {
		nextStrLen, _ := strconv.Atoi(encoded[idx:idx+3])
		idx = idx+3
		strs = append(strs, encoded[idx:idx+nextStrLen])
		idx = idx + nextStrLen
	}

	return strs
}
